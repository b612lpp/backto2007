package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ydResp struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

func main() {
	var a ydResp
	publicURL := getURL()
	longURL := generateURL(publicURL)
	rawBody := getInfo(longURL)
	downloadLync := extractD(rawBody, a)
	downloadFile(downloadLync)

}

//get public link from console
func getURL() (h string) {
	fmt.Println("Inser public url")
	fmt.Scan(&h)
	return
}

//generate link to get meta information
func generateURL(h string) (g string) {

	g = "https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=" + h
	fmt.Println(g)
	return
}

//get meta information
func getInfo(u string) (body []byte) {
	req, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
	}
	body, _ = ioutil.ReadAll(req.Body)
	//fmt.Println(body)
	return
}

//extract download link from meta(json)
func extractD(b []byte, d ydResp) (u string) {
	jsonErr := json.Unmarshal(b, &d)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	u = d.Href

	return
}

//download file(function from internet)
func downloadFile(url string) (err error) {
	var filepath string
	fmt.Println("input full path for downloaded file with name")
	fmt.Scan(&filepath)
	fmt.Println(filepath)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
