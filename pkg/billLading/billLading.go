package billLading

import (
	"days4/pkg/address"
	"days4/pkg/customer"
	"days4/pkg/product"
)

type BillLading struct {
	Id       uint
	Prod     product.Product
	Address  address.Address
	Customer customer.Customer
}

func NewBillLading(id uint) *BillLading {
	return &BillLading{
		Id: id,
	}
}

func (b *BillLading) AppendProd(prod *product.Product) {
	b.Prod = *prod
}

func (b *BillLading) AddAddress(a *address.Address) {
	b.Address = *a
}

func (b *BillLading) AddCustomer(c *customer.Customer) {
	b.Customer = *c
}
