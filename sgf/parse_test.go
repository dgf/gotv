// test specification examples, see:
// http://www.red-bean.com/sgf/sgf4.html
// http://www.red-bean.com/sgf/var.htm
package sgf_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/sgf"
)

// SGF basics and property value clean up
func TestParse_cleanup(t *testing.T) {
	for _, check := range []struct{ name, sgf, exp string }{
		{"2.1. EBNF spaces",
			" \r\n \r \n \n\r \t ( \n\r\t\r\n\n ; \r\n\f\n\n\r AN \r\n\t\n [ebnf] \n\r\f\n ) \n\n\r\n ",
			"(;AN[ebnf])"},
		{"3. UcLetter, Digit, None",
			"(;U[AZ];D[9876543210];N[])",
			"(;U[AZ];D[9876543210];N[])"},
		{"3. Number +/-",
			"(;N[1];N[+1];N[-1])",
			"(;N[1];N[+1];N[-1])"},
		{"3. Real +/-",
			"(;R[0.1];R[+0.2];R[-0.3])",
			"(;R[0.1];R[+0.2];R[-0.3])"},
		{"3.1. Double",
			"(;CB[1];CB[2])",
			"(;CB[1];CB[2])"},
		{"3.2. Text escape",
			`(;C["t \\ e \: ' x \] t])`,
			`(;C["t \\ e \: ' x \] t])`},
		{"3.2. Text unicode",
			"(;C[\xE2\x98\xA0] ",
			"(;C[☠])"},
		{"3.2. Text chars",
			"(;C[?!§$%&/ ()=?`*'_:;><´+ #-.,¹²³¼½¬ {[}\\¸~ µ@])",
			"(;C[?!§$%&/ ()=?`*'_:;><´+ #-.,¹²³¼½¬ {[}\\¸~ µ@])"},
		{"3.2. Text feeds",
			"(;C[ \f \n\n \n\r \n \r \r\n \t t \f\t\r e \n\f\n\t\n x \t\f\n t \f \n\n \n\r \n \r \r\n \t ])",
			"(;C[t\ne\nx\nt])"},
		{"3.2. Text comments",
			"(;C[ /* \f\n\r\t c\no\rm\tm\ne\fn\nt \f\n\r\t*/ ])",
			"(;C[/*\nc\no\nm m\ne n\nt\n*/])"},
		{"3.3. SimpleText spaces",
			"(;AN[ \f\n\t\r s \f \t \n \t\f p \n\r a\rc \f\r\t e \n s])",
			"(;AN[s p ac e s])"},
	} {
		act := sgf.Parse(check.sgf).String()
		if check.exp != act {
			t.Errorf("%s\nEXP: %s\nACT: %s\n", check.name, check.exp, act)
		}
	}
}

// No Variation
func ExampleParse_noVariation() {
	s := "(;GM[1]SZ[19]FF[4];B[aa];W[bb];B[cc])"
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb];B[cc])
}

// One variation at move 3
func ExampleParse_oneVariationAtMove3() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc];W[dd];B[ad];W[bd])
	        (;B[hh];W[hg]))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc];W[dd];B[ad];W[bd])(;B[hh];W[hg]))
}

// Two variations at move 3
func ExampleParse_twoVariationAtMove3() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc]N[A];W[dd];B[ad];W[bd])
	        (;B[hh]N[B];W[hg])
	        (;B[gg]N[C];W[gh];B[hh];W[hg];B[kk]))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc]N[A];W[dd];B[ad];W[bd])(;B[hh]N[B];W[hg])(;B[gg]N[C];W[gh];B[hh];W[hg];B[kk]))
}

// Two variations at different moves
func ExampleParse_twoVariationAtDifferentMoves() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc];W[dd]
	          (;B[ad];W[bd])
	          (;B[ee];W[ff])
	        )
	        (;B[hh];W[hg]))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc];W[dd](;B[ad];W[bd])(;B[ee];W[ff]))(;B[hh];W[hg]))
}

// Variation of a variation
func ExampleParse_variationOfVariation() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc]N[A];W[dd];B[ad];W[bd])
	          (;B[hh]N[B];W[hg])
	          (;B[gg]N[C];W[gh];B[hh]
	            (;W[hg]N[A];B[kk])
	            (;W[kl]N[B])))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc]N[A];W[dd];B[ad];W[bd])(;B[hh]N[B];W[hg])(;B[gg]N[C];W[gh];B[hh](;N[A]W[hg];B[kk])(;N[B]W[kl])))
}

var benchSGF = "(;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc];W[dd];B[ad];W[bd])(;B[hh];W[hg])(;B[gg];W[gh];B[hh](;W[hg];B[kk])(;W[kl])))"

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = sgf.Parse(benchSGF)
	}
}
