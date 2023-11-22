package main

import (
	"fmt"
	"os"
)


func main() {
	if err := removeFile("pidFile.pid"); err != nil {
		fmt.Println(err)
	}
}

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("can't remove file: %s.  %w", fileName, err)
	}
	return nil
}
