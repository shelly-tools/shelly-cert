package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const CHUNK_LEN = 1024

func deleteCA(host string, method string) error {
	url := fmt.Sprintf("http://%s/rpc/Shelly.%s", host, method)
	payload := strings.NewReader("{\"data\": null}")
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 2}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func putChunk(host string, method string, data []byte) error {
	url := fmt.Sprintf("http://%s/rpc/Shelly.%s", host, method)
	payload := strings.NewReader(fmt.Sprintf("{\"data\": \"%s\", \"append\": true}", data))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 2}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	hostPtr := flag.String("host", "", "the host to send the file to")
	caFilePtr := flag.String("file", "", "the path to the CA file")
	methodPtr := flag.String("type", "PutUserCA", "Method is PutUserCa, PutTLSClientCert, PutTLSClientKey")
	flag.Parse()

	caData, err := ioutil.ReadFile(*caFilePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	caDataStr := string(caData)

	fmt.Println("Deleting existing file")
	err = deleteCA(*hostPtr, *methodPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	pos := 0
	fmt.Printf("total %d bytes\n", len(caDataStr))
	for pos < len(caDataStr) {
		chunk := caDataStr[pos:min(pos+CHUNK_LEN, len(caDataStr))]
		err := putChunk(*hostPtr, *methodPtr, []byte(chunk))
		if err != nil {
			fmt.Println(err)
			return
		}
		pos += len(chunk)
	}
	fmt.Println("Done")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
