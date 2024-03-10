package revenue

import (
	"backend/internal/dto"
	"backend/internal/interfaces"
	"backend/internal/transport/http/handlers/middleware"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type revenueRoutes struct {
	usecase interfaces.Usecase
}

func NewRevenueRoutes(routerGroup *gin.RouterGroup, usecase interfaces.Usecase) {
	router := &revenueRoutes{
		usecase: usecase,
	}

	group := routerGroup.Group("/revenue")
	group.Use(middleware.Auth())

	group.POST("/", router.CreateRevenue)
	group.GET("/", router.GetRevenues)
}

// CreateRevenue create revenue
// @Tags revenue
// @Summary Retrieve revenue id
// @Security ApiKeyAuth
// @Description Retrieve revenue id
// @Param request body dto.Revenue true "request"
// @Success 200 {object} dto.Response "Successful operation"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /revenue [post]
func (r *revenueRoutes) CreateRevenue(c *gin.Context) {

	var body dto.Revenue
	err := c.ShouldBind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: fmt.Sprintf("Неверный формат данных %v", err),
		})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			ErrorCode:   "INTERNAL_SERVER_ERROR",
			Description: "Пользователь не распознан",
		})
	}
	body.UserId = userId.(string)
	id, err := r.usecase.Revenue().CreateRevenue(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.Response{
		Payload: dto.CreateRevenueResponse{
			Id: id,
		},
	})

}

// GetRevenues returns revenues
// @Tags revenue
// @Summary Retrieve revenues
// @Security ApiKeyAuth
// @Description Retrieve revenues
// @Param page query int false "0"
// @Param size query int false "10"
// @Success 200 {object} dto.Response "Successful operation"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /revenue [get]
func (r *revenueRoutes) GetRevenues(c *gin.Context) {

	var (
		pageString string
		sizeString string
	)
	pageString = c.Query("page")
	sizeString = c.Query("size")

	page, err := strconv.ParseUint(pageString, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: fmt.Sprintf("Неверный формат данных для page %v", err),
		})
		return
	}
	size, err := strconv.ParseUint(sizeString, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: fmt.Sprintf("Неверный формат данных для size %v", err),
		})
		return
	}

	revenues, total, err := r.usecase.Revenue().GetRevenues(c.Request.Context(), dto.GetRevenuesRequest{
		Pagination: dto.Pagination{
			Page: page,
			Size: size,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.Response{
		Payload: dto.GetRevenuesResponse{
			Revenues: revenues,
			Total:    total,
		},
	})

}
