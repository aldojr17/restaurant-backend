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
		UpdateMenuRating(menu_id int, data map[string]interface{}) *domain.Response
		AddMenuOption(options *[]domain.MenuOption) *domain.Response
		UpdateMenuOption(option *domain.MenuOption) *domain.Response
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
	var err error
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
		pageable.FilterParams()[util.FILTER_BY_CATEGORY],
	}

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Model(domain.Menu{}).Where("name ILIKE ?", arguments[0].(string)).
			Where("category_id = ?", arguments[1]).
			Count(&count).Error
	} else {
		err = repo.db.Model(domain.Menu{}).Where("name ILIKE ?", arguments[0].(string)).
			Count(&count).Error
	}

	if err != nil {
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

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Preload("Category").Preload("MenuOption").Preload("Reviews.User").
			Where("name ILIKE ?", arguments[0].(string)).Where("category_id = ?", arguments[1]).Order(arguments[2]).
			Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&menus).Error
	} else {
		err = repo.db.Preload("Category").Preload("MenuOption").Preload("Reviews.User").
			Where("name ILIKE ?", arguments[0].(string)).Order(arguments[2]).
			Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&menus).Error
	}

	if err != nil {
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

	if err := repo.db.Preload("Category").Preload("MenuOption").Preload("Reviews.User").Where("id = ?", menu_id).First(&menu).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, domain.ErrMenuNotFound)
	}

	return util.SetResponse(menu, 0, nil)
}

func (repo *menuRepository) UpdateMenuRating(menu_id int, data map[string]interface{}) *domain.Response {
	if err := repo.db.Model(&domain.Menu{}).Where("id = ?", menu_id).UpdateColumns(data).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseMenuUpdated, 0, nil)
}

func (repo *menuRepository) AddMenuOption(options *[]domain.MenuOption) *domain.Response {
	if err := repo.db.Create(&options).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(options, 0, nil)
}

func (repo *menuRepository) UpdateMenuOption(option *domain.MenuOption) *domain.Response {
	if err := repo.db.Table("menu_options").Updates(&option).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(option, 0, nil)
}
