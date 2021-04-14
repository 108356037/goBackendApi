package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func testWrite(filename string, content string) error {
	if err := ioutil.WriteFile(filename, []byte(content), 0755); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := testWrite("test.py", "#!/usr/bin/env python\nprint(\"Hello World from Python\")"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("success write\n")

	b, err := ioutil.ReadFile("./test.py")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	splitStrArr := strings.Split(string(b), "/")
	fmt.Printf("%v\n", string(b))
	for i, val := range splitStrArr {
		fmt.Printf("the item at index %d: %v\n", i, val)
	}
}
