package logger

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func Log(test string) {
	fmt.Println(aurora.Cyan(test))
}
