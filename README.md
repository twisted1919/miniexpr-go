# MiniExpr - A Tiny Go Math Expression Parser

**MiniExpr** is a small, self-contained math expression scanner, parser and interpreter written in Go.  
It was built for **learning purposes** and for safely evaluating simple arithmetic expressions.

### Why?
I always wanted to write a scanner, parser and interpreter and since writing a programming language is a huge task, 
writing this library was the next best thing.  

## Features
- Fully deterministic **recursive-descent parser**    
- Supports:
    - Basic arithmetic: `+`, `-`, `*`, `/`
    - Exponentiation: `**` and `^` (right-associative)
    - Bit shifts: `<<`, `>>`
    - Unary negation: `-x`
    - Grouping with parentheses: `( ... )`
    - Numeric literals with underscores: `1_000_000`
    - Decimal numbers: `3.1415`
- Strict error detection for:
    - Stray or mismatched parentheses
    - Invalid underscore placement (`1__0`, `5_`)
    - Malformed numbers (`5.`, `.5`, `5.2.`)
    - Trailing or missing operands (`2**`, `1+`)
- Works directly with both `[]byte` and `string` input without extra allocations.
- Produces clear error messages with byte position. 

## Grammar Overview
```bnf
<expression> ::= <term>
<term>       ::= <factor> (("+" | "-") <factor>)*
<factor>     ::= <unary> (("*" | "/" | "<<" | ">>") <unary>)*
<unary>      ::= ("-" | "+") <unary> | <pow>
<pow>        ::= <primary> ( ("**" | "^") <pow> )?
<primary>    ::= NUMBER | "(" <expression> ")"
```  

## Why no multibyte support?  
This parser focuses on plain ASCII math expressions.    
While using `rune` iteration or `bufio.Reader` would make Unicode support straightforward,  
working directly with raw bytes was more interesting and educational for this project.  

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/twisted1919/miniexpr-go"
)

func main() {
	expr := "2**3**2 + (1_000_000 / (2<<3)) - (10 - 2) * (3 + 4) / 2 + -(5<<2) + (16>>3) + -2**2 + (-2)**3"

	result, err := miniexpr.EvaluateString(expr)
	if err != nil {
		log.Fatalf("unable to evaluate expression: %v", err)
	}

	fmt.Printf("The result for '%s' is '%v'.\n", expr, result) // The result is: 62954
}

```

### Testing
```bash 
go test ./...
```