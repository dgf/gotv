package model_test

import (
	"encoding/json"
	"fmt"

	"github.com/dgf/gotv/model"
)

func ExampleDate_After() {
	today := model.Today()
	yesterday := today.Add(0, 0, -1)

	fmt.Printf("> %v\n", today.After(yesterday))
	fmt.Printf("< %v\n", yesterday.After(today))

	// Output:
	// > true
	// < false
}

func ExampleDate_Before() {
	today := model.Today()
	tomorrow := today.Add(0, 0, 1)

	fmt.Printf("< %v\n", today.Before(tomorrow))
	fmt.Printf("> %v\n", tomorrow.Before(today))

	// Output:
	// < true
	// > false
}

func ExampleDate_MarshalJSON() {
	date := model.ToDate("2017-07-23")
	data := map[string]model.Date{"date": date}

	if b, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		fmt.Print(string(b))
	}

	// Output:
	// {"date":"2017-07-23"}
}
