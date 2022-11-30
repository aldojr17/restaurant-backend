package domain

import (
	"fmt"
	"strings"
)

type PageableRequest struct {
	Filters   map[string]interface{}
	Search    map[string]interface{}
	Limit     int
	Page      int
	Sort_by   string
	Desceding string
	Type      string
}

func (p *PageableRequest) SearchParams() map[string]interface{} {
	return p.Search
}
func (p *PageableRequest) FilterParams() map[string]interface{} {
	return p.Filters
}
func (p *PageableRequest) GetPage() int {
	return p.Page
}
func (p *PageableRequest) GetLimit() int {
	return p.Limit
}
func (p *PageableRequest) SortBy() string {
	if p.Sort_by == "" {
		return "name ASC"
	}

	switch p.Type {
	case "menu":
		if strings.ToLower(p.Sort_by) == "date" {
			p.Sort_by = "created_at"
			break
		}

		if strings.ToLower(p.Sort_by) != "price" && strings.ToLower(p.Sort_by) != "name" && strings.ToLower(p.Sort_by) != "rating" {
			return "name ASC"
		}
	case "order":
		if strings.ToLower(p.Sort_by) != "order_date" && strings.ToLower(p.Sort_by) != "total_price" {
			return "order_date DESC"
		}
	default:
		return "name ASC"
	}

	if p.Desceding == "asc" {
		return fmt.Sprintf("%s %s", strings.ToLower(p.Sort_by), "ASC")
	}

	return fmt.Sprintf("%s %s", strings.ToLower(p.Sort_by), "DESC")
}
