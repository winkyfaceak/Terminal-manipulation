package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 0 {
		log.Fatal("Please provide a file or directory path to delete")
	}

	for _, path := range os.Args[1:] {
		info, err := os.Stat(path)

		if os.IsNotExist(err) {
			log.Fatalf("The path '%s' does not exist.", path)
		} else if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatalf("Failed to delete directory '%s': %v", path, err)
			}
			fmt.Printf("Successfully deleted directory: %s\n", path)
		} else {
			err = os.Remove(path)
			if err != nil {
				log.Fatalf("Failed to delete file '%s': %v", path, err)
			}
			fmt.Printf("Successfully deleted file: %s\n", path)
		}
	}
}
