package main

import (
	"fmt"
	"os"
)

func main()  {
	f, err := os.OpenFile("./access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err: ", err)
	}
	defer f.Close()

	f.WriteString(`172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854 >> access.log`+"\n")
}