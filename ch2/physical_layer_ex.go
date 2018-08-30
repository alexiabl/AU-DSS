package main
import ( "net" ; "os" ; "fmt" ; "strconv" )

func main() {
	name, _ := os.Hostname()
	addrs, _ := net.LookupHost(name)
	fmt.Println("Name: " + name)
	fmt.Println("address 1")
	for indx, addr := range addrs {
		fmt.Println("Address number " + strconv.Itoa(indx) + ": " + addr)
	}
}