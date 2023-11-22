package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	wd, _ := os.Getwd()
	file, err := os.Open( wd + "/map/road.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		fmt.Printf("The most common work is: %q", mcw(buf.String()))
	}
}

func mcw(content string) string {
	wordCount := make(map[string]int)
	maxCount, mcWord := 0, ""
	for _, word := range strings.Fields(content) {
		w := strings.ToLower(word)
		wordCount[w] = wordCount[w] + 1
		if wordCount[w] > maxCount {
			maxCount, mcWord = wordCount[w], w
		}
	}
	return mcWord
}
