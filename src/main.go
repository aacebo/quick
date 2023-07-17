package main

import (
	"fmt"
	"log"
	"os"
	"quick/src/ast"
	"quick/src/ast/stmt"
	"quick/src/parser"
)

func main() {
	stmts := [][]stmt.Stmt{}

	if len(os.Args) < 2 {
		log.Fatalln("usage: quick main.q")
	}

	for i := 1; i < len(os.Args); i++ {
		_stmts, errs := parser.New(os.Args[i]).Parse()

		if errs != nil && len(errs) > 0 {
			for _, err := range errs {
				fmt.Println(err.String())
			}

			log.Fatalln("failed to compile")
		}

		stmts = append(stmts, _stmts)
	}

	for _, _stmts := range stmts {
		value, err := ast.New().Interpret(_stmts)

		if err != nil {
			log.Fatalln(err)
		}

		if value != nil {
			log.Println(value.String())
		}
	}
}
