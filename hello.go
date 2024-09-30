package main

import "fmt"

func main() {
	bytesWritten, err := fmt.Println("Hello World!")

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Bytes written", bytesWritten)
	}
}
