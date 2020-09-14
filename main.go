package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/wakatakeru/pubpass/cmd"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := cmd.ApplyCmd.Execute(); err != nil {
		glog.Fatal("Command Error")
	}
}
