package address

import (
	"days4/pkg/input"
	"fmt"
	"strconv"
)

type Address struct {
	PostalCode string
	State      string
	City       string
	Street     string
	House      string
}

func NewAddress() *Address {
	fmt.Println("Введите почтовый индекс, государство, город, улицу и дом ")

	var address []string

	for {
		address = input.Input()
		if len(address) != 5 {
			fmt.Println("Адрес заполнен не верно!")
			continue
		}
		_, err := strconv.Atoi(address[4])
		if err != nil {
			fmt.Println("Некорректно введён номер дома")
			continue
		}
		_, err = strconv.Atoi(address[0])
		if err != nil {
			fmt.Println("Некорректно введён почтовый индекс")
			continue
		} else {
			break
		}
	}

	return &Address{
		PostalCode: address[0],
		State:      address[1],
		City:       address[2],
		Street:     address[3],
		House:      address[4],
	}
}
