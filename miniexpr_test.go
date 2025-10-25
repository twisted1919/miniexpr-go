package miniexpr_test

import (
	"fmt"
	"testing"

	// nolint: nolintlint, depguard
	"github.com/twisted1919/miniexpr-go"
)

// nolint: nolintlint, funlen
func TestEvaluate_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expr string
		want float64
	}{
		// literals and whitespace
		{"0", 0},
		{"1", 1},
		{"42", 42},
		{"  7  ", 7},
		{"1234567890", 1234567890},

		// underscore grouping (between digits only)
		{"1_000_000", 1000000},
		{"12_345 + 6_789", 19134},
		{"(1_2_3) + (4_5_6)", 579},
		{"9_9_9_9 - 8_8_8_8", 1111},
		{"1_000*2 + 3_000/6", 2000 + 500},
		{"(1_000_000) / (1_0_0)", 1000000.0 / 100.0},

		// +, -, *, / precedence and associativity
		{"1+2", 3},
		{"1+2*3", 7},
		{"(1+2)*3", 9},
		{"2*3+4*5-6/3", 24},
		{"8/4/2", 1}, // left-assoc division
		{"7-4-1", 2}, // left-assoc subtraction
		{"4/2*3", 6}, // left-to-right at same precedence
		{"4*(2/3)", 4.0 * (2.0 / 3.0)},

		// unary +/- and binding vs power
		{"-5+2", -3},
		{"+5", 5},
		{"-(1+2)", -3},
		{"-(1)", -1},
		{"-(-(-2))", -2},

		// power (both ** and ^), right-assoc, tighter than unary
		{"2**3", 8},
		{"2^3", 8},
		{"2**3**2", 512}, // 2 ** (3 ** 2)
		{"3^2^3", 6561},  // 3 ^ (2 ^ 3)
		{"2**(3**2)", 512},
		{"(2**3)**2", 64},
		{"-2**2", -4}, // -(2**2)
		{"(-2)**2", 4},
		{"-(2**3)", -8},
		{"-2**3", -8},
		{"0**0", 1},
		{"2**0", 1},
		{"2**1", 2},

		// mixed power with parentheses
		{"2**(3+2)", 32},
		{"2**3+2", 10},
		{"2**(3+2*2)", 128},
		{"((2))**((3))", 8},
		{"(((((1+2)))))", 3},

		// shifts (same precedence as * and / in our grammar)
		{"2<<3", 16},
		{"16>>3", 2},
		{"1+2<<3", 17}, // 1 + (2<<3)
		{"(1+2)<<3", 24},
		{"8>>2+1", 3}, // (8>>2) + 1
		{"8>>(2+1)", 1},
		{"2<<1<<2", 16}, // left-assoc
		{"32>>1>>3", 2}, // left-assoc
		{"2*3<<2", 24},  // (2*3)<<2
		{"2<<3*2", 32},  // (2<<3)*2
		{"-(2<<3)", -16},
		{"((2<<2)) + ((8>>1))", 12},
		{"2<<3/2", 8}, // (2<<3)/2

		// whitespace variants
		{"\t2\t+\n3", 5},
	}

	// nolint: nolintlint, varnamelen
	for i, tc := range tests {
		t.Run(fmt.Sprintf("%02d_%s", i+1, tc.expr), func(t *testing.T) {
			t.Parallel()

			got, err := miniexpr.EvaluateBytes([]byte(tc.expr))
			if err != nil {
				t.Fatalf("Evaluate(%q) returned error: %v", tc.expr, err)
			}

			if got != tc.want {
				t.Fatalf("Evaluate(%q) = %v, want %v", tc.expr, got, tc.want)
			}
		})
	}
}

func TestEvaluate_SyntaxErrors(t *testing.T) {
	t.Parallel()

	errCases := []string{
		"1__0",
		"_1",
		"1_",
		"1_000_",
		"5_2__1",
		"__",
		"1___2",
		"1_ 2", // underscore not directly between digits

		// decimals not accepted
		"5.2.",
		".5",
		"5.",
		"1_.0",
		"1._0",

		// power token issues
		"2**",  // missing rhs
		"**2",  // missing lhs
		"2^^3", // not a valid power token sequence
		"2**^3",
		"^",
		"**",

		// binary operator missing operand
		"2+",
		"2*/3",
		"2//3",
		"*2",
		"/2",

		// malformed shifts
		"2<<",
		">>2",
		"2<<>>1",
		"2<<<3",
		"3>>>>1",

		// paren mismatches
		"(1+2",
		"1+2)",
		")1+2(",
		"(",

		// unknown tokens / identifiers
		"X",
		"--X",
	}

	for i, expr := range errCases {
		t.Run(fmt.Sprintf("err_%02d_%s", i+1, expr), func(t *testing.T) {
			t.Parallel()

			_, err := miniexpr.EvaluateBytes([]byte(expr))
			if err == nil {
				t.Fatalf("Evaluate(%q) expected error, got nil", expr)
			}
		})
	}
}

func TestEvaluateString(t *testing.T) {
	t.Parallel()

	expr := `2**3**2 + (1_000_000 / (2<<3)) - (10 - 2) * (3 + 4) / 2 + -(5<<2) + (16>>3) + -2**2 + (-2)**3`

	res, err := miniexpr.EvaluateString(expr)
	if err != nil {
		t.Fatalf("Evaluate(%q) returned error: %v", expr, err)
	}

	if res != 62954 {
		t.Fatalf("Evaluate(%q) = %v, want %v", expr, res, 1)
	}
}
