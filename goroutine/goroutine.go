package goroutine

import (
	"fmt"
	"time"
)

func print(task string) {
	fmt.Println(task)
}

func Goroutine() {
	tasks := []string{
		"a",
		"b",
		"c",
	}

	for _, task := range tasks {
		go print(task) // a, b, cが出力される
		go func() {
			fmt.Println(task) // 全てcが出力される。goroutineが起動する時にはすでにループが周り切っていて、taskは最後のcになっている
		}()
	}

	time.Sleep(time.Second)
}
