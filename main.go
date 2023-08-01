package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Data []struct {
		Da       string `json:"DA"`
		Pa       string `json:"PA"`
		Moz      string `json:"Moz_Rank"`
		BackLink string `json:"Back_Links"`
	} `json:"data"`
}

func main() {
	var list string
	fmt.Println("Credit @zeruproject")
	fmt.Print("Your list: ")
	fmt.Scanln(&list)
	CheckDapa(list)
}

func Api(url string) {
	var data Response
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   120 * time.Second,
	}
	request, err := http.NewRequest("GET", "https://tools.helixs.id/API/dapa?url="+url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	body, _ := io.ReadAll(response.Body)

	defer response.Body.Close()
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Print(err.Error())
	}
	result := fmt.Sprintf("Url : %s [DA : %v] [PA : %v] [Moz : %v] [Backlink : %v]", url, data.Data[0].Da, data.Data[0].Pa, data.Data[0].Moz, data.Data[0].BackLink)
	fmt.Println(result)
	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = fmt.Fprintf(file, result+"\n")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func CheckDapa(list string) {
	file, err := os.Open(list)
	if err != nil {
		fmt.Print(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weblist := scanner.Text()
		Api(weblist)
	}
}
