package main

import (
	"go-test/s3/jcloud"
	"log"
)

func main() {
	jcloud.JCloudS3.Init()

	if data, err := jcloud.JCloudS3.Get("form/12935/uploader/d7792554-2c1f-4cf4-83bb-23117736f03e.xlsx"); err != nil {
		log.Printf("出错: %s", err.Error())
	} else {
		log.Printf("找到: %d", len(data))
	}

}
