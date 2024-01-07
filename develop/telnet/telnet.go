package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type connection struct {
	addr    string
	socket  net.Conn
	timeout string
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("- telnet command shell -")

	for {
		fmt.Print("cmd> ")
		buf, _ := reader.ReadString('\n')

		splited := strings.Split(buf, " ")
		cmd := splited[0]
		if cmd == "quit" {
			os.Exit(1)
		}

		if cmd != "go-telnet" {
			fmt.Println("usage: go-telnet [--timeout=time.Second] host port")
			continue
		}

		if len(splited) < 2 && len(splited) > 3 {
			fmt.Println("usage: go-telnet [--timeout=time.Second] host port")
			continue
		}

		host := splited[1]
		port := splited[2]

		var timeout = ""

		if len(splited) == 4 {
			timeout = splited[1]
			host = splited[2]
			port = splited[3]
		}

		port = strings.TrimSuffix(port, "\n")
		port = strings.TrimSuffix(port, "\r")

		fmt.Println(host + ":" + port)

		connector := &connection{addr: host + ":" + port, timeout: timeout}
		connector.run()
	}

}

func (c connection) run() {

	log.Println("starting connection to", c.addr)

	timeout := 0 * time.Second
	if c.timeout != "" {
		re, _ := regexp.Compile("[0-9]+")
		parsedTimeout, _ := strconv.Atoi(re.FindAllString(c.timeout, 1)[0])
		timeout = time.Duration(parsedTimeout) * time.Second
	}

	conn, errConn := net.DialTimeout("tcp", c.addr, timeout)
	if errConn != nil {
		log.Fatal("error at read stdin: ", errConn)
	}
	c.socket = conn
	log.Println("successfully connected to socket:")

	buf := make([]byte, 8192)
	for {
		fmt.Print("telnet>")
		recv, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal("error at read stdin:", err)
		}
		_, errWrite := c.socket.Write(buf[:recv])
		if errWrite != nil {
			log.Fatal("error at write socket:", errWrite)
		}
	}
}

func read(c connection) {
	buf := make([]byte, 8192)
	for {

		recv, err := c.socket.Read(buf)
		if err != nil {
			log.Fatal("error at read socket:", err)
		}
		fmt.Println("recieved -> ", string(buf[:recv]))
	}
}
