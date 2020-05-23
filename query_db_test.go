package dbtoapi

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	println("******************test queryDB start******************")

	data := QueryDB("table", 1, 1)
	fmt.Println(data)
}
