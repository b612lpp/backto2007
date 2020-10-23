package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	listPages := getPagesToSearch()
	searchWord := getStringToSearch()
	for _, v := range listPages {
		wg.Add(1)
		go searchString(v, searchWord, &wg)

	}
	wg.Wait()
}

func searchString(v, sw string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(v)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	a := string(body)
	//fmt.Println(a)
	defer resp.Body.Close()
	if strings.Contains(a, sw) {
		fmt.Println(v)
	} else {
		fmt.Println("false", v)
	}

}
func getStringToSearch() (t string) {

	fmt.Println("Let us know what should be found")
	fmt.Scanln(&t)
	return t
}
func getPagesToSearch() (p []string) {
	var str string
	fmt.Println("Please, specify pages should be checked. BE NOTIFIED: pages must be separated by comma like https://ya.ru,https://google.com")
	fmt.Scan(&str)
	p = strings.Split(str, ",")
	return p
}
