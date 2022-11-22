package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	MenuRepository interface {
		GetAllMenus(pageable util.Pageable) (*util.Page, error)
		CreateMenu(menu *domain.MenuPayload) *domain.Response
		UpdateMenu(menu *domain.MenuPayload, menu_id int) *domain.Response
		DeleteMenu(menu_id int) *domain.Response
		GetMenu(menu_id int) *domain.Response
		UpdateMenuRating(menu_id int, rating float32) *domain.Response
	}

	menuRepository struct {
		db *gorm.DB
	}
)

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{
		db: db,
	}
}

func (repo *menuRepository) GetAllMenus(pageable util.Pageable) (*util.Page, error) {
	var count int64
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
	}

	if err := repo.db.Model(domain.Menu{}).Where("name ILIKE ?", arguments[0].(string)).
		Count(&count).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Menus{}), err
	}

	if count == 0 {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Menus{}), nil
	}

	paginator := util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), int(count))
	arguments = append(arguments, pageable.SortBy(), paginator.PerPageNums, paginator.Offset())

	var menus domain.Menus

	if err := repo.db.Preload("Category").Where("name ILIKE ?", arguments[0].(string)).Order(arguments[1]).
		Limit(arguments[2].(int)).Offset(arguments[3].(int)).Find(&menus).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Menus{}), err
	}

	return paginator.Pageable(menus), nil
}

func (repo *menuRepository) CreateMenu(menu *domain.MenuPayload) *domain.Response {
	if err := repo.db.Table("menus").Create(&menu).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(menu, 0, nil)
}

func (repo *menuRepository) UpdateMenu(menu *domain.MenuPayload, menu_id int) *domain.Response {
	if err := repo.db.Table("menus").Where("id = ?", menu_id).Updates(&menu).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(menu, 0, nil)
}

func (repo *menuRepository) DeleteMenu(menu_id int) *domain.Response {
	menu := new(domain.Menu)

	if err := repo.db.Clauses(clause.Returning{}).Delete(menu, menu_id).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(menu, 0, nil)
}

func (repo *menuRepository) GetMenu(menu_id int) *domain.Response {
	menu := new(domain.Menu)

	if err := repo.db.Where("id = ?", menu_id).First(&menu).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(menu, 0, nil)
}

func (repo *menuRepository) UpdateMenuRating(menu_id int, rating float32) *domain.Response {
	if err := repo.db.Model(&domain.Menu{}).Where("id = ?", menu_id).UpdateColumns(map[string]interface{}{
		"rating":       gorm.Expr("(rating + ?) / (total_review + 1)", rating),
		"total_review": gorm.Expr("total_review + 1"),
	}).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseMenuUpdated, 0, nil)
}
