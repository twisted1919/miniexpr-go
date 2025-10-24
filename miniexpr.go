// Package miniexpr evaluates expressions to their result.
package miniexpr

func evaluate(src DataSource) (float64, error) {
	scanner := NewScanner(src)

	tokens, err := scanner.Scan()
	if err != nil {
		return 0, err
	}

	parser := NewParser(tokens)

	expr, err := parser.Parse()
	if err != nil {
		return 0, err
	}

	interpreter := NewInterpreter()

	res, err := interpreter.Interpret(expr)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// EvaluateBytes will scan, parse and interpret given byte slice source.
func EvaluateBytes(src []byte) (float64, error) {
	return evaluate(ByteSource(src))
}

// EvaluateString will scan, parse and interpret given string source.
func EvaluateString(src string) (float64, error) {
	return evaluate(StringSource(src))
}
