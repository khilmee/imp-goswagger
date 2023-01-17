package repository

import (
	"errors"
	"imp-goswagger/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepo struct {
	base BaseRepository
}

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	Update(id int, input *model.Product) (*model.Product, error)
	FindById(id int) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Delete(id int) error
}

func NewProductRepository(br BaseRepository) ProductRepository {
	return &productRepo{br}
}

func (r *productRepo) Create(product *model.Product) (*model.Product, error) {
	err := r.base.GetDB().
		Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepo) Update(id int, input *model.Product) (*model.Product, error) {
	result := &model.Product{}
	err := r.base.GetDB().Model(result).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(input).
		Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *productRepo) FindById(id int) (*model.Product, error) {
	var product model.Product
	err := r.base.GetDB().
		Where("id=?", id).
		First(&product).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) FindAll() ([]model.Product, error) {
	var product []model.Product
	query := r.base.GetDB()
	err := query.Find(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

func (r *productRepo) Delete(id int) error {
	var product model.Product
	err := r.base.GetDB().
		Where("id = ?", id).
		Delete(&product).
		Error
	if err != nil {
		return err
	}
	return nil
}
