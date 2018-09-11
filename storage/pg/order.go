package pg

import (
	"fmt"

	"github.com/bobisme/some-go-app/model"
	"github.com/jackc/pgx"
)

type OrderMapper struct{}

func (m OrderMapper) Columns() string {
	return "id,created_at,modified_at,deleted_at"
}

func (m OrderMapper) Fields(order *model.Order) []interface{} {
	return []interface{}{
		&order.ID,
		// &order.CreatedAt,
		// &order.ModifiedAt,
		// &order.DeletedAt,
	}
}

//--

type Orders struct {
	DB *pgx.ConnPool
}

func (o Orders) Fetch(id int) model.Order {
	panic("not implemented")
}

func (o Orders) Save(order *model.Order) (int, error) {
	orderMapper := OrderMapper{}
	err := o.DB.QueryRow(
		fmt.Sprintf(`
			INSERT INTO orders (created_at, modified_at)
			VALUES (NOW(), NOW())
			RETURNING %s
		`, orderMapper.Columns())).
		Scan(orderMapper.Fields(order)...)
	if err != nil {
		return -1, err
	}
	return order.ID, nil
}
