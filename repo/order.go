package repo

import "github.com/bobisme/some-go-app/model"

type OrderRepo interface {
	Fetch(id int) model.Order
	Save(order *model.Order) (int, error)
}
