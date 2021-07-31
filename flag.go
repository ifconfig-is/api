package main

import (
	"flag"
	"os"
)

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

func init_flag() {
	flag.BoolVar(&DEV, "dev", false, "development mode")
	flag.Parse()
}
