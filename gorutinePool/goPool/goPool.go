package goPool

import "fmt"

type Task struct {
	task func() error
}

func (t *Task) NewTask(f func() error) {
	t.task = f
}
func (t *Task) Execute() {
	t.task()
}

type Pool struct {
	EntryChannel chan Task
	JobsChannel  chan Task
	WorkNum      int
}

func (p *Pool) NewPool(workNum int) {
	p.WorkNum = workNum
	p.EntryChannel = make(chan Task)
	p.JobsChannel = make(chan Task)
}

//取出任务工作
func (p *Pool) Worker(workId int) {
	workNum := 0
	for task := range p.JobsChannel {
		task.Execute()
		workNum += 1
		fmt.Println("work id : ", workId)
		fmt.Println("执行任务:")
	}
}
func (p *Pool) Run() {
	for i := 0; i < p.WorkNum; i++ {
		go p.Worker(i)
	}
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}
