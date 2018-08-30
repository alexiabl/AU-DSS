package main
import ( "net" ; "time" ; "strconv" )

func main() {
	ServerAddr, _ := net.ResolveUDPAddr("udp","127.0.0.1:8080")
	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)
	defer conn.Close()
	i := 0

	for {
		i++
		msg := strconv.Itoa(i)
		conn.Write([]byte(msg))
		time.Sleep(time.Second * 1)
	}
}