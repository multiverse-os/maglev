package job

import (
	"fmt"
)

type Job struct {
	Unique bool
	Action func()
}

func PreformSCheduledTask() {
	fmt.Println("vim-go")
}
