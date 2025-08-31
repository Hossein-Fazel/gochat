package main
import(
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"runtime"
	"encoding/json"
	"gochat/client"
	"gochat/server"
)

type Config struct{
	Your_erver string `json:"your_server"`
	Your_port int `json:"your_port"`
	Friend_server string `json:"friend_server"`
	Friend_port int `json:"friend_port"`
}

func clearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error clearing screen:", err)
    }
}

func main(){
	clearScreen()
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

	go server.StartServer(config.Your_erver, config.Your_port)
	client.StartClient(config.Friend_server, config.Friend_port, strings.TrimSpace(name))

	fmt.Println("\t\t*** Closing the app ***")
	fmt.Println("\t\t    *** Goodbye ***")
}