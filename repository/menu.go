package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	MenuRepository interface {
		GetAllMenus(pageable util.Pageable) (*util.Page, error)
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
