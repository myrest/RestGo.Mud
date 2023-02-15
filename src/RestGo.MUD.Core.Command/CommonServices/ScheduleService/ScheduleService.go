package ScheduleService

import (
	"fmt"
	"time"
)

// Timer represents a timer task
type scheduler struct {
	interval     time.Duration //多久執行一次
	task         func()        //要執行的函式
	stop         chan struct{} //執行結束後，是否要停止Job
	isRunning    bool          //Job是否在執行中
	endExecution chan struct{} //執行結束後要發出一個訊號，可以直接停止Job
	isStopping   bool          //是否正在停止中，限制Stop只能被執行一次
}

// var jobs = []*scheduler{}
var jobs = make(map[string]*scheduler)

// NewTimer returns a new Timer and starts the timer task
func NewTimerInMins(jobName string, mins int, task func()) {
	if mins == 0 {
		panic(fmt.Sprintf("[NewTimerInMins]Job mins 為 0: [%s]", jobName))
	}
	duration := time.Duration(mins) * time.Minute
	NewTimer(jobName, duration, task)
}

// NewTimer returns a new Timer and starts the timer task
func NewTimerInHours(jobName string, hour int, task func()) {
	duration := time.Duration(hour) * time.Hour
	NewTimer(jobName, duration, task)
}

func NewTimer(jobName string, interval time.Duration, task func()) {
	if interval == 0 {
		panic(fmt.Sprintf("[NewTimer]傳入的 interval 為 0: [%s]", jobName))
	}

	if _, ok := jobs[jobName]; ok {
		panic(fmt.Sprintf("Job Name重覆定義[%s]", jobName))
	}

	t := &scheduler{
		interval:   interval,            // No changes here
		task:       task,                // No changes here
		stop:       make(chan struct{}), // No changes here
		isRunning:  false,               // No changes here
		isStopping: false,               // No changes here
	}

	if t.interval == 0 {
		panic(fmt.Sprintf("[NewTimer]t.interval 為 0: [%s]", jobName))
	}

	jobs[jobName] = t // Moved this line to the end of the function to avoid panic if job name already exists
	go t.start()      // Moved this line to the end of the function to ensure that jobs are added before starting them
}

func (t *scheduler) start() {
	for {
		select {
		case <-time.After(t.interval):
			t.isRunning = true
			t.endExecution = make(chan struct{})
			t.task()
			t.isRunning = false
			close(t.endExecution)
		case <-t.stop:
			return
		}
	}
}

func (t *scheduler) Stop() {
	if !t.isStopping {
		t.isStopping = true
		if t.isRunning {
			select {
			case <-t.stop:
				return
			default:
				if _, isChannelOPen := <-t.endExecution; !isChannelOPen {
					close(t.stop)
				}
			}
		} else {
			close(t.stop)
		}
	}
}

func Shutdown() {
	for jobname, j := range jobs {
		j.Stop()
		fmt.Println(jobname, " is stopped.")
	}
}
