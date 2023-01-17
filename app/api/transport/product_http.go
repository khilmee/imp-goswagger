package transport

import (
	"context"
	"encoding/json"
	"imp-goswagger/app/api/endpoint"
	"imp-goswagger/app/model"
	"imp-goswagger/app/model/base"
	"imp-goswagger/app/service"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func ProductHttpHandler(s service.ProductService, logger log.Logger) http.Handler {
	pr := mux.NewRouter()

	ep := endpoint.MakeProductEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(base.EncodeError),
	}

	pr.Methods("POST").Path("/api/v1/product").Handler(httptransport.NewServer(
		ep.Create,
		decodeSaveProduct,
		base.EncodeResponseHTTP,
		options...,
	))

	return pr
}

func decodeSaveProduct(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req model.SaveProductRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
