package model

import "fmt"

// Komi value
type Komi float64

func (k Komi) String() string {
	return fmt.Sprintf("%.1f", k)
}
