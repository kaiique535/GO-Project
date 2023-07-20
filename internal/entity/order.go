package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) *Order {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validade()
	if err != nil {
		return nil
	}
	return order
}

func (o *Order) Validade() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	if o.Tax < 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() { //* Equivale ao ponteiro, ponteiros são utilizados para atualizar variaveis fora do escopo declarado.
	o.FinalPrice = o.Price + o.Tax

}
