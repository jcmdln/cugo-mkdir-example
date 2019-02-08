package main

import (
	"flag"

	"github.com/jcmdln/cugo/src/mkdir"
)

var opt = &mkdir.Options{}

func init() {
	flag.UintVar(&opt.Mode, "m", 0755, "")
	flag.BoolVar(&opt.Parents, "p", false, "Make parent directories")
	flag.BoolVar(&opt.Verbose, "v", false, "Show verbose output")
	flag.Parse()
}

func main() {
	opt.Mkdir(flag.Args())
}
