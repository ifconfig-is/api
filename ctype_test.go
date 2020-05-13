package main

import (
	"fmt"
	"testing"
)

func TestGetCType(t *testing.T) {
	fmt.Println("Testing CType...")
	p := SetupParam("89.160.20.112")
	data, err := GetCType(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
