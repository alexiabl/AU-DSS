package main

import("fmt";"net";"bufio";"os")

//channel
var channel = make(chan string)
var connections = []net.Conn{}
 
//broadcast of messages
func broadcast(){
	msg := <-channel
	for _,conn := range connections {
		fmt.Fprintf(conn, msg)
	}
}


func handleConn(connection net.Conn){
	fmt.Println("Connected to "+ connection.RemoteAddr().String())
	for {
		msg, error := bufio.NewReader(connection).ReadString('\n')
		if (error !=nil) { 
			fmt.Println("ERROR")
			// remove connection
			return
		}else{
			msg_from := "["+connection.RemoteAddr().String()+"]: "+msg
			channel <- msg_from
		}
	}
}

func main(){

	fmt.Println("Server started...")
	name, _ := os.Hostname()
	fmt.Println("Server name is: ",name)
	listener, _ := net.Listen("tcp",":4444")


	for{
		go broadcast()
		conn, _ := listener.Accept()
		_, port, _ := net.SplitHostPort(listener.Addr().String())
		fmt.Println("Listening on port " + port)
		connections = append(connections,conn)
		go handleConn(conn)
	}
}

/*
 When the client is started it prompts the user for and IP address and port
number. Then it makes a TCP connection to the corresponding IP address
and port. Then it will prompt the user for input until quit is typed. The input
is sent to the server. Concurrently it reads lines of text coming from
the server and prints them. Notice that several Go routines may access
Conn.Read and Conn.Write concurrently.

 When the server is started it determines its local IP address and prints it.
Then it start listening for connections on a random open port and prints
the port number. It should be able to handle several concurrent connections.
When the server receives a strong on any incoming connections, it will
prepend the IP address and port number of the sender to the string and send
the string back to all its current connections.*/