package format

import "fmt"

func NewJsonString() string {
	return fmt.Sprintf("{\"time\": \"2024-04-11 23:46:47.2163963\", \"level\":\"INF\", \"caller\": \"./internal/service/tails.go:30\", \"msg\": \"hello word\"}\n")
}
