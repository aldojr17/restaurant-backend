package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error)
		GetAllOrders(pageable util.Pageable) (*util.Page, error)
		UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response
	}

	orderRepository struct {
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (repo *orderRepository) GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error) {
	var count int64
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
	}

	if err := repo.db.Model(domain.Order{}).Where("user_id = ?", user_id).Count(&count).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), err
	}

	if count == 0 {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), nil
	}

	paginator := util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), int(count))
	arguments = append(arguments, pageable.SortBy(), paginator.PerPageNums, paginator.Offset())

	var orders domain.Orders

	if err := repo.db.Preload("MenuOptions.MenuDetail.Category").Where("user_id = ?", user_id).Order(arguments[1]).Limit(arguments[2].(int)).Offset(arguments[3].(int)).Find(&orders).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), err
	}

	return paginator.Pageable(orders), nil
}

func (repo *orderRepository) GetAllOrders(pageable util.Pageable) (*util.Page, error) {
	var count int64
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
	}

	if err := repo.db.Model(domain.Order{}).Count(&count).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), err
	}

	if count == 0 {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), nil
	}

	paginator := util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), int(count))
	arguments = append(arguments, pageable.SortBy(), paginator.PerPageNums, paginator.Offset())

	var orders domain.Orders

	if err := repo.db.Preload("MenuOptions.MenuDetail.Category").Order(arguments[1]).Limit(arguments[2].(int)).Offset(arguments[3].(int)).Find(&orders).Error; err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), err
	}

	return paginator.Pageable(orders), nil
}

func (repo *orderRepository) UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response {

	if err := repo.db.Table("orders").Where("id", order.Id).Updates(&order).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseOrderStatusUpdated, 0, nil)
}
