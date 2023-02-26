// @Title
// @Description
// @Author
// @Update
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var counter = 0

func Greet(writer io.Writer, name string) {
	counter++
	fmt.Printf("called %d times \n", counter)
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Jack Boss")
}

func main() {
	// Greet(os.Stdout, "Jack")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
