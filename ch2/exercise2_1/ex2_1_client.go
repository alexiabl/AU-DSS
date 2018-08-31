package main

import ( "bufio" ; "fmt" ; "net"; "os"; "strings")

func sendMessage(conn net.Conn){
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if text == "quit\n" { return }
	fmt.Fprintf(conn,text)
}


//go routine for receiving messages

func receiveMessage(conn net.Conn){
	for {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil { 
		fmt.Println("SERVER ERROR")
		return 
	}
	fmt.Println(msg)
	}
}

func main(){
	fmt.Println("Enter IP Address:")
	reader := bufio.NewReader(os.Stdin)
	ip, _ := reader.ReadString('\n')

	fmt.Println("Enter Port:")
	port, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip,"\n")
	port = strings.TrimSuffix(port,"\n")

	final_addr := net.JoinHostPort(ip,port)

	fmt.Println("The address will be: ",final_addr)
	conn, _ := net.Dial("tcp",final_addr)
	
	defer conn.Close()

	go receiveMessage(conn)

	for{
		sendMessage(conn)
	}

}

