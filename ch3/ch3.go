package ch3

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Demo35() {
	resp, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
}
