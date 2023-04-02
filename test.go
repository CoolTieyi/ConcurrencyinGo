package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	fmt.Print("Enter your grade : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(reflect.TypeOf(input))
	// string->Float
	//inpu := strings.Trim(input, '\n')

	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(grade))
	fmt.Println(grade)
}
