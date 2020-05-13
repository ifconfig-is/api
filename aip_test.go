package main

import (
	"fmt"
	"testing"
)

func TestGetAIP(t *testing.T) {
	fmt.Println("Testing AIP...")
	p := SetupParam("89.160.20.112")
	data, err := GetAIP(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
