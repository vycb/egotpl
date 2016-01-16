package main

import (
	"time"
	"fmt"
	//"bytes"
	"net/http"
	."github.com/vycb/egotpl/tpl"
)


func main() {

	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
	//var buf bytes.Buffer
	//fmt.Println(&buf)
}

func root(w http.ResponseWriter, r *http.Request) {

	page := &Page{
		Title: "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
		Year: time.Now().Year(),
	}
	if err := Home(w, page); err != nil {
		fmt.Println(err)
	}

}

