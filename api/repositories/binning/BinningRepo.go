package binning

import (
	"context"
	"gorm.io/gorm"
	"user-services/api/entities/binningEntities"
	"user-services/api/exceptions"
	"user-services/api/payloads/request"
)

type BinningRepo interface {
	FindAll(ctx context.Context, db *gorm.DB, requestBody []request.BinningHeaderRequest) ([]binningEntities.BinningStockHeader, error, string)
	GetAll(db *gorm.DB, requestBody []request.BinningHeaderRequest) ([]binningEntities.BinningStockHeader, *exceptions.BaseErrorResponse)
}
