package produk

import (
	"context"
	"database/sql"
	"fmt"
	"fullstack_toko/backend/model/domain"
)

type Repository interface {
	GetAllProduct(ctx context.Context, tx *sql.Tx) ([]domain.Product, error)
	GetById(ctx context.Context, tx *sql.Tx, id int) (*domain.Product, error)
	CreateProduct(ctx context.Context, tx *sql.Tx, pr *domain.Product) (int, error)
	CreateProductStat(ctx context.Context, tx *sql.Tx, ps *domain.ProductStat) error
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetAllProduct(ctx context.Context, tx *sql.Tx) ([]domain.Product, error) {
	rows, err := tx.QueryContext(ctx, "select * from product;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dp := []domain.Product{}
	for rows.Next() {
		pr := &domain.Product{}
		err := rows.Scan(&pr.Id,  &pr.Url_image,&pr.Name,  &pr.Deskripsi, &pr.Category,&pr.Price, &pr.Quantity, &pr.Userid)
		if err != nil {
			return nil, err
		}
		dp = append(dp, *pr)
	}
	return dp, nil
}

func (r *RepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (*domain.Product, error) {
	pr := &domain.Product{}
	err := tx.QueryRowContext(ctx, "select * from product where id = ?", id).Scan(&pr.Id,  &pr.Url_image,&pr.Name,  &pr.Deskripsi, &pr.Category,&pr.Price, &pr.Quantity, &pr.Userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cannot find requested id")
		}
		return nil, err
	}
	return pr, nil
}

func (r *RepositoryImpl) CreateProduct(ctx context.Context, tx *sql.Tx, pr *domain.Product) (int, error) {
	result, err := tx.ExecContext(ctx, "insert into product (name,price,quantity,deskripsi,category,userid,url_image) values(?,?,?,?,?,?,?)",pr.Name,pr.Price,pr.Quantity,pr.Deskripsi,pr.Category,pr.Userid,pr.Url_image)
	if err != nil {
		return 0, err
	}
	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return 0, fmt.Errorf("no rows affected ,invalid userid")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

func (r *RepositoryImpl) CreateProductStat(ctx context.Context, tx *sql.Tx, ps *domain.ProductStat) error {
	_, err := tx.ExecContext(ctx, "insert into product_Stat(createdAT,lastUpdated,productID) values(?,?,?)", ps.CreatedAT, ps.LastUpdate, ps.ProductID)
	if err != nil {
		return err
	}
	return nil
}
