package main

import (
	"fmt"
	"testing"
)

func TestGetCity(t *testing.T) {
	fmt.Println("Testing City...")
	p := SetupParam("89.160.20.112")
	data, err := GetCity(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
