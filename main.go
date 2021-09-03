package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := ":" + arguments[1]
	l, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println("Error 1 =>: ", err.Error())
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error 2 =>: ", err.Error())
			return
		}
		go ConcurrentConnection(c)
	}
}

func someComputations(n int) int {
	values := make(map[int]int)
	for i := 0; i <= n; i++ {
		var value int
		switch i % 3 {
		case 0:
			value = 1
		case 1:
			value = values[i-1] + values[i-2]
		case 2:
			value = 10
		}
		values[i] = value
	}
	return values[n]
}

func ConcurrentConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error ConcurrentConnection =>: ", err.Error())
			os.Exit(100)
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		comp := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			comp = strconv.Itoa(someComputations(n)) + "\n"
		}
		c.Write([]byte(comp))
	}
	time.Sleep(time.Second * 5)
	c.Close()
}
