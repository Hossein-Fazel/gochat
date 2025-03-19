package main
import(
	"fmt"
	"sync"
	"gochat/client"
	"gochat/server"
)


func main(){
	var wg sync.WaitGroup

	fmt.Println("\t\t*** Welcome to gochat ***")
	wg.Add(1)
	go server.StartServer()
	go client.StartClient(&wg)
	wg.Wait()

	fmt.Println("\t\t*** Closing the app ***")
	fmt.Println("\t\t    *** Goodbye ***")
}