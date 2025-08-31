package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Hossein-Fazel/Gobar/progressbar"
)

func StartClient(server string, port int, name string) {
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
		message = strings.TrimSpace(message)
		if strings.ToLower(message) == "exit" {
			return
		} else if len(message) >= 6 && message[:6] == "<SEND " {
			fp := message[6 : len(message)-1]
			finfo, err := os.Stat(fp)
			if err != nil {
				fmt.Println(err)
				continue
			}
			send_file(finfo, fp, conn)
		} else {
			_, err := conn.Write([]byte(fmt.Sprintf("%v : %v\n", name, message)))
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}
	}
}

func send_file(file_info os.FileInfo, file_path string, conn net.Conn) {
	var header string = fmt.Sprintf("<RECEIVE:%v:%v>\n", file_info.Name(), file_info.Size())

	_, err := conn.Write([]byte(header))
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}
	file, _ := os.Open(file_path)
	defer file.Close()
	buffer := make([]byte, 4096)
	pbar := progressbar.NewProgressBar()
	pbar.Set_filled("-")
	pbar.Set_total(int(file_info.Size()))
	for {
		count, err := file.Read(buffer)
		if err != nil {
			if count == 0 {
				break
			}
			fmt.Println("An error occurred")
			return
		}

		_, err = conn.Write(buffer[:count])
		pbar.Update(count)
		pbar.Show()
		if err != nil {
			fmt.Println("Error sending file:", err)
			return
		}
	}
	pbar.Stop()
	fmt.Printf("File '%v' sent successfully.\n", file_info.Name())
}
