package api

import (
	"fmt"
	"testing"
)

func Test_calculateTokens(t *testing.T) {
	got := calculateTokens("hello world")
	fmt.Printf("got:%v\n", got)

	got = calculateTokens("世界你好")
	fmt.Printf("got:%v\n", got)
}
