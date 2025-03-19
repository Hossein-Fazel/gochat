package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func StartClient(wg *sync.WaitGroup, server string, port int, name string) {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", fmt.Sprintf("%v:%v", server, port))
		if err == nil {
			fmt.Println("Your friend is online")
			break
		}
		fmt.Println("Your friend is offline")
		time.Sleep(2 * time.Second)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if strings.ToLower(message) == "exit"{
			wg.Done()
		}else{
			_, err := conn.Write([]byte(fmt.Sprintf("%v : ", name) + message + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}
	}
}