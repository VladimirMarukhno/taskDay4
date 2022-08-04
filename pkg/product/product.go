package product

import (
	"days4/pkg/input"
	"fmt"
	"strconv"
)

type Product struct {
	ProductName string
	NumberItems string
}

func NewCustomer() *Product {
	fmt.Println("Введите наименование товара и его инвентарный номер.")

	var product []string

	for {

		product = input.Input()

		if len(product) != 2 {
			fmt.Println("Таблица заполнена неверно!")
			continue
		}

		_, err := strconv.Atoi(product[1])
		if err != nil {
			fmt.Println("Некорректно введён id товара")
			continue
		} else {
			break
		}
	}

	return &Product{
		ProductName: product[0],
		NumberItems: product[1],
	}
}
