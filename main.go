package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求参数", r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)
}

func main() {
	fmt.Println("run web start...")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
