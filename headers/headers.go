package headers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	Params1 string
	Params2 string
	Url     string
	Headers []string

	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyz"
	lengthParm = 10
)

func report(data string) {
	currentTime := time.Now()
	date := currentTime.Format("01-02-2006 15:04:05")
	f, err := os.OpenFile("report"+date+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data + "\n"); err != nil {
		log.Println(err)
	}
}

/*
Genere una stringa di 10 caratteri
*/
func GeneretorRandomString() string {
	b := make([]byte, lengthParm)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/*
Legge il file headers.dat
*/
func Readfile() []string {
	dat, err := os.ReadFile("./headers/headers.dat")
	if err != nil {
		fmt.Println(err)
	}
	f := string(dat)
	return strings.Split(f, "\n")
}

/*
Richiesta con headers
*/
func MakeRequestWithHeaders(headers string) *http.Response {
	req, err := http.NewRequest("GET", Url+"?"+Params2, nil)
	fmt.Printf("[*]Try %s?%s with headers %s\n", Url, Params2, headers)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Set("Connection", "keep-alive, "+headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err)
		return nil
	}
	//defer resp.Body.Close()

	return resp
}

/*
Richiesta senza headers
*/
func MakeRequestNoHeaders() *http.Response {
	req, err := http.NewRequest("GET", Url+"?"+Params1, nil)
	fmt.Printf("[*]Try %s?%s no headers\n", Url, Params1)
	if err != nil {
		log.Println(err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err)
		return nil
	}
	//defer resp.Body.Close()
	return resp
}

func Check(resp1 *http.Response, resp2 *http.Response, headers string) {

	text, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Lunghezza byte response uno
	lengResp1 := len(string(text))

	text, err = ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Lunghezza byte response due
	lengResp2 := len(string(text))

	if resp1.StatusCode != resp2.StatusCode || lengResp1 != lengResp2 {
		if resp1.StatusCode != resp2.StatusCode {
			fmt.Printf("[+]URL: %s\nRequest one - Status Code: %d\nRequest two - Status Code: %d\n\nHeaders used: %s\n", Url, resp1.StatusCode, resp2.StatusCode, headers)
			report("URL: " + Url + "\nRequest one - Status Code: " + string(resp1.StatusCode) + "\nRequest two - Status Code: " + string(resp2.StatusCode) + "\n\nHeaders used: " + headers + "\n")
		}
		if lengResp1 != lengResp2 {
			fmt.Printf("[+]URL: %s\nRequest one - Size: %d Bytes\nRequest two - Size: %dBytes\n\nHeaders used: %s\n", Url, lengResp1, lengResp2, headers)
			report("URL: " + Url + "\nRequest one - Size: " + string(lengResp1) + " Bytes\nRequest two - Size: " + string(lengResp2) + " Bytes\n\nHeaders used: " + headers + "\n")
		}
	}
	resp3 := MakeRequestNoHeaders()
	if resp3 != nil {
		if resp3.StatusCode == resp2.StatusCode {
			fmt.Printf("[+]%s?cb=%s poisoned?\n", Url, Params2)
			report(Url + "?cb=" + Params2 + " poisoned?\n")
		} else {
			fmt.Printf("[-]No poisoning detected\n")
		}
	}
}
