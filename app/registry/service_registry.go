package registry

import (
	rp "imp-goswagger/app/repository"
	"imp-goswagger/app/service"

	"gorm.io/gorm"
)

func RegisterProductService(db *gorm.DB) service.ProductService {
	return service.NewProductService(
		rp.NewBaseRepository(db),
		rp.NewProductRepository(rp.NewBaseRepository(db)),
	)
}
