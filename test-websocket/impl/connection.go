package impl

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	closeChan chan []byte
	mutex sync.Mutex
	isClosed bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn: wsConn,
		inChan: make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),
	}

	// 读 chan 协程
	conn.readLoop()

	// 写 chan 协程
	conn.writeLoop()

	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	data = <-conn.inChan
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	conn.outChan <- data
	return
}

func (conn *Connection) Close() {
	// 线程安全，可重入的 Close
	conn.wsConn.Close()

	// 加 mutex 锁保证线程安全，readLoop 和 writeLoop 都可能调用本方法导致 close(conn.closeChan) 出现多次调用而出错（closeChan 只能 close 一次）
	// 这里确保 closeChan 只 close 一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var data []byte
	var err error

	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			// todo: goto err
			return
		}

		// 当 writeLoop 中出错关闭 conn，这里会因 conn.inChan 满等待读出而仍处于阻塞状态...
		//conn.inChan <- data
		select {
		case conn.inChan <- data:
		case <- conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var data []byte
	var err error

	for {
		// 当 readLoop 中出错关闭 conn，这里会因 conn.outChan 空等待写入而仍处于阻塞状态...
		//data = <- conn.outChan
		select {
		case data = <- conn.outChan:
		case <- conn.closeChan:
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}