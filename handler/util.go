package handler

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinHandlerWrapper(f func(c *gin.Context) *domain.Response) gin.HandlerFunc {
	return func(c *gin.Context) {
		if res := f(c); res.Err != nil {
			ErrorResponse(c, res.Code, res.Err)
		} else {
			switch res.Data.(type) {
			case *util.Page:
				PaginationSuccessResponse(c, res.Data.(*util.Page))
			case domain.LoginResponse:
				TokenSuccessResponse(c, res.Data.(domain.LoginResponse))
			default:
				SuccessResponse(c, res.Data)
			}
		}
	}
}

func ErrorResponse(c *gin.Context, code int, e error) {
	result := map[string]interface{}{
		"isSuccess": false,
		"data":      nil,
		"error":     e.Error(),
	}
	c.JSON(code, result)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	result := map[string]interface{}{
		"isSuccess": true,
		"data":      data,
		"error":     nil,
	}
	c.JSON(http.StatusOK, result)
}

func PaginationSuccessResponse(c *gin.Context, data *util.Page) {
	result := map[string]interface{}{
		"isSuccess":    true,
		"data":         data.Data,
		"error":        nil,
		"current_page": data.CurrentPage,
		"total":        data.Total,
		"total_page":   data.TotalPage,
	}
	c.JSON(http.StatusOK, result)
}

func TokenSuccessResponse(c *gin.Context, data domain.LoginResponse) {
	result := map[string]interface{}{
		"isSuccess": true,
		"token":     data.Token,
		"error":     nil,
	}
	c.JSON(http.StatusOK, result)
}

func queryParamOrNull(req *http.Request, key string) interface{} {
	if value := req.FormValue(key); value != "" {
		return value
	}
	return nil
}

func queryLikeParamOrNull(req *http.Request, key string) interface{} {
	likeParam := queryParamOrNull(req, key)
	if likeParam != nil {
		return "%" + likeParam.(string) + "%"
	} else {
		return "%%"
	}
}
