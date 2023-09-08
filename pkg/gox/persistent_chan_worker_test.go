package gox

import (
	"fmt"
	"testing"
	"time"
)

type WorkerMock struct {
	dataChan    chan int
	TaskManager *PersistentChanWorker[int]
	BatchPool   []int
	BatchSize   int
}

func (m *WorkerMock) handleItem(info int) {
	m.BatchPool = append(m.BatchPool, info)
	if len(m.BatchPool) < m.BatchSize {
		return
	}
	m.batchSave()
}

func (m *WorkerMock) batchSave() {
	if len(m.BatchPool) == 0 {
		return
	}
	fmt.Printf("batchSave data:%+v\n", m.BatchPool)
	m.BatchPool = m.BatchPool[:0]
}

func (m *WorkerMock) finish() {
	fmt.Printf("finish data:%+v\n", m.BatchPool)
}

func TestPersistentChanWorker(t *testing.T) {
	workerMock := WorkerMock{
		dataChan:  make(chan int, 100),
		BatchPool: make([]int, 0, 5),
		BatchSize: 5,
	}
	workerMock.TaskManager = NewPersistentChanWorker[int](
		"mock",
		workerMock.dataChan,
		workerMock.handleItem,
		WithCronFunc(time.Second*6, workerMock.batchSave),
		WithFinishFunc(workerMock.finish),
	)
	workerMock.TaskManager.Start()

	for i := 0; i < 1000; i++ {
		workerMock.dataChan <- i
		time.Sleep(time.Second)
	}
}
