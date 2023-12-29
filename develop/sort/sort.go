package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdconfig struct {
	command  string
	filepath string
	flags    []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Starting sort util")
	fmt.Println("--------------------------------")

	for {
		fmt.Print(">> ")

		buffer, _ := reader.ReadString('\n')
		length := len(strings.Split(buffer, " "))

		cmd := strings.Split(buffer, " ")[0]
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSuffix(cmd, "\r")

		if cmd == "q" {
			fmt.Print("quit from sort shell")
			os.Exit(1)
		}

		if length < 2 {
			fmt.Println("error: sort command needs at least two arguments")
			continue
		}

		filepath := strings.Split(buffer, " ")[length-1]
		filepath = strings.TrimSuffix(filepath, "\n")
		filepath = strings.TrimSuffix(filepath, "\r")

		cmdcfg := &cmdconfig{}
		cmdcfg.filepath = filepath
		cmdcfg.command = cmd

		for _, flag := range strings.Split(buffer, " ")[1 : length-1] {
			flag = strings.TrimSuffix(flag, "\n")
			flag = strings.TrimSuffix(flag, "\r")
			cmdcfg.flags = append(cmdcfg.flags, flag)
		}

	}
}

func (c *cmdconfig) readf() ([]string, error) {
	file, err := os.Open(c.filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func (c *cmdconfig) writef() error {
	file, err := os.OpenFile(c.filepath, os.O_WRONLY|os.O_CREATE, 0666)

	return err
}

func (c *cmdconfig) parseInput() {

	if c.command != "sort" {
		fmt.Println("wrong command")
		return
	}

	for i, _ := range c.flags {
		switch c.flags[i] {
		case "-k":
			column := c.flags[i + 1]
			ncol, err := strconv.Atoi(column)
			
			if err != nil {
				fmt.Println("cannot convert char to int")
				return
			}
			
		}
	}
}
