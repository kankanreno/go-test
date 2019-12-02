package main

import (
	"fmt"
	"github.com/chenhg5/go-task"
	"runtime"
	"strconv"
	"time"
)

func main() {

	// init
	task.InitTaskReceiver(runtime.NumCPU())

	// 有十个人同时点菜
	for i := 0; i < 10; i++ {
		task.AddTask(task.NewTask(
			map[string]interface{}{
				"user": strconv.Itoa(i),
			},                                              // 参数
			[]task.FacFunc{ordering, cooking, deliverying}, // 任务列表
			-1), // -1代表任务不超时
		)
	}

	time.Sleep(time.Second * 50)
}

// 下单任务
func ordering(uuid string, param map[string]interface{}) (string, error) {
	fmt.Println("user " + param["user"].(string) + " is ordering")
	time.Sleep(time.Second * 1)
	return uuid, nil
}

// 做菜任务
func cooking(uuid string, param map[string]interface{}) (string, error) {
	fmt.Println("user " + param["user"].(string) + " is cooking")
	time.Sleep(time.Second * 1)
	return uuid, nil
}

// 配送任务
func deliverying(uuid string, param map[string]interface{}) (string, error) {
	fmt.Println("user " + param["user"].(string) + " is deliverying")
	time.Sleep(time.Second * 1)
	return uuid, nil
}
