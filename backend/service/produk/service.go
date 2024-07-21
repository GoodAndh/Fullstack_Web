package produk

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
	Repo Repository
	Db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) *ServiceImpl {
	return &ServiceImpl{
		Repo: repo,
		Db:   db,
	}
}

func (s *ServiceImpl) GetAllProduct(ctx context.Context) ([]web.Product, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	pr, err := s.Repo.GetAllProduct(ctx, tx)
	if err != nil {
		return nil, err
	}

	if len(pr) < 1 {
		return nil, fmt.Errorf("no product found")
	}

	return utils.ConvertProductIntoSlice(pr), nil
}

func (s *ServiceImpl) GetById(ctx context.Context, id int) (*web.Product, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	pr, err := s.Repo.GetById(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return utils.ConvertProductIntoWeb(pr), nil
}

func (s *ServiceImpl) CreateProduct(ctx context.Context, pr *web.ProductCreatePayload) (int, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer exception.CommitOrRollback(tx)

	prID, err := s.Repo.CreateProduct(ctx, tx, &domain.Product{
		Name:      pr.Name,
		Deskripsi: pr.Deskripsi,
		Category:  pr.Category,
		Price:     float64(pr.Price),
		Quantity:  pr.Quantity,
		Userid:    pr.Userid,
		Url_image: pr.Url_image,
	})
	if err != nil {
		return 0, err
	}

	return prID, nil
}

func (s *ServiceImpl) CreateProductStat(ctx context.Context, productID int) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer exception.CommitOrRollback(tx)

	err = s.Repo.CreateProductStat(ctx, tx, &domain.ProductStat{
		CreatedAT:  time.Now(),
		LastUpdate: time.Now(),
		ProductID:  productID,
	})

	return err

}

func (s *ServiceImpl) UpdateProduct(ctx context.Context, ps *web.ProductUpdatePayload) error {

	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer exception.CommitOrRollback(tx)

	err = s.Repo.UpdateProduct(ctx, tx, &domain.Product{
		Id:        ps.Id,
		Name:      ps.Name,
		Deskripsi: ps.Deskripsi,
		Category:  ps.Category,
		Price:     float64(ps.Price),
		Quantity:  ps.Quantity,
		Userid:    ps.Userid,
		Url_image: ps.Url_image,
	})

	return err
}

func (s *ServiceImpl) GetByUserID(ctx context.Context, userID int) ([]web.Product, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	ps, err := s.Repo.GetByUserID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	return utils.ConvertProductIntoSlice(ps), nil
}



func checkIfIDisExist(ps []web.Product,productID int)error {
	psMap:=make(map[int]web.Product)
	for _,product  := range ps {
		psMap[product.Id]=product
	}

	// check if id exist
	_,ok:=psMap[productID]
	if !ok{
		return fmt.Errorf("cannot find the id")
	}else {
		return nil
	}
}