package cmdmanager

import "fmt"

type CMD struct{}

func (cmd CMD) WriteJson(data interface{}) error {
	fmt.Println(data)
	return nil
}

func (cmd CMD) ReadLines() ([]string, error) {

	var prices []string
	fmt.Println("Please Enter your prices. Confirm every price with enter")

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil

}

func New() CMD {
	return CMD{}
}
