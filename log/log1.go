package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {

    filename := os.Args[0]+".log"
    file,err := os.Create(filename)
	if err != nil {
		panic("Failed to create log file")
	}

	logger := log.New(file,"",log.LstdFlags | log.Llongfile) 
	logger.Println("1.Println log with log.LstdFlags ...")
	logger.Println("2.Println log with log.LstdFlags ...")

	logger.SetFlags(log.LstdFlags)

	log.Println("3.Println log without log.LstdFlags ...")
	logger.Println("3.Println log without log.LstdFlags ...")

	return 0 
}


