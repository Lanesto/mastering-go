package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

func main() {
	logger, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, os.Args[0])
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(logger)
	}
	log.Panic(fmt.Sprintf("%#v", logger))
	fmt.Println("You can't see this!")
}
