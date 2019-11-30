package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var group sync.WaitGroup

func main() {
	//设置利用核心数量
	//runtime.GOMAXPROCS(8) //3.1453302
	runtime.GOMAXPROCS(1) //

	start := time.Now()
	for i := 0; i < 10000; i++ {
		go costTimeTask()
	}
	//同步等待
	group.Wait()
	end := time.Now()
	//输出执行时间，单位为秒。
	fmt.Println(end.Sub(start).Seconds())
}

func costTimeTask() {
	fmt.Println("xxxxxxxxxxx")
	time.Sleep(time.Duration(2) * time.Second)

}
