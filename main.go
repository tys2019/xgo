package main

import (
	"flag"
	"fmt"
)

var (
	configFile , pidFile	string
	v						bool
)

var (
	//VERSION
	VERSION 	string
	//BUILD
	BUILD 		string
	//COMMITSHA1 git commit ID
	COMMITSHA1	string
)

func parseArgs()  {
	flag.StringVar(&configFile,"svrConFile","./config/config","config file")
	flag.StringVar(&pidFile,"pidFile","/tmp/xgo.pid","pid file")
	flag.BoolVar(&v,"v",false,"show version")
	flag.Parse()
}
func usage()  {
	fmt.Printf(`Info:
		version : %s
		release time : %s
		commit sha1 : %s
	`,VERSION,BUILD,COMMITSHA1)
	fmt.Println("Usage:")
	flag.PrintDefaults()
}

func main()  {
	flag.Parse()
	if v{
		usage()
		return
	}
	
	reload()
}

func reload()  {
	
}


