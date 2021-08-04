package service

import (
	"context"
	"fmt"
	pbfiles "shop-service/pbfiles"
)

type CatalogServiceServer struct {
}

func (rs *CatalogServiceServer) List(ctx context.Context, request *pbfiles.ViewCatalogRequest) (*pbfiles.ViewCatalogResponse, error) {
	var catType = request.GetCatalogType()

	// logic goes here (search from DB ...)
	resposeMsg := fmt.Sprintf("%s - %s", "Requested type is ", catType)	
	fmt.Println("Request is :",request)
	fmt.Println("Response msg is :",resposeMsg)
	return &pbfiles.ViewCatalogResponse{
		Code:     200,
		JsonString: resposeMsg,
	}, nil
}