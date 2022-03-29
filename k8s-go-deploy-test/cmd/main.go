package main

import (
	"flag"
	"fmt"
	"kgdt/inits"
	"kgdt/internal/controller"
	"os"
)

var (
	BuildTime string
	CommitID  string
	GitTag    string
)

func main() {
	showVer := flag.Bool("version", false, "show version")
	flag.Parse()
	if *showVer {
		fmt.Printf("%s\t%s\t%s\n", GitTag, CommitID, BuildTime)
		os.Exit(0)
	}
	controller.InitWeb(inits.Configs.Project)

}
