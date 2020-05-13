package main

import (
	"fmt"
	"testing"
)

func TestGetEnterprise(t *testing.T) {
	fmt.Println("Testing Enterprise...")
	p := SetupParam("89.160.20.112")
	data, err := GetEnterprise(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
