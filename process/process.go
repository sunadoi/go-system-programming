package process

import (
	"fmt"
	"os"
)

func Process() {
	fmt.Printf("プロセスID: %d\n", os.Getegid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())
}
