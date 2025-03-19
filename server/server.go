package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		message = strings.TrimSpace(message)
		if len(message) >= 9 && message[:9] == "<RECEIVE:"{
			parts := strings.Split(message, ":")
		
			file, err := os.Create(parts[1])
			if err != nil{
				fmt.Printf("An error occurred: %v\n", err)
			}
			defer file.Close()
			
			fmt.Printf("Receiving file '%v' (%v bytes)\n", parts[1], parts[2][:len(parts[2])-1])
			buffer := make([]byte, 4096)
			for {
				count, err := reader.Read(buffer)
				if err != nil{
					fmt.Println("An error occurred")
					return
				}
				if string(buffer[:count]) == "<END>"{
					break
				}

				file.Write(buffer[:count])
			}
			fmt.Printf("File '%v' received successfully.\n", parts[1])

		}else{
			fmt.Printf("%s\n", message)
		}
	}
}

func StartServer(server string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", server, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Server started on %v:%v\n", server, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}