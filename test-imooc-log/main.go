package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

// === ReadFromFile ===
type Reader interface {
	Read(rc chan []byte)
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

func (r *ReadFromFile) Read(rc chan []byte) {
	//打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}

	//从文件末尾开始逐行读取（2读最新）
	f.Seek(0, 2)

	rd := bufio.NewReader(f)

	//写入 Read Channel
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes: %s", err.Error()))
		}
		TypeMonitorChan <- TypeHandleLine
		rc <- line[:len(line)-1]
	}
}

func NewReadFromFile(path string) *ReadFromFile {
	return &ReadFromFile{
		path: path,
	}
}

// === WriteFromFile ===
type Writer interface {
	Write(wc chan *Message)
}

type WriteFromFile struct {
	influxDBsn string // influx data source
}

func (w *WriteFromFile) Write(wc chan *Message) {
	// 解析参数
	infSli := strings.Split(w.influxDBsn, "@")

	// create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     infSli[0],
		Username: infSli[1],
		Password: infSli[2],
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	for v := range wc {
		// create a new batch point
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  infSli[3],
			Precision: infSli[4],
		})
		if err != nil {
			log.Fatal(err)
		}

		// create a point and add to batch
		tags := map[string]string{
			"Path":   v.Path,
			"Method": v.Method,
			"Scheme": v.Scheme,
			"Status": v.Status,
		}
		fields := map[string]interface{}{
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
			"BytesSent":    v.BytesSent,
		}
		pt, err := client.NewPoint("nginx_log", tags, fields, v.TimeLocal)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		// write the batch
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

		log.Info("write success!")
	}

	//q := client.NewQuery("SELECT count(value) FROM cpu_load", "mydb", "")
	//if response, err := c.Query(q); err == nil && response.Error() == nil {
	//	fmt.Println(response.Results)
	//}
}

func NewWriteFromFile(influxDBsn string) *WriteFromFile {
	return &WriteFromFile{
		influxDBsn: influxDBsn,
	}
}

// === LogProcess ===
type LogProcess struct {
	rc    chan []byte
	wc    chan *Message
	read  Reader
	write Writer
}

type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

func (l *LogProcess) Process() {
	//从 Read Channel 中读取每行日志数据
	//正则提取所需的监控数据（path, status, method 等）
	//写入 Write Channel

	reg := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := reg.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			log.Warning("FindStringSubmatch fail: ", string(v))
			TypeMonitorChan <- TypeErrNum
			continue
		}

		msg := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			log.Warning("ParseInLocation fail: ", err)
			TypeMonitorChan <- TypeErrNum
			continue
		}

		msg.TimeLocal = t
		msg.BytesSent = cast.ToInt(ret[8])

		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			log.Warning("strings.Split fail", ret[6])
			TypeMonitorChan <- TypeErrNum
			continue
		}

		msg.Method = reqSli[0]

		u, err := url.Parse(reqSli[1])
		if err != nil {
			log.Warning("url.Parse fail", reqSli[1])
			TypeMonitorChan <- TypeErrNum
			continue
		}
		msg.Path = u.Path

		msg.Scheme = ret[5]

		msg.Status = ret[7]

		msg.UpstreamTime = cast.ToFloat64(ret[12])
		msg.RequestTime = cast.ToFloat64(ret[13])

		l.wc <- msg
	}
}

func NewLogProcess(r Reader, w Writer) *LogProcess {
	return &LogProcess{
		rc:    make(chan []byte, 200),
		wc:    make(chan *Message, 200),
		read:  r,
		write: w,
	}
}

// === Monitor ===
type SystemInfo struct {
	HandleLine int 	`json:"handleLine"`
	Tps float64 `json:"tps"`
	ReadCHanLen int `json:"readChanLen"`
	WriteChanLen int `json:"writeChanLen"`
	RunTime string `json:"runTime"`
	ErrNum int `json:"errNum"`
}

const (
	TypeHandleLine = 0
	TypeErrNum = 1
)

var TypeMonitorChan = make(chan int, 200)

type Monitor struct {
	startTime time.Time
	data SystemInfo
	tpsSli []int
}

func (m *Monitor) start(lp *LogProcess) {
	go func() {
		for tmc := range TypeMonitorChan {
			switch tmc {
			case TypeErrNum:
				m.data.ErrNum += 1
			case TypeHandleLine:
				m.data.HandleLine += 1
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	go func() {
		for {
			<- ticker.C
			m.tpsSli = append(m.tpsSli, m.data.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()

	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		m.data.RunTime = time.Now().Sub(m.startTime).String()
		m.data.ReadCHanLen = len(lp.rc)
		m.data.WriteChanLen = len(lp.wc)

		ret, _ := json.MarshalIndent(m.data, "", "\t")
		io.WriteString(writer, string(ret))
	})

	http.ListenAndServe(":9193", nil)
}

func main() {
	//var path, influxDsn string
	//flag.StringVar(&path, "path", "./access.log", "read file path")
	//flag.StringVar(&influxDsn, "influxDsn", "http://127.0.0.1:8086@imooc@imoocpass@imooc@s", "influx data source")
	//flag.Parse()
	//
	//// new..
	//r := NewReadFromFile(path)
	//w := NewWriteFromFile(influxDsn)
	//lp := NewLogProcess(r, w)
	r := NewReadFromFile("access.log")
	w := NewWriteFromFile("http://127.0.0.1:8086@imooc@imoocpass@imooc@s")
	lp := NewLogProcess(r, w)

	// TODO: 参数化
	go lp.read.Read(lp.rc)
	for i := 0; i < 2; i++ {
		go lp.Process()
	}
	for i := 0; i < 4; i++ {
		go lp.write.Write(lp.wc)
	}

	//time.Sleep(3000000 * time.Second)
	m := Monitor{
		startTime: time.Now(),
		data: SystemInfo{},
	}
	m.start(lp)
}
