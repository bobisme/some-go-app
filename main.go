package main

import (
	"fmt"

	"github.com/bobisme/some-go-app/db"
	"github.com/bobisme/some-go-app/model"
	"github.com/bobisme/some-go-app/repo"
	"github.com/bobisme/some-go-app/storage/pg"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func saveOrder(repo repo.OrderRepo, order *model.Order) (int, error) {
	return repo.Save(order)
}

func main() {
	db, err := db.Init()
	panicIfErr(err)
	fmt.Println(db)
	order := model.Order{}
	orderRepo := pg.Orders{db}
	_, err = saveOrder(orderRepo, &order)
	panicIfErr(err)
	fmt.Println("order id is", order.ID)
}
