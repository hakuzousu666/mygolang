package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	r, err := http.Get("https://v.hantools.top/bot/data/get/six")
	if err != nil {
		log.Fatalln("报错", err)
	}
	fmt.Println(r)
	time.Sleep(time.Second * 3)
	return
}
