package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	logger, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(logger)
	}

	log.Fatal(fmt.Sprintf("%#v", logger))
	fmt.Println("You will never see this!")
}
