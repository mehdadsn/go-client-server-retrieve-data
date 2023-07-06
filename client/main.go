package main

import (
	"fmt"
	"net"
)

func main() {
	var input string
	function := ""
	param := ""
	for {
		fmt.Println("\n----NEW REQUEST----\n1-Get Name By Id\n2-Get In Stock\n3-Order\n4-Get By category\nEnter the number")
		fmt.Scan(&input)
		switch input {
		case "1":
			function = "getnamebyid"
			fmt.Println("enter id: ")
			fmt.Scan(&param)
		case "2":
			function = "getinstock"
		case "3":
			function = "order"
			fmt.Println("enter id: ")
			fmt.Scan(&param)
		case "4":
			function = "getstockbycategory"
			fmt.Println("enter category name: ")
			fmt.Scan(&param)
		}
		request := function + "\n" + param
		connect(request)

	}

}

func connect(request string) {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	_, err = connection.Write([]byte(request))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println(string(buffer[:mLen]))
}
