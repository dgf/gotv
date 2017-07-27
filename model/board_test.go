package model_test

import (
	"fmt"
	"testing"

	"github.com/dgf/gotv/model"
)

func TestBoardValidation(t *testing.T) {
	b := model.NewBoard(model.X13)

	_, err := b.PlaceSGF(B, "zz")
	if err == nil {
		t.Error("unknown SGF placement expected")
	}

	_, err = b.PlaceIGS(B, "I7")
	if err == nil {
		t.Error("unknown IGS placement expected")
	}

	_, err = b.PlaceIGS(B, "Q17")
	if err == nil {
		t.Error("out of range expected")
	}

	_, _ = b.Place(B, P(7, 4))
	_, err = b.Place(W, P(7, 4))
	if err == nil {
		t.Error("occupied before expected")
	}
}

func ExampleBoard_Groups() {
	b := model.NewBoard(model.Size(4))
	for _, a := range []struct {
		c model.Color
		p model.Point
		r model.Points
	}{
		{c: B, p: P(1, 1), r: model.Points{}},
		{c: W, p: P(2, 1), r: model.Points{}},
		{c: B, p: P(2, 2), r: model.Points{}},
		{c: W, p: P(1, 2), r: model.Points{P(1, 1)}},
		{c: B, p: P(1, 3), r: model.Points{}},
		{c: W, p: P(1, 1), r: model.Points{}},
		{c: B, p: P(3, 1), r: model.Points{P(1, 1), P(1, 2), P(2, 1)}},
		{c: W, p: P(1, 1), r: model.Points{}},
		{c: B, p: P(1, 2), r: model.Points{}},
	} {
		fmt.Println(a.c, a.p, a.r)
		r, err := b.Place(a.c, a.p)
		if err != nil {
			panic(err)
		}
		d := model.DiffPoints(r, a.r)
		if len(d) != 0 {
			panic(fmt.Sprintf("captures diff %v", d))
		}
		fmt.Println(b.Groups)
	}
	fmt.Println(b)

	// Output:
	// B {1 1} []
	//  1 B m[{1 1}] l[{1 2} {2 1}]
	// W {2 1} []
	//  1 B m[{1 1}] l[{1 2}]
	//  2 W m[{2 1}] l[{2 2} {3 1}]
	// B {2 2} []
	//  1 B m[{1 1}] l[{1 2}]
	//  2 W m[{2 1}] l[{3 1}]
	//  3 B m[{2 2}] l[{1 2} {2 3} {3 2}]
	// W {1 2} [{1 1}]
	//  2 W m[{2 1}] l[{1 1} {3 1}]
	//  3 B m[{2 2}] l[{2 3} {3 2}]
	//  4 W m[{1 2}] l[{1 1} {1 3}]
	// B {1 3} []
	//  2 W m[{2 1}] l[{1 1} {3 1}]
	//  3 B m[{2 2}] l[{2 3} {3 2}]
	//  4 W m[{1 2}] l[{1 1}]
	//  5 B m[{1 3}] l[{1 4} {2 3}]
	// W {1 1} []
	//  3 B m[{2 2}] l[{2 3} {3 2}]
	//  5 B m[{1 3}] l[{1 4} {2 3}]
	//  6 W m[{1 1} {1 2} {2 1}] l[{3 1}]
	// B {3 1} [{1 1} {1 2} {2 1}]
	//  3 B m[{2 2}] l[{1 2} {2 1} {2 3} {3 2}]
	//  5 B m[{1 3}] l[{1 2} {1 4} {2 3}]
	//  7 B m[{3 1}] l[{2 1} {3 2} {4 1}]
	// W {1 1} []
	//  3 B m[{2 2}] l[{1 2} {2 1} {2 3} {3 2}]
	//  5 B m[{1 3}] l[{1 2} {1 4} {2 3}]
	//  7 B m[{3 1}] l[{2 1} {3 2} {4 1}]
	//  8 W m[{1 1}] l[{1 2} {2 1}]
	// B {1 2} []
	//  7 B m[{3 1}] l[{2 1} {3 2} {4 1}]
	//  8 W m[{1 1}] l[{2 1}]
	//  9 B m[{1 2} {1 3} {2 2}] l[{1 4} {2 1} {2 3} {3 2}]
	// +abcd+
	// a◯ ● |
	// b●●  |
	// c●   |
	// d    |
	// {1 2} w1 b3
}

func ExampleBoard_KO() {
	b := model.NewBoard(model.Size(4))

	_, _ = b.Place(B, P(1, 1))
	_, _ = b.Place(W, P(2, 1))
	_, _ = b.Place(B, P(2, 2))
	fmt.Println(b)

	r, _ := b.Place(W, P(1, 2))
	fmt.Println(b)
	d := model.DiffPoints(r, model.Points{P(1, 1)})
	if len(d) != 0 {
		panic(fmt.Sprintf("captures diff %v", d))
	}

	_, _ = b.Place(B, P(1, 3))
	_, _ = b.Place(W, P(2, 3))
	r, _ = b.Place(B, P(1, 1))
	fmt.Println(b)
	d = model.DiffPoints(r, model.Points{P(1, 2)})
	if len(d) != 0 {
		panic(fmt.Sprintf("captures diff %v", d))
	}

	if _, err := b.Place(W, P(1, 2)); err == nil {
		panic("KO placement")
	}

	_, _ = b.Place(W, P(3, 1))
	_, _ = b.Place(B, P(3, 2))
	r, _ = b.Place(W, P(1, 2))
	fmt.Println(b)
	d = model.DiffPoints(r, model.Points{P(1, 1)})
	if len(d) != 0 {
		panic(fmt.Sprintf("captures diff %v", d))
	}

	if _, err := b.Place(B, P(1, 1)); err == nil {
		panic("KO placement")
	}

	_, _ = b.Place(B, P(1, 4))
	_, _ = b.Place(W, P(4, 2))
	r, _ = b.Place(B, P(1, 1))
	fmt.Println(b)
	d = model.DiffPoints(r, model.Points{P(1, 2)})
	if len(d) != 0 {
		panic(fmt.Sprintf("captures diff %v", d))
	}

	if _, err := b.Place(W, P(1, 2)); err == nil {
		panic("KO placement")
	}

	_, _ = b.Place(W, P(4, 1))
	_, _ = b.Place(B, P(3, 3))
	_, _ = b.Place(W, P(1, 2))

	if _, err := b.Place(B, P(1, 1)); err == nil {
		panic("KO placement")
	}

	r, _ = b.Place(B, P(2, 4))
	d = model.DiffPoints(r, model.Points{P(2, 3)})
	if len(d) != 0 {
		panic(fmt.Sprintf("captures diff %v", d))
	}
	_, _ = b.Place(W, P(1, 1))

	fmt.Println(b)
	fmt.Println(b.Groups)
	// Output:
	// +abcd+
	// a●◯  |
	// b ●  |
	// c    |
	// d    |
	// {2 2} w0 b0
	// +abcd+
	// a ◯  |
	// b◯●  |
	// c    |
	// d    |
	// {1 2} w1 b0
	// +abcd+
	// a●◯  |
	// b+●  |
	// c●◯  |
	// d    |
	// {1 1} w1 b1
	// +abcd+
	// a+◯◯ |
	// b◯●● |
	// c●◯  |
	// d    |
	// {1 2} w2 b1
	// +abcd+
	// a●◯◯ |
	// b+●●◯|
	// c●◯  |
	// d●   |
	// {1 1} w2 b2
	// +abcd+
	// a◯◯◯◯|
	// b◯●●◯|
	// c● ● |
	// d●●  |
	// {1 1} w3 b3
	// 15 B m[{2 2} {3 2} {3 3}] l[{2 3} {3 4} {4 3}]
	// 17 B m[{1 3} {1 4} {2 4}] l[{2 3} {3 4}]
	// 18 W m[{1 1} {1 2} {2 1} {3 1} {4 1} {4 2}] l[{4 3}]
}

func ExampleBoard_suicide() {
	b := model.NewBoard(model.Size(4))

	// setup
	for _, a := range []struct {
		c model.Color
		p model.Point
	}{
		{c: B, p: P(1, 2)},
		{c: B, p: P(1, 4)},
		{c: B, p: P(2, 1)},
		{c: B, p: P(2, 2)},
		{c: B, p: P(2, 3)},
		{c: B, p: P(2, 4)},
		{c: W, p: P(3, 1)},
		{c: W, p: P(3, 2)},
		{c: W, p: P(3, 3)},
		{c: W, p: P(3, 4)},
		{c: W, p: P(4, 1)},
		{c: W, p: P(4, 3)},
	} {
		_, _ = b.Place(a.c, a.p)
	}

	// suicides
	for _, a := range []struct {
		c model.Color
		p model.Point
	}{
		{c: B, p: P(4, 2)},
		{c: B, p: P(4, 4)},
		{c: W, p: P(1, 1)},
		{c: W, p: P(1, 3)},
	} {
		if r, err := b.Place(a.c, a.p); err == nil || len(r) != 0 {
			fmt.Println(fmt.Sprintf("suicide %s %v expected", a.c, a.p))
		}
	}

	fmt.Println(b)
	// Output:
	//+abcd+
	//a ●◯◯|
	//b●●◯ |
	//c ●◯◯|
	//d●●◯ |
	//{4 3} w0 b0
}

func ExampleBoard_PlaceSGF() {
	b := model.NewBoard(model.X9)
	for _, a := range []struct {
		c model.Color
		i string
	}{
		{c: B, i: "gd"},
		{c: W, i: "df"},
		{c: B, i: "gg"},
		{c: W, i: "ec"},
		{c: B, i: "cc"},
		{c: W, i: "jj"},
		{c: B, i: "aj"},
		{c: W, i: "aa"},
		{c: B, i: "ja"},
	} {
		_, err := b.PlaceSGF(a.c, a.i)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(b)
	_, err := b.PlaceSGF(W, "ja")
	fmt.Println(err)
	// Output:
	// +abcdefghj+
	// a◯       ●|
	// b         |
	// c  ● ◯    |
	// d      ●  |
	// e         |
	// f   ◯     |
	// g      ●  |
	// h         |
	// j●       ◯|
	// {9 1} w0 b0
	// {9 1} is occupied
}

func TestCaptureSGF(t *testing.T) {
	b := model.NewBoard(model.Size(3))
	_, _ = b.PlaceSGF(B, "cc")
	_, _ = b.PlaceSGF(B, "bc")
	_, _ = b.PlaceSGF(W, "ac")
	_, _ = b.PlaceSGF(W, "bb")
	r, _ := b.PlaceSGF(W, "cb")
	if len(r) != 2 || r[0] != "bc" || r[1] != "cc" {
		t.Error("capture expected", r)
	}
}

func ExampleBoard_PlaceIGS() {
	b := model.NewBoard(model.X13)
	for _, a := range []struct {
		c model.Color
		i string
	}{
		{c: B, i: "L10"},
		{c: W, i: "D4"},
		{c: B, i: "K3"},
		{c: W, i: "D10"},
		{c: B, i: "J11"},
		{c: W, i: "N13"},
		{c: B, i: "A13"},
		{c: W, i: "A1"},
		{c: B, i: "N1"},
	} {
		_, err := b.PlaceIGS(a.c, a.i)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(b)
	_, err := b.PlaceIGS(W, "N1")
	fmt.Println(err)
	// Output:
	// +abcdefghjklmn+
	// a●           ◯|
	// b             |
	// c        ●    |
	// d   ◯      ●  |
	// e             |
	// f             |
	// g             |
	// h             |
	// j             |
	// k   ◯         |
	// l         ●   |
	// m             |
	// n◯           ●|
	// {13 13} w0 b0
	// {13 13} is occupied
}

func TestCaptureIGS(t *testing.T) {
	b := model.NewBoard(model.Size(3))
	_, _ = b.PlaceIGS(B, "C1")
	_, _ = b.PlaceIGS(B, "B1")
	_, _ = b.PlaceIGS(W, "A1")
	_, _ = b.PlaceIGS(W, "B2")
	r, _ := b.PlaceIGS(W, "C2")
	if len(r) != 2 || r[0] != "B1" || r[1] != "C1" {
		t.Error("capture expected", r)
	}
}

func BenchmarkBoardString(b *testing.B) {
	a := model.NewBoard(model.X9)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = a.String()
	}
}

var endlessPlaySetup = []struct {
	c model.Color
	p model.Point
}{
	{c: B, p: P(1, 2)},
	{c: B, p: P(2, 1)},
	{c: B, p: P(3, 2)},
	{c: W, p: P(1, 3)},
	{c: W, p: P(2, 2)},
	{c: W, p: P(3, 3)},
}
var endlessPlayLoop = []struct {
	c model.Color
	p model.Point
}{
	{c: W, p: P(1, 1)},
	{c: B, p: P(2, 3)},
	{c: W, p: P(3, 1)},
	{c: B, p: P(1, 2)},
	{c: W, p: P(3, 3)},
	{c: B, p: P(2, 1)},
	{c: W, p: P(1, 3)},
	{c: B, p: P(3, 2)},
}

func ExampleBoard_endlessPlay() {
	b := model.NewBoard(model.Size(3))
	// setup
	for _, a := range endlessPlaySetup {
		_, _ = b.Place(a.c, a.p)
	}
	fmt.Println(b)
	// loop
	for _, a := range endlessPlayLoop {
		_, _ = b.Place(a.c, a.p)
	}
	fmt.Println(b)
	// Output:
	// +abc+
	// a ● |
	// b●◯●|
	// c◯ ◯|
	// {3 3} w0 b0
	// +abc+
	// a ●+|
	// b●◯●|
	// c◯ ◯|
	// {3 2} w4 b4
}

func BenchmarkBoardLoop(b *testing.B) {
	board := model.NewBoard(model.Size(3))
	for _, a := range endlessPlaySetup {
		_, _ = board.Place(a.c, a.p)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, a := range endlessPlayLoop {
			_, _ = board.Place(a.c, a.p)
		}
	}
}
