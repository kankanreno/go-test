package main

import (
	"fmt"
	"time"
)

// --- task ---
// 定义一个任务类型 Task
type Task struct {
	f func() error // biz
}

// 创建一个 Task 任务
func NewTask(argF func() error) *Task {
	return &Task{
		f: argF,
	}
}

// 执行任务
func (t *Task) Execute() {
	t.f()
}

// --- pool ---
// 定义一个协程池类型 Pool
type Pool struct {
	JobsChannel  chan *Task // 对内任务出口
	EntryChannel chan *Task // 对外任务入口

	workerNum int // 协程数量
}

// 创建 Pool
func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    cap,
	}
}

// 不断地 从 JobsChannel 取出一个 task，并干活
func (p *Pool) worker(ID int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("worker", ID, "执行完了一个任务")
	}
}

// 根据协程数量，创建 worker 协程
func (p *Pool) run1() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
}

// 不断地 从 EntryChannel 中取出 task，并发送给 JobsChannel
func (p *Pool) run2() {
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

// --- main ---
func biz() error {
	fmt.Println(time.Now())
	return nil
}

func main() {
	t := NewTask(biz)

	p := NewPool(4)

	p.run1()

	// 不断地 放入任务
	taskNum := 0
	go func() {
		for {
			p.EntryChannel <- t
			taskNum += 1
			fmt.Println("当前一共放入了 ", taskNum, " 个任务")
		}
	}()

	p.run2()
}
