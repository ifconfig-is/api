package main

import (
	"fmt"
	"testing"
)

func TestGetCountry(t *testing.T) {
	fmt.Println("Testing Country...")
	p := SetupParam("89.160.20.112")
	data, err := GetCountry(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}
