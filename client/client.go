package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func StartClient(wg *sync.WaitGroup) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to server")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if strings.ToLower(message) == "exit"{
			wg.Done()
		}else{
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}
	}
}