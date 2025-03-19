package main
import(
	"fmt"
	"sync"
	"os"
	"bufio"
	"strings"
	"encoding/json"
	"gochat/client"
	"gochat/server"
)

type Config struct{
	Server string `json:"server"`
	Your_port int `json:"your_port"`
	Friend_port int `json:"friend_port"`
}

func main(){
	var wg sync.WaitGroup
	var config Config
	data, err := os.ReadFile("config.json")
	if err != nil{
		fmt.Println("Config file does not exist")
		os.Exit(0)
	}

	err = json.Unmarshal(data, &config)

	if err != nil{
		fmt.Println("Config file is corrupt")
		os.Exit(0)
	}

	fmt.Println("\t\t*** Welcome to gochat ***")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")
	name, _ := reader.ReadString('\n')

	wg.Add(1)
	go server.StartServer(config.Server, config.Your_port)
	go client.StartClient(&wg, config.Server, config.Friend_port, strings.TrimSpace(name))
	wg.Wait()

	fmt.Println("\t\t*** Closing the app ***")
	fmt.Println("\t\t    *** Goodbye ***")
}