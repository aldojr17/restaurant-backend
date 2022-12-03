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
		CreateOrder(order *domain.OrderPayload) *domain.Response
		CreateOrderDetails(orders *domain.OrderDetails) *domain.Response
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
	var err error
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
		pageable.FilterParams()[util.FILTER_BY_CATEGORY],
	}

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Model(domain.Order{}).Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Where("menus.category_id = ?", arguments[1]).
			Where("user_id = ?", user_id).Count(&count).Error
	} else {
		err = repo.db.Model(domain.Order{}).Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).
			Where("user_id = ?", user_id).Count(&count).Error
	}

	if err != nil {
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

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Preload("Payment").Preload("OrderDetails.MenuDetail.Category").
			Preload("OrderDetails.MenuDetail.MenuOption").
			Preload("OrderDetails.MenuDetail.Reviews").
			Preload("Coupon").
			Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Where("menus.category_id = ?", arguments[1]).Where("user_id = ?", user_id).Order(arguments[2]).Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&orders).Error
	} else {
		err = repo.db.Preload("Payment").Preload("OrderDetails.MenuDetail.Category").
			Preload("OrderDetails.MenuDetail.MenuOption").
			Preload("OrderDetails.MenuDetail.Reviews").
			Preload("Coupon").
			Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Where("user_id = ?", user_id).Order(arguments[2]).Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&orders).Error
	}

	if err != nil {
		return util.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable(domain.Orders{}), err
	}

	return paginator.Pageable(orders), nil
}

func (repo *orderRepository) GetAllOrders(pageable util.Pageable) (*util.Page, error) {
	var count int64
	var err error
	arguments := []interface{}{
		pageable.SearchParams()[util.SEARCH_BY_NAME],
		pageable.FilterParams()[util.FILTER_BY_CATEGORY],
	}

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Model(domain.Order{}).Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Where("menus.category_id = ?", arguments[1]).
			Count(&count).Error
	} else {
		err = repo.db.Model(domain.Order{}).Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).
			Count(&count).Error
	}

	if err != nil {
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

	if arguments[1] != nil && arguments[1] != "0" {
		err = repo.db.Preload("Payment").Preload("OrderDetails.MenuDetail.Category").
			Preload("Coupon").
			Preload("User").
			Preload("OrderDetails.MenuDetail.MenuOption").
			Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Where("menus.category_id = ?", arguments[1]).
			Order(arguments[2]).Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&orders).Error
	} else {
		err = repo.db.Preload("Payment").Preload("OrderDetails.MenuDetail.Category").
			Preload("Coupon").
			Preload("User").
			Preload("OrderDetails.MenuDetail.MenuOption").
			Joins("left join order_details on order_details.order_id  = orders.id").
			Joins("left join menus on menus.id  = order_details.menu_id").
			Joins("left join categories on categories.id  = menus.category_id").
			Group("orders.id").
			Where("COALESCE(menus.name, '') ILIKE ?", arguments[0]).Order(arguments[2]).Limit(arguments[3].(int)).Offset(arguments[4].(int)).Find(&orders).Error
	}

	if err != nil {
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

func (repo *orderRepository) CreateOrder(order *domain.OrderPayload) *domain.Response {
	if err := repo.db.Table("orders").Create(&order).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(order, 0, nil)
}

func (repo *orderRepository) CreateOrderDetails(orders *domain.OrderDetails) *domain.Response {
	if err := repo.db.Table("order_details").Create(&orders.Orders).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseOrderDetailsCreated, 0, nil)
}
