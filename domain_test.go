package main

import (
	"fmt"
	"testing"
)

func TestGetDomain(t *testing.T) {
	fmt.Println("Testing Domain...")
	p := SetupParam("89.160.20.112")
	data, err := GetDomain(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
