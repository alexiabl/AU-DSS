package main

import("fmt";"net";"bufio";"os")

func main(){

	fmt.Println("Server started...")

	listener, _ := net.Listen("tcp",":4444")

	name, _ := os.Hostname()
	addrs, _ := net.LookupHost(name)
	fmt.Println("Server address is: ",addrs)
	conn, _ := listener.Accept()


	for{
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		if msg == "quit\n" { return }
		fmt.Print("Message received - ", string(msg))
		conn.Write([]byte(msg))
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