package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("go> ")
	for scanner.Scan() {
		fmt.Println(strconv.Quote(scanner.Text()))
		fmt.Print("go> ")
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
