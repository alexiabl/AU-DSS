package main

import ( "bufio" ; "fmt" ; "net"; "os")


func main(){
	fmt.Println("Enter IP Address and port ")
	reader := bufio.NewReader(os.Stdin)
	address, _ := reader.ReadString('\n')

	fmt.Println("The address will be: ",address)
	conn, _ := net.Dial("tcp",address)
	if (conn!=nil) { defer conn.Close() }

	for{
		    // read in input from stdin
			fmt.Print("Text to send: ")
			text, err := reader.ReadString('\n')
			if text == "quit\n" { return }
			fmt.Fprintf(conn, text)
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil { return }
			fmt.Print("From server: " + msg)
	}
}

