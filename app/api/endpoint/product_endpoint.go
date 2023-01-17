package endpoint

import (
	"context"
	"imp-goswagger/app/model"
	"imp-goswagger/app/model/base"
	"imp-goswagger/app/service"

	"github.com/go-kit/kit/endpoint"
)

type ProductEndpoint struct {
	Create  endpoint.Endpoint
	Update  endpoint.Endpoint
	Show    endpoint.Endpoint
	FindAll endpoint.Endpoint
	Delete  endpoint.Endpoint
}

func MakeProductEndpoints(s service.ProductService) ProductEndpoint {
	return ProductEndpoint{
		Create:  makeSaveProduct(s),
		Update:  makeUpdateProduct(s),
		Show:    makeShowProduct(s),
		FindAll: makeFindAllProduct(s),
		Delete:  makeDeleteProduct(s),
	}
}

func makeSaveProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(model.SaveProductRequest)
		result, msg := s.CreateProduct(req)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, result), nil
	}
}

func makeUpdateProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(model.ReqSampleProductBodyUpdate)
		result, msg := s.UpdateProduct(req.ID, req.Body)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, result), nil
	}
}

func makeShowProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(model.GetProductByParamGet)
		result, msg := s.GetProduct(req.ID)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, result), nil
	}
}

func makeFindAllProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		result, msg := s.GetList()
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, result), nil
	}
}

func makeDeleteProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(model.GetProductByParamDelete)
		msg := s.DeleteProduct(req.ID)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, nil), nil
	}
}
