package main

import (
	"fmt"
	"time"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "Hello World")
	// })

	// http.ListenAndServe(":3000", nil)

	for {
		fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		time.Sleep(1 * time.Second)
	}

}
