package main

import (
	"errors"
	"fmt"
	"os"
)

// struct - blueprint, describes a custom data type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// method - has receiver argument of type bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	total := 0.0

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%.2f\n", key+":", value)
		total += value
	}

	// tip
	fs += fmt.Sprintf("%-25v ...$%.2f\n", "tip:", b.tip)
	total += b.tip

	// total
	fs += fmt.Sprintf("%-25v ...$%.2f", "total:", total)

	return fs
}

// pointer receiver - update struct value
// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill to file inside bills folder
func (b *bill) save() {
	data := []byte(b.format())

	// create bills folder if it doesn't exist
	if _, err := os.Stat("bills/"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("bills", 0777)
		if err != nil {
			panic(err)
		}
	}

	// dump byte slice into new text file
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}
