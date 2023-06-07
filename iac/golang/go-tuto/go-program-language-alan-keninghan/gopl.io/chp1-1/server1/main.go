package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	server1()
	//echo1()
	//echo2()

}

func server1() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func echo1() {
	args := ""

	if len(os.Args) == 1 {
		fmt.Printf("No command line provided")
	} else {

		for i := 1; i < len(os.Args); i++ {
			args += os.Args[i]
			args += " "

		}

	}

	fmt.Printf("%s", args)

}

func echo2() {
	fmt.Printf("%s", strings.Join(os.Args[1:], " "))
}
*/
