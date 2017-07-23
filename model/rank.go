package model

import (
	"encoding/json"
)

// Rank sorted as int
type Rank int

// rank constants  NR = 0, 30 Kyu = 1, ...
const (
	NR Rank = iota
	Kyu30
	Kyu29
	Kyu28
	Kyu27
	Kyu26
	Kyu25
	Kyu24
	Kyu23
	Kyu22
	Kyu21
	Kyu20
	Kyu19
	Kyu18
	Kyu17
	Kyu16
	Kyu15
	Kyu14
	Kyu13
	Kyu12
	Kyu11
	Kyu10
	Kyu9
	Kyu8
	Kyu7
	Kyu6
	Kyu5
	Kyu4
	Kyu3
	Kyu2
	Kyu1
	Dan1
	Dan2
	Dan3
	Dan4
	Dan5
	Dan6
	Dan7
	Dan8
	Dan9
	Pro1
	Pro2
	Pro3
	Pro4
	Pro5
	Pro6
	Pro7
	Pro8
	Pro9
)

var (
	ranks  = make(map[string]Rank)
	labels = []string{
		NR:    "NR",
		Kyu30: "30 Kyu",
		Kyu29: "29 Kyu",
		Kyu28: "28 Kyu",
		Kyu27: "27 Kyu",
		Kyu26: "26 Kyu",
		Kyu25: "25 Kyu",
		Kyu24: "24 Kyu",
		Kyu23: "23 Kyu",
		Kyu22: "22 Kyu",
		Kyu21: "21 Kyu",
		Kyu20: "20 Kyu",
		Kyu19: "19 Kyu",
		Kyu18: "18 Kyu",
		Kyu17: "17 Kyu",
		Kyu16: "16 Kyu",
		Kyu15: "15 Kyu",
		Kyu14: "14 Kyu",
		Kyu13: "13 Kyu",
		Kyu12: "12 Kyu",
		Kyu11: "11 Kyu",
		Kyu10: "10 Kyu",
		Kyu9:  "9 Kyu",
		Kyu8:  "8 Kyu",
		Kyu7:  "7 Kyu",
		Kyu6:  "6 Kyu",
		Kyu5:  "5 Kyu",
		Kyu4:  "4 Kyu",
		Kyu3:  "3 Kyu",
		Kyu2:  "2 Kyu",
		Kyu1:  "1 Kyu",
		Dan1:  "1 Dan",
		Dan2:  "2 Dan",
		Dan3:  "3 Dan",
		Dan4:  "4 Dan",
		Dan5:  "5 Dan",
		Dan6:  "6 Dan",
		Dan7:  "7 Dan",
		Dan8:  "8 Dan",
		Dan9:  "9 Dan",
		Pro1:  "1 Pro",
		Pro2:  "2 Pro",
		Pro3:  "3 Pro",
		Pro4:  "4 Pro",
		Pro5:  "5 Pro",
		Pro6:  "6 Pro",
		Pro7:  "7 Pro",
		Pro8:  "8 Pro",
		Pro9:  "9 Pro",
	}
)

// init reverse rank map
func init() {
	for r, s := range labels {
		ranks[s] = Rank(r)
	}
}

// String rank, unknown returns NR
func (r Rank) String() string {
	if r <= 0 || int(r) >= len(labels) {
		return labels[NR]
	}
	return labels[r]
}

// MarshalJSON encodes rank
func (r Rank) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// ToRank decodes rank
func ToRank(r string) Rank {
	return ranks[r]
}
