package order

import (
	"context"
	"database/sql"
	"fmt"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/domain"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
	"time"
)

type ServiceImpl struct {
	repo Repository
	Db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
		Db:   db,
	}
}

func (s *ServiceImpl) CreateOrders(ctx context.Context, userID int) (int, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer exception.CommitOrRollback(tx)

	OrderID, err := s.repo.CreateOrders(ctx, tx, &domain.Orders{
		LastUpdated: time.Now(),
		UserID:      userID,
	})
	if err != nil {
		return 0, err
	}
	return OrderID, nil
}

func (s *ServiceImpl) CreateOrderitems(ctx context.Context, ord *web.OrderitemsCreatePayload) error {

	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer exception.CommitOrRollback(tx)

	err = s.repo.CreateOrderitems(ctx, tx, &domain.Order_items{
		Status:     ord.Status,
		Quantity:   ord.Quantity,
		TotalPrice: ord.TotalPrice,
		ProductID:  ord.ProductID,
		OrderID:    ord.OrderID,
		UserID:     ord.UserID,
	})

	return err

}

func (s *ServiceImpl) SearchOrders(ctx context.Context, userID int) ([]web.Orders, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	ord, err := s.repo.SearchOrders(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	if len(ord) < 1 {
		return nil, fmt.Errorf("you dont have any order")
	}
	return utils.ConvertOrdersIntoSlice(ord), nil

}

func (s *ServiceImpl) SearchOrderItems(ctx context.Context, userID int, orderID []int) ([]web.Order_items, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	ord, err := s.repo.SearchOrderItems(ctx, tx, userID, orderID)
	if err != nil {
		return nil, err
	}
	if len(ord) < 1 {
		return nil, fmt.Errorf("you dont have any order")
	}

	return utils.ConvertOrderItemsIntoSlice(ord), nil
}
