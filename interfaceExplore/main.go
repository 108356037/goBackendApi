package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
type customError struct{}

func (customError) Error() string {
	return "This is from custom error"
}

// our custom writer implements the writer interface (a Write() function required)
func (logWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, new(customError)
	}
	fmt.Println(string(bs))
	fmt.Printf("bs len: %v\n", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://i.nccu.edu.tw/Login.aspx?ReturnUrl=%2f")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// defer resp.Body.Close()
	// fmt.Println(string(bs))
	a := new(logWriter)
	io.Copy(a, resp.Body)
}
