package menurepoimpl

import (
	"net/http"
	menuentities "user-services/api/entities/master/menu"
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
	menurepo "user-services/api/repositories/master/menu"

	"gorm.io/gorm"
)

type MenuUrlRepositoryImpl struct {
}

// GetMenuURLByCompanyAndUser implements menurepo.MenuUrlRepository.
func (*MenuUrlRepositoryImpl) GetByRoleID(tx *gorm.DB, roleID int) ([]string, *exceptions.BaseErrorResponse) {
	var menuAccess menuentities.MenuAccess
	var response []string

	err := tx.
		Preload("MenuListByAccess.MenuUrl", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id",
				"path",
			)
		}).
		Select("id").
		Find(&menuAccess, "role_id = ?", roleID).
		Error

	if err != nil {

		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	for i := range menuAccess.MenuListByAccess {
		response = append(response,
			menuAccess.MenuListByAccess[i].MenuUrl.Path,
		)
	}

	return response, nil
}

func NewMenuURLRepository() menurepo.MenuUrlRepository {
	return &MenuUrlRepositoryImpl{}
}

func (*MenuUrlRepositoryImpl) Create(tx *gorm.DB, request []menupayloads.CreateMenuUrlRequest) (bool, *exceptions.BaseErrorResponse) {
	var menuUrlEntities []menuentities.MenuUrl

	for i := range request {
		menuUrlEntities = append(menuUrlEntities, menuentities.MenuUrl{
			Path:        request[i].MenuUrlPath,
			Description: request[i].MenuUrlDescription,
		})
	}

	rows, err := tx.
		Create(&menuUrlEntities).
		Rows()

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return true, nil
}

func (*MenuUrlRepositoryImpl) GetByName(tx *gorm.DB, menuUrlPath []string) (menupayloads.GetMenuUrlByName, *exceptions.BaseErrorResponse) {
	var menuUrlEntities menuentities.MenuUrl
	var response menupayloads.GetMenuUrlByName

	rows, err := tx.Model(&menuUrlEntities).
		Where("path in (?)", menuUrlPath).
		Scan(&response).
		Rows()

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return response, nil
}
