package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() {
	data, err := os.ReadFile(dbfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file read:%v\n", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		readPiece := strings.Split(line, "-")
		id, _ := strconv.Atoi(readPiece[0])
		price, _ := strconv.Atoi(readPiece[3])
		stock, _ := strconv.Atoi(readPiece[4])
		pieces = append(pieces, Piece{
			Id:          id,
			Name:        readPiece[1],
			Category:    readPiece[2],
			Price:       price,
			Stock:       stock,
			Description: readPiece[5],
		})
	}
}
