package main

import (
	"log"
	"os"
	"quick/src/scanner"
	"quick/src/token"
)

func main() {
	files := map[string][]byte{}

	if len(os.Args) < 2 {
		log.Fatalln("usage: quick main.q")
	}

	for i := 1; i < len(os.Args); i++ {
		bytes, err := os.ReadFile(os.Args[i])

		if err != nil {
			log.Fatal(err)
		}

		files[os.Args[i]] = bytes
		scanner := scanner.New(os.Args[i], bytes)

		for {
			t, err := scanner.Next()

			if err != nil {
				log.Fatal(err)
			}

			if t.Kind == token.EOF {
				break
			}

			log.Printf("[%s] -> %s\n", t.Path, t.String())
		}
	}
}
