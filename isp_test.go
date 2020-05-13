package main

import (
	"fmt"
	"testing"
)

func TestGetISP(t *testing.T) {
	fmt.Println("Testing ISP...")
	p := SetupParam("89.160.20.112")
	data, err := GetISP(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
