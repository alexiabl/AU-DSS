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
	fmt.Println("Enter peer's IP Address")
	reader := bufio.NewReader(os.Stdin)
	ip, _ := reader.ReadString('\n')

	fmt.Println("Enter peer's port")
	port, _ := reader.ReadString('\n')

	address := strings.Join(ip,":")
	final_addr := strings.Join(address,port)

	fmt.Println("The address is: "+final_addr)
	//how to append a string from user input without the \n

	//check if peer exists in the network

	//
	createConnection(final_addr)


}