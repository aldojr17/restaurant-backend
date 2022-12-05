package handler

import (
	"final-project-backend/domain"
	"final-project-backend/initialize"
	"final-project-backend/mocks"
	"final-project-backend/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewCouponHandler(t *testing.T) {
	app := initialize.App()
	couponHandler := NewCouponHandler(app)

	assert.NotNil(t, couponHandler)
}

func TestCreateCoupon(t *testing.T) {
	payload := &domain.Coupon{
		Code:       "TEST01",
		Discount:   10000,
		ValidUntil: "2022-12-12",
	}

	jsonBody := `{
		"code":"TEST01",
		"discount":10000,
		"valid_until":"2022-12-12"
		}`

	body := strings.NewReader(jsonBody)

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("POST", "/admin/coupons", body)

	c.Request = req
	c.Set("role_id", 0)

	s := mocks.NewCouponService(t)
	h := couponHandler{
		s: s,
	}

	s.On("CreateCoupon", payload).Return(util.SetResponse(nil, 0, nil))

	res := h.CreateCoupon(c)
	assert.Nil(t, res.Err)
}
