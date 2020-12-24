package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	logger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(logger)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	logger, err = syslog.New(syslog.LOG_MAIL, "Some Program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(logger)
	}
	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}
