package main

import (
	"days4/pkg/address"
	"days4/pkg/billLading"
	"days4/pkg/customer"
	"days4/pkg/product"
	"days4/pkg/repo/repoBillLading"
	"fmt"
	"log"
)

/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/
// Заполняем хранилище накладными
func addBill() repoBillLading.BillLadingStorage {
	r := repoBillLading.NewCustomerStorage()
	var id uint = 1
	resume := "N"
	for {
		bill := billLading.NewBillLading(id)
		id++
		as := address.NewAddress()
		bill.AddAddress(as)
		cs := customer.NewCustomer()
		bill.AddCustomer(cs)
		ps := product.NewCustomer()
		bill.AppendProd(ps)
		r.Put(bill)
		fmt.Print("Вы хотите ввести данные ещё одной накладной?\n" +
			"Да[Y] Нет[N]\n")
		_, err := fmt.Scan(&resume)
		if err != nil {
			log.Println(err)
		}
		if resume == "Y" {
			continue
		} else if resume == "N" {
			break
		} else {
			log.Fatal("Недопустимый символ")
		}
	}
	return r
}

// Запрашиваем вывод накладных.
func show(r *repoBillLading.BillLadingStorage) {
	input := "N"
	for {
		fmt.Println("===========================================================================")
		fmt.Println("Вы хотите посмотреть все накладные или определённую накладную по её id?")
		fmt.Println("Все[Y] Одну[Z] Выход[N]")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("===========================================================================")
		if input == "N" {
			fmt.Println("Операция отменена!")
			break
		} else if input == "Y" {
			for _, c := range *r {
				fmt.Printf("\nid: %v\nФИО:%v %v %v\nТелефонный номер: %v\nАдресс доставки: %v, %v, %v, %v, %v\nid товара: %v\nНаименование товара: %v\n\n",
					c.Id, c.Customer.FirstName, c.Customer.LastName, c.Customer.Patronymic, c.Customer.TelephoneNumber,
					c.Address.PostalCode, c.Address.State, c.Address.City, c.Address.Street, c.Address.House,
					c.Prod.NumberItems, c.Prod.ProductName)
			}
		} else {
			fmt.Println("Введите id накладной")
			var tmp uint
			_, err = fmt.Scan(&tmp)
			if err != nil {
				log.Fatal(err)
			}
			c, err := r.Get(tmp)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Printf("\nid: %v\nФИО:%v %v %v\nТелефонный номер: %v\nАдресс доставки: %v, %v, %v, %v, %v\nid товара: %v\nНаименование товара: %v\n\n",
					c.Id, c.Customer.FirstName, c.Customer.LastName, c.Customer.Patronymic, c.Customer.TelephoneNumber,
					c.Address.PostalCode, c.Address.State, c.Address.City, c.Address.Street, c.Address.House,
					c.Prod.NumberItems, c.Prod.ProductName)
			}
		}
	}
}

func main() {
	repo := addBill()
	show(&repo)
}
