package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"strconv"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r)

	w.Write([]byte("Hello, world."))
}

type ListItem struct {
	Name string
	Age  int64
}

var data = []ListItem{
	{"dubby.cn", 2},
	{"道玄", 1},
	{"杨正", 24},
}

func list(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("json marshal fail: ", err)
	}
	w.Write(jsonData)
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	ageStr := r.Form.Get("age")

	if name != "" && ageStr != "" {
		age, err := strconv.ParseInt(ageStr, 0, 64)
		if err != nil {
			log.Fatal("age format error: ", ageStr)
		}
		item := ListItem{Name: name, Age: age}
		data = append(data, item)
	}

	list(w, r)
}

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/list", list)
	http.HandleFunc("/add", add)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen fail: ", err)
	}
}
