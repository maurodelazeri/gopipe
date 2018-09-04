package country

import "github.com/maurodelazeri/gopipe/server-gateway/service/common"

type VisitRequest struct {
	Country string `json:"country"`
}

type VisitResponse struct {
	Message string `json:"message"`
	common.ErrResponse
}
