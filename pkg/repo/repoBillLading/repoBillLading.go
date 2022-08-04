package repoBillLading

import (
	"days4/pkg/billLading"
	"fmt"
)

type BillLadingStorage map[uint]*billLading.BillLading

func NewCustomerStorage() BillLadingStorage {
	return make(map[uint]*billLading.BillLading)
}

func (bs BillLadingStorage) Put(b *billLading.BillLading) {
	bs[b.Id] = b
}

func (bs BillLadingStorage) Get(billLadingKey uint) (*billLading.BillLading, error) {
	c, ok := bs[billLadingKey]
	if !ok {
		return nil, fmt.Errorf("Накладной не существует!")
	}
	return c, nil
}
