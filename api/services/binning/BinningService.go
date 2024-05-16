package binning

import (
	"context"
	"user-services/api/exceptions"
	"user-services/api/payloads/request"
	"user-services/api/payloads/response"
)

type BinningService interface {
	FindAll(ctx context.Context, RequestBody []request.BinningHeaderRequest) ([]response.BinningHeaderResponses, error, string)
	GetAll(ctx context.Context, RequestBody []request.BinningHeaderRequest) ([]response.BinningHeaderResponses, *exceptions.BaseErrorResponse)
}
