package main

import ( "bufio" ; "fmt" ; "os" ; "strings";"net")

//must use a mutex with map since it is not thread safe.

type peer struct {
	ip_address string
	port int
	//map var MessagesSent map[string]bool
}

func createConnection(address string){
	listener, _ := net.Listen("tcp",address)
	connection , _ := listener.Accept()

}

func main(){
	// initialize peers network to empty
	peers_network := []peer{}

	fmt.Println("Hello! Welcome to my toy P2P Network\n")
	fmt.Println("Enter peer's IP Address and port")
	reader := bufio.NewReader(os.Stdin)
	address, _ := reader.ReadString('\n')

	listener 



	//how to append a string from user input without the \n

	//check if peer exists in the network

	//
	createConnection(final_addr)


}