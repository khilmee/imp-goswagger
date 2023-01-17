package model

// swagger:model Product
type Product struct {
	// ID of product
	// in: int64
	ID uint64 `gorm:"primary_key:auto_increment" json:"-"`

	// Name of product
	// in: string
	Name string `json:"name" gorm:"not null;type:varchar" faker:"first_name"`

	// SKU of product
	// in: string
	SKU string `json:"sku" gorm:"type:varchar" faker:"time_period"`

	// UOM of product
	// in: string
	UOM string `json:"uom" gorm:"type:varchar" faker:"uuid_digit"`

	// Weight of product
	// in: int32
	Weight int32 `json:"weight" faker:"oneof: 15, 27, 61"`
}

// swagger:parameters SaveProductRequest
type ReqProductBody struct {
	//  in: body
	Body SaveProductRequest `json:"body"`
}

// swagger:parameters UpdateProductRequest
type ReqSampleProductBodyUpdate struct {
	// name: id
	// in: path
	// required: true
	ID int `json:"id"`
	//  in: body
	Body SaveProductRequest `json:"body"`
}

type SaveProductRequest struct {
	// Name of product
	// in: string
	Name string `json:"name"`

	// SKU of product
	// in: string
	SKU string `json:"sku"`

	// UOM of the product
	// in: string
	UOM string `json:"uom"`

	// Weight of the product
	// in: int32
	Weight int32 `json:"weight"`
}

// swagger:parameters byParamDelete
type GetProductByParamDelete struct {
	// name: id
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters byParamGet
type GetProductByParamGet struct {
	// name: id
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model CommonSuccess
type CommonSuccess struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}
