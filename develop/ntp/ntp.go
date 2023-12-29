package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func timeNTP() (time.Time, error) {
	response, err := ntp.Query("pool.ntp.org")
	time := time.Now().Add(response.ClockOffset)
	return time, err
}

func main() {
	time, err := timeNTP()
	if err != nil {
		fmt.Printf("bad response: %v\n", err)
		return
	}
	fmt.Println(time.String())
}
