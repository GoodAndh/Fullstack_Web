package order

import (
	"context"
	"database/sql"
	"fmt"
	"fullstack_toko/backend/model/domain"
	"strings"
)

type Repository interface {
	CreateOrders(ctx context.Context, tx *sql.Tx, ord *domain.Orders) (int, error)
	CreateOrderitems(ctx context.Context, tx *sql.Tx, ord *domain.Order_items) error
	SearchOrders(ctx context.Context, tx *sql.Tx, userID int) ([]domain.Orders, error)
	SearchOrderItems(ctx context.Context, tx *sql.Tx, userID int, orderID []int) ([]domain.Order_items, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) CreateOrders(ctx context.Context, tx *sql.Tx, ord *domain.Orders) (int, error) {
	result, err := tx.ExecContext(ctx, "insert into orders (lastupdated,userID) values (?,?) ", ord.LastUpdated, ord.UserID)
	if err != nil {
		return 0, err
	}

	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return 0, fmt.Errorf("no rows affected ,message:%v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *RepositoryImpl) CreateOrderitems(ctx context.Context, tx *sql.Tx, ord *domain.Order_items) error {
	_, err := tx.ExecContext(ctx, "insert into order_items(status,quantity,totalprice,productID,orderID,userID) values(?,?,?,?,?,?)", ord.Status, ord.Quantity, ord.TotalPrice, ord.ProductID, ord.OrderID, ord.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) SearchOrders(ctx context.Context, tx *sql.Tx, userID int) ([]domain.Orders, error) {
	or := []domain.Orders{}
	rows, err := tx.QueryContext(ctx, "select * from orders where userID = ?", userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ord := &domain.Orders{}
		err := rows.Scan(&ord.Id, &ord.LastUpdated, &ord.UserID)
		if err != nil {
			return nil, err
		}
		or = append(or, *ord)
	}

	return or, nil

}

func (r *RepositoryImpl) SearchOrderItems(ctx context.Context, tx *sql.Tx, userID int, orderID []int) ([]domain.Order_items, error) {
	placeholders := strings.Repeat("?,", len(orderID))
	query := fmt.Sprintf("select * from order_items where userid = %d and orderid in (%s%d)", userID, placeholders, 0)
	//convert orderID into slice of any
	args := make([]any, len(orderID))
	for i, v := range orderID {
		args[i] = v
	}

	or := []domain.Order_items{}
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Next()
	for rows.Next() {
		ord := &domain.Order_items{}
		err := rows.Scan(&ord.Id, &ord.Status, &ord.Quantity, &ord.TotalPrice, &ord.ProductID, &ord.OrderID, &ord.UserID)
		if err != nil {
			return nil, err
		}
		or = append(or, *ord)
	}

	return or, nil
}
