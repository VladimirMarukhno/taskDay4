package customer

import (
	"days4/pkg/input"
	"fmt"
	"strconv"
)

type Customer struct {
	FirstName       string
	LastName        string
	Patronymic      string
	TelephoneNumber string
}

func NewCustomer() *Customer {
	var customer []string

	fmt.Println("Введите фамилию, имя, отчество и контактный телефон.")

	for {
		customer = input.Input()
		if len(customer) != 4 {
			continue
		}

		_, err := strconv.Atoi(customer[3])
		if err != nil {
			fmt.Println("Некорректно введён номер дома")
			continue
		} else {
			break
		}

	}
	return &Customer{
		FirstName:       customer[0],
		LastName:        customer[1],
		Patronymic:      customer[2],
		TelephoneNumber: customer[3],
	}

}
