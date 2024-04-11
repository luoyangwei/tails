package format

import (
	"fmt"
	"time"
)

func NewConsoleString() string {
	now := time.Now().Format("2006-01-02 15:04:05.0000000")
	caller := "./internal/service/tails.go:30"
	hw := "hello word"
	return fmt.Sprintf("%-28s %c[;32m%-6s%c[0m%-34s%s\n", now, 0x1B, "INF", 0x1B, caller, hw)
}
