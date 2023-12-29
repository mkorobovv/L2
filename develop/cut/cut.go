package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cutConfig struct {
	flags    []string
	filepath string
	cmd      string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("cut utility")
	fmt.Println("--------------------------------")
	for {

		fmt.Print(">> ")
		buffer, _ := reader.ReadString('\n')

		length := len(strings.Split(buffer, " "))

		cmd := strings.Split(buffer, " ")[0]
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSuffix(cmd, "\r")

		if cmd == "q" {
			fmt.Print("end cut utility")
			os.Exit(1)
		}

		if length < 6 {
			fmt.Println("usage: cut -f list [-s] -d delim [file ...] ")
			continue
		}
	}
}

func (c *cutConfig) readf() (f *os.File, err error) {
	f, err = os.Open(c.filepath)
}
