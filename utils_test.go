package dbtoapi

import (
	"fmt"
	"testing"
)

func TestRespJson(t *testing.T) {
	println("******************test respJson start******************")

	resp := respJson([]string{"a", "b", "c"})
	fmt.Println(resp)
}
