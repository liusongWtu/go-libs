package gox

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// 异步任务worker
// 可持续运行,自我监控，异常退出后可自启
// 通过chan传递数据
// 可指定时间周期性执行任务(如间隔一定时间保存一次数据)

var (
	defaultWorkerOptions = WorkerOptions{
		CronDuration: time.Hour * 24,
		AlertFunc: func(title, detail string) {
			logx.Errorf("%s, %s", title, detail)
		},
	}
)

type WorkerOptions struct {
	FinishFunc   func()
	CronDuration time.Duration
	CronFunc     func()
	AlertFunc    func(title, detail string)
}

type WorkerOption func(*WorkerOptions)

func WithFinishFunc(finishFunc func()) WorkerOption {
	return func(o *WorkerOptions) {
		o.FinishFunc = finishFunc
	}
}

func WithCronFunc(cronDuration time.Duration, cronFunc func()) WorkerOption {
	return func(o *WorkerOptions) {
		o.CronDuration = cronDuration
		o.CronFunc = cronFunc
	}
}

func WithAlertFunc(alert func(title, detail string)) WorkerOption {
	return func(o *WorkerOptions) {
		o.AlertFunc = alert
	}
}

func NewPersistentChanWorker[T any](name string, dataChan chan T, handleItemFunc func(T), opts ...WorkerOption) *PersistentChanWorker[T] {
	options := defaultWorkerOptions
	for _, o := range opts {
		o(&options)
	}

	worker := &PersistentChanWorker[T]{
		name:              name,
		DataChan:          dataChan,
		keepAliveChan:     make(chan struct{}),
		shutdownSignal:    make(chan os.Signal, 1),
		isShutdown:        false,
		stopKeepAliveChan: make(chan struct{}),
		handleItemFunc:    handleItemFunc,
		finishFunc:        options.FinishFunc,
		cronDuration:      options.CronDuration,
		cronFunc:          options.CronFunc,
		alertFunc:         options.AlertFunc,
	}
	return worker
}

type PersistentChanWorker[T any] struct {
	name string
	//数据chan
	DataChan chan T
	//保活，存储协程意外退出
	keepAliveChan  chan struct{}
	shutdownSignal chan os.Signal
	//程序收到退出信号
	isShutdown bool
	//停止保活
	stopKeepAliveChan chan struct{}
	handleItemFunc    func(T)
	finishFunc        func()
	cronDuration      time.Duration
	cronFunc          func()
	alertFunc         func(title, detail string)
}

func (w *PersistentChanWorker[T]) alert(title, detail string) {
	if w.alertFunc != nil {
		w.alertFunc(title, detail)
	}
}

func (w *PersistentChanWorker[T]) run() {
	go func() {
		ticker := time.NewTicker(w.cronDuration)
		defer func() {
			ticker.Stop()
			if err := recover(); err != nil {
				w.alert(fmt.Sprintf("%s worker run recover", w.name), fmt.Sprintf("%s PersistentChanWorker run recover error:%+v,msg:%s", w.name, err, GetStack()))
			}
			if !w.isShutdown {
				time.Sleep(time.Second)
				w.keepAliveChan <- struct{}{}
			}
		}()

		for {
			select {
			case info := <-w.DataChan:
				w.handleItemFunc(info)
			case <-ticker.C:
				if w.cronFunc != nil {
					w.cronFunc()
				}
			case <-w.shutdownSignal:
				w.finishFunc()
				w.isShutdown = true
				w.stopKeepAliveChan <- struct{}{}
				return
			}
		}
	}()
}

func (w *PersistentChanWorker[T]) AddItem(info T) {
	w.DataChan <- info
}

func (w *PersistentChanWorker[T]) Start() {
	signal.Notify(w.shutdownSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	w.keepAlive()
	w.run()
}

func (w *PersistentChanWorker[T]) keepAlive() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				w.alert(fmt.Sprintf("%s worker run recover", w.name),
					fmt.Sprintf("%s PersistentChanWorker keepAlive recover error:%+v,msg:%s", w.name, err, GetStack()))
			}
		}()

		for {
			select {
			case <-w.keepAliveChan:
				w.run()
			case <-w.stopKeepAliveChan:
				w.alert(fmt.Sprintf("%s worker run closed", w.name),
					fmt.Sprintf("%s PersistentChanWorker keepAlive exit", w.name))
				return
			}
		}
	}()
}

// 主动停止
func (w *PersistentChanWorker[T]) Stop() {
	w.shutdownSignal <- syscall.SIGHUP
}
