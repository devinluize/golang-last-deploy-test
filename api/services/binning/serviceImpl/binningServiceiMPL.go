package serviceImpl

import (
	"context"
	"gorm.io/gorm"
	"user-services/api/entities/binningEntities"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads/request"
	"user-services/api/payloads/response"
	"user-services/api/repositories/binning"
	binning2 "user-services/api/services/binning"
)

type BinningServiceImpl struct {
	DB   *gorm.DB
	Repo binning.BinningRepo
}

func NewBinningServiceImpl(DB *gorm.DB, repo binning.BinningRepo) binning2.BinningService {
	return &BinningServiceImpl{DB: DB, Repo: repo}
}
func ItemToHeaderResponse(detail binningEntities.BinningStockDetail) response.BinningDetailResponses {
	return response.BinningDetailResponses{
		BinDocNo:  detail.BinDocNo,
		BinLineNo: detail.BinLineNo,
		PoLineNo:  detail.PoLineNo,
		ItemCode:  detail.ItemCode,
		LocCode:   detail.LocCode,
		CaseNo:    detail.CaseNo,
		GrpoQty:   detail.GrpoQty,
	}
}
func ItemToHeaderResponses(dataDetail []binningEntities.BinningStockDetail) []response.BinningDetailResponses {
	var detail []response.BinningDetailResponses
	for _, i := range dataDetail {
		detail = append(detail, ItemToHeaderResponse(i))
	}
	return detail
}
func ToHeaderResponse(data binningEntities.BinningStockHeader) response.BinningHeaderResponses {
	return response.BinningHeaderResponses{
		CompanyCode: data.CompanyCode,
		PoDocNo:     data.PoDocNo,
		WHSGroup:    data.WHSGroup,
		WHSCode:     data.WHSCode,
		Item:        ItemToHeaderResponses(data.Item),
	}
}
func ToHeaderResponses(binningHeader []binningEntities.BinningStockHeader) []response.BinningHeaderResponses {
	var binningResponses []response.BinningHeaderResponses
	for _, i := range binningHeader {

		binningResponses = append(binningResponses, ToHeaderResponse(i))
	}
	return binningResponses
}
func (b *BinningServiceImpl) FindAll(ctx context.Context, RequestBody []request.BinningHeaderRequest) ([]response.BinningHeaderResponses, error, string) {
	//TODO implement me
	allBining, err, errmsg := b.Repo.FindAll(ctx, b.DB, RequestBody)
	return ToHeaderResponses(allBining), err, errmsg
}

func (b *BinningServiceImpl) GetAll(ctx context.Context, RequestBody []request.BinningHeaderRequest) ([]response.BinningHeaderResponses, *exceptions.BaseErrorResponse) {
	tx := b.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := b.Repo.GetAll(tx, RequestBody)

	if err != nil {
		return ToHeaderResponses(get), err
	}

	return ToHeaderResponses(get), nil
}
