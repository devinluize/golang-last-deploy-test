package repoImpl

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"user-services/api/entities/binningEntities"
	"user-services/api/exceptions"
	"user-services/api/payloads/request"
	"user-services/api/repositories/binning"
)

type BinningRepositoryImpl struct{}

func NewBinningRepositoryImpl() binning.BinningRepo {
	return &BinningRepositoryImpl{}
}
func GetBinningHeader(db *gorm.DB, RequestHeader request.BinningHeaderRequest, data *binningEntities.BinningStockHeader) error {
	errs := db.Where("VIA_BINNING = ? AND PO_DOC_NO = ? AND COMPANY_CODE = ?", "1", RequestHeader.PoDocNo, RequestHeader.CompanyCode).First(&data)
	return errs.Error
}
func getBinningDetail(db *gorm.DB, RequestHeader request.BinningHeaderRequest, data *[]binningEntities.BinningStockDetail) error {
	errs := db.Table("atbinningstock0 A").
		Select("A.BIN_DOC_NO, B.BIN_LINE_NO, B.REF_LINE AS PO_LINE_NO, B.ITEM_CODE, B.LOC_CODE, A.SUPPLIER_CASE_NO AS CASE_NO, CAST(B.DO_QTY AS FLOAT) AS GRPO_QTY").
		Joins("INNER JOIN atbinningstock1 B ON A.BIN_SYS_NO = B.BIN_SYS_NO").
		Where("REF_DOC_NO = ? AND COMPANY_CODE = ?", RequestHeader.PoDocNo, RequestHeader.CompanyCode).
		Find(&data)
	if errors.Is(errs.Error, gorm.ErrRecordNotFound) {
		recover()
		return nil
	}
	return errs.Error
}
func (c *BinningRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, requestBody []request.BinningHeaderRequest) ([]binningEntities.BinningStockHeader, error, string) {
	//TODO implement me
	tx := db.Begin()
	defer tx.Rollback()
	var BinningHeader []binningEntities.BinningStockHeader
	var err error = nil
	var errOnDetail error = nil
	for _, i := range requestBody {
		var stock binningEntities.BinningStockHeader
		var detail []binningEntities.BinningStockDetail
		err = GetBinningHeader(db, i, &stock)
		if err != nil {
			panic(err)
		}
		errOnDetail = getBinningDetail(db, i, &detail)
		if err != nil || errOnDetail != nil {
			tx.Rollback()
			continue
		}
		stock.Item = detail
		BinningHeader = append(BinningHeader, stock)
	}
	if err != nil || errOnDetail != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return BinningHeader, err, "Record Not Found"
		}
		return BinningHeader, err, "Error Reading to Database"
	}
	return BinningHeader, nil, ""
}

func (c *BinningRepositoryImpl) GetAll(tx *gorm.DB, requestBody []request.BinningHeaderRequest) ([]binningEntities.BinningStockHeader, *exceptions.BaseErrorResponse) {
	//var approval []approvalentities.Approval
	//err := tx.
	//	Model(&approval).
	//	Scopes(utils.Paginate(approval, &pagination, tx.Model(&approval))).
	//	Scan(&approval).
	//	Error
	defer tx.Rollback()
	var BinningHeader []binningEntities.BinningStockHeader
	var err error = nil
	var errOnDetail error = nil
	for _, i := range requestBody {
		var stock binningEntities.BinningStockHeader
		var detail []binningEntities.BinningStockDetail
		err = GetBinningHeader(tx, i, &stock)
		if err != nil {
			panic(err)
		}
		errOnDetail = getBinningDetail(tx, i, &detail)
		if err != nil || errOnDetail != nil {
			tx.Rollback()
			continue
		}
		stock.Item = detail
		BinningHeader = append(BinningHeader, stock)
	}
	if err != nil {
		return BinningHeader, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if errOnDetail != nil {
		return BinningHeader, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return BinningHeader, nil
}
