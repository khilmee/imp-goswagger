package service

import (
	"imp-goswagger/app/model"
	"imp-goswagger/app/repository"
	"imp-goswagger/helper/message"
)

type ProductService interface {
	CreateProduct(input model.SaveProductRequest) (*model.Product, message.Message)
	UpdateProduct(id int, input model.SaveProductRequest) (*model.Product, message.Message)
	GetProduct(id int) (*model.Product, message.Message)
	GetList() ([]model.Product, message.Message)
	DeleteProduct(id int) message.Message
}

type productServiceImpl struct {
	baseRepo    repository.BaseRepository
	productRepo repository.ProductRepository
}

func NewProductService(
	br repository.BaseRepository,
	pr repository.ProductRepository,
) ProductService {
	return &productServiceImpl{br, pr}
}

// swagger:operation POST /product Product SaveProductRequest
// Create Product
//
// ---
// responses:
//
//	401: CommonError
//	200: CommonSuccess
func (s *productServiceImpl) CreateProduct(input model.SaveProductRequest) (*model.Product, message.Message) {
	s.baseRepo.BeginTx()
	product := model.Product{
		Name:   input.Name,
		SKU:    input.SKU,
		UOM:    input.UOM,
		Weight: input.Weight,
	}
	result, err := s.productRepo.Create(&product)
	if err != nil {
		return nil, message.ErrSaveData
	}
	return result, message.SuccessMsg
}

// swagger:operation PUT /product/{id} Product SaveProductRequest
// Update Product
//
// ---
// responses:
//
//	401: CommonError
//	200: CommonSuccess
func (s *productServiceImpl) UpdateProduct(id int, input model.SaveProductRequest) (*model.Product, message.Message) {
	s.baseRepo.BeginTx()
	product, err := s.productRepo.FindById(id)
	if err != nil {
		return nil, message.ErrNotFound
	}
	if product != nil {
		return nil, message.ErrNotFound
	}

	product.Name = input.Name
	product.SKU = input.SKU
	product.UOM = input.UOM
	product.Weight = input.Weight

	result, err := s.productRepo.Update(id, product)
	if err != nil {
		return nil, message.ErrSaveData
	}
	return result, message.SuccessMsg
}

// swagger:operation GET /product/{id} Product byParamGet
// Get Product by Id
//
// ---
// responses:
//
//	401: CommonError
//	200: CommonSuccess
func (s *productServiceImpl) GetProduct(id int) (*model.Product, message.Message) {
	s.baseRepo.BeginTx()
	result, err := s.productRepo.FindById(id)
	if err != nil {
		return nil, message.ErrNotFound
	}
	return result, message.SuccessMsg
}

// swagger:operation GET /product Product getList
// Get Product List
//
// ---
// responses:
//
//	401: CommonError
//	200: CommonSuccess
func (s *productServiceImpl) GetList() ([]model.Product, message.Message) {
	s.baseRepo.BeginTx()
	results, err := s.productRepo.FindAll()
	if err != nil {
		return nil, message.ErrNotFound
	}
	return results, message.SuccessMsg
}

// swagger:operation DELETE /product/{id} Product byParamDelete
// Delete Product by Id
//
// ---
// responses:
//
//	401: CommonError
//	200: CommonSuccess
func (s *productServiceImpl) DeleteProduct(id int) message.Message {
	s.baseRepo.BeginTx()
	err := s.productRepo.Delete(id)
	if err != nil {
		return message.ErrDeleteData
	}
	return message.SuccessMsg
}
