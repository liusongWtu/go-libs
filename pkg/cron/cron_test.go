package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestAddJob(t *testing.T) {
	AddJob("@every 3s",
		func() {
			fmt.Println(time.Now().String())
			time.Sleep(time.Second * 2)
		},
		SkipIfStillRunning)
	AddJob("@every 1s",
		func() {
			fmt.Println("RunIfStillRunning", time.Now().String())
			time.Sleep(time.Second * 2)
		},
		RunIfStillRunning)

	// AddJob("@every 1s",
	// 	func() {
	// 		fmt.Println("hi")
	// 		time.Sleep(time.Second * 2)
	// 	},
	// 	DelayIfStillRunning)

	time.Sleep(time.Hour)
}
