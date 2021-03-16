package tool

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Read() {
	fmt.Println("Hi")
	readText, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %s", err)
	}
	fmt.Printf("\nLength: %d", len(readText))
	fmt.Printf("\nData Read: \n%s", readText)
}
