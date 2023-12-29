package main

import (
	"bufio"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

func wget(url, filename string) error {

	r, err := getResponse(url)
	if err != nil {
		fmt.Printf("error getting %s: %v\n", url, err)
		return err
	}
	if len(filename) == 0 {
		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
		filename = params["filename"]
		return err
	}
	write(filename, r)
	return nil
}

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

func write(filename string, r *http.Response) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("error opening file %s: %v\n", filename, err)
	}
	defer f.Close()

	bufWriter := bufio.NewWriterSize(f, 1024*4)

	_, err = io.Copy(bufWriter, r.Body)
	if err != nil {
		fmt.Printf("error copying %v\n", err)
	}
}

func main() {
	err := wget("https://pkg.go.dev/net/http", "")
	if err != nil {
		fmt.Printf("error wget %v\n", err)
	}
}
