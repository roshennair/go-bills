package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')     // read string input until newline character
	return strings.TrimSpace(input), err // remove leading/trailing whitespace
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // create buffered reader for standard input

	name, _ := getInput("Enter a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, t - add tip, s - save bill): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		p, _ := getInput("Item price: $", reader)

		price, err := strconv.ParseFloat(p, 64) // convert string to float64
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, price)
		fmt.Println("Item added -", name, price)
		promptOptions(b)
	case "t":
		t, _ := getInput("Tip amount: $", reader)

		tip, err := strconv.ParseFloat(t, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(tip)
		fmt.Println("Tip added -", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("File saved -", b.name+".txt")
	default:
		fmt.Println("That wasn't a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
