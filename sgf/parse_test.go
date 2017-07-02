package sgf_test

// test specification examples, see http://www.red-bean.com/sgf/var.htm

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/sgf"
)

// text escaping
func ExampleParse_Comment() {
	s := "(;GM[1]SZ[19]FF[4];B[aa]\f\n\r\tC\f\n\r\t [\n \fc\t[o \f\t  m\\]\r\n\tm\\\\e\\:n\n\r (FF[1\\])t\r])"
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa]C[c [o m\]
	//  m\\e\:n
	//  (FF[1\])t
	// ])
}

// No Variation, with unsorted properties and nearly all possible space and line feed combinations
func ExampleParse_NoVariation() {
	s := "\n(\r;\tGM [\n1\r]\nSZ\r[\t19\n]\rFF\n[\r4\t]\n;\rB\t[\naa\r]\t;\nW\r[\tbb\n]\r;\tB [ cc ]\n\r\t\r\n)"
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb];B[cc])
}

// One variation at move 3
func ExampleParse_OneVariationAtMove3() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc];W[dd];B[ad];W[bd])
	        (;B[hh];W[hg]))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc];W[dd];B[ad];W[bd])(;B[hh];W[hg]))
}

// Two variations at move 3
func ExampleParse_TwoVariationAtMove3() {
	s := `(;FF[4]GM[1]SZ[19];B[aa];W[bb]
	        (;B[cc]N[A];W[dd];B[ad];W[bd])
	        (;B[hh]N[B];W[hg])
	        (;B[gg]N[C];W[gh];B[hh];W[hg];B[kk]))`
	fmt.Println(sgf.Parse(s))
	// Output:
	// (;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc]N[A];W[dd];B[ad];W[bd])(;B[hh]N[B];W[hg])(;B[gg]N[C];W[gh];B[hh];W[hg];B[kk]))
}

// Two variations at different moves
func ExampleParse_TwoVariationAtDifferentMoves() {
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
func ExampleParse_VariationOfVariation() {
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

func BenchmarkParse(b *testing.B) {
	s := "(;FF[4]GM[1]SZ[19];B[aa];W[bb](;B[cc];W[dd];B[ad];W[bd])(;B[hh];W[hg])(;B[gg];W[gh];B[hh](;W[hg];B[kk])(;W[kl])))"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sgf.Parse(s)
	}
}
