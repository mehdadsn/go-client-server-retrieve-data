package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	dbfile = "db.txt"
)

type Piece struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

// works as our database
var pieces []Piece

func main() {
	readData()
	fmt.Println("Starting Server...")

	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("error acepting:", err.Error())
		}
		processConnection(connection)

	}
}

func processConnection(connection net.Conn) {
	defer connection.Close()
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("error reading: ", err.Error())
	}
	message := string(buffer[:mLen])
	params := strings.Split(message, "\n")

	//Get Name By ID
	if strings.Contains(message, "getnamebyid") {
		id, _ := strconv.Atoi(params[1])
		connection.Write([]byte(pieces[id].Name))
	}

	//Get In Stock
	if strings.Contains(message, "getinstock") {
		var res string
		for _, v := range pieces {
			if v.Stock > 0 {
				res += "ID:" + strconv.Itoa(v.Id) + " Name:" + v.Name + " Category:" + v.Category + " Price:" + strconv.Itoa(v.Price) + " Stock:" + strconv.Itoa(v.Stock) + " Description:" + v.Description + "\n"
			}
		}
		connection.Write([]byte(res))
	}

	//Order
	if strings.Contains(message, "order") {
		id, _ := strconv.Atoi(params[1])
		var response string
		if pieces[id].Stock < 1 {
			response = "Product Out Of Stock!"
			//connection.Write([]byte("Product Out Of Stock!"))
		} else {
			pieces[id].Stock -= 1
			response = fmt.Sprintf("Order Completed of one %v", pieces[id].Name)
		}
		connection.Write([]byte(response))
	}

	//In Stock By category
	if strings.Contains(message, "getstockbycategory") {
		var res string
		for _, v := range pieces {
			if params[1] == v.Category && v.Stock > 0 {
				res += "ID:" + strconv.Itoa(v.Id) + " Name:" + v.Name + " Category:" + v.Category + " Price:" + strconv.Itoa(v.Price) + " Stock:" + strconv.Itoa(v.Stock) + " Description:" + v.Description + "\n"
			}
		}
		connection.Write([]byte(res))
	}
	//req, _ := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(buffer[:mLen])))

	// if strings.Contains(req.URL.Path, "getnamebyid") {
	// 	getNameById(req, connection)
	// }
	// if strings.Contains(req.URL.Path, "getinstock") {

	// }
	// if strings.Contains(req.URL.Path, "order") {

	// }
	// if strings.Contains(req.URL.Path, "getinstockbycategory") {

	// }
}

// func getNameById(r *http.Request, connection net.Conn) {
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
// 	piece := pieces[id]
// 	connection.Write([]byte(piece.Name))
// }
