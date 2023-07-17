package main

import (
	"fmt"
	"log"
	"os"
	"quick/src/ast"
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/parser"
)

func main() {
	stmts := []stmt.Stmt{}
	errs := []*error.Error{}

	if len(os.Args) < 2 {
		log.Fatalln("usage: quick main.q")
	}

	for i := 1; i < len(os.Args); i++ {
		_stmts, _errs := parser.New(os.Args[i]).Parse()
		stmts = append(stmts, _stmts...)
		errs = append(errs, _errs...)
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.String())
		}

		log.Fatalln("failed to compile")
	}

	value, err := ast.New().Interpret(stmts)

	if err != nil {
		log.Fatalln(err)
	}

	if value != nil {
		log.Println(value.String())
	}
}
