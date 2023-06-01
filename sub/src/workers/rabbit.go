package workers

import (
	"fmt"
	"time"
)

type fn func()

func New(routine fn, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})
	func() {
		for {
			select {
			case <-ticker.C:
				routine()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func Counter() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
}
