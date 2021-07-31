package main

import (
	"flag"
	"os"
)

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

func init_flag() {
	flag.BoolVar(&PROD, "prod", false, "production mode")
	flag.Parse()
}
