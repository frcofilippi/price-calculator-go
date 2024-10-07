package commandmanager

import "fmt"

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Welcome!!")
	fmt.Println("Please enter the priced and confirm using ENTER on each price!!")
	var price string
	prices := make([]string, 0)
	for {
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)

	}
	return prices, nil
}

func (cmd CMDManager) WriteJSON(data interface{}) error {
	fmt.Printf("Data: %v", data)
	return nil
}
