package revenue

import (
	"backend/internal/dto"
	"backend/internal/interfaces"
	"backend/internal/transport/http/handlers/middleware"
	"net/http"

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
}

// CreateRevenue returns revenue id
// @Tags revenue
// @Summary Retrieve revenue id
// @Description Retrieve revenue id
// @Param request body dto.CreateRevenueRequest true "request"
// @Success 200 {object} dto.Response "Successful operation"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /revenue [post]
func (r *revenueRoutes) CreateRevenue(c *gin.Context) {

	var body dto.CreateRevenueRequest
	err := c.ShouldBind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: "Неверный формат данных",
		})
		return
	}
	userId, ok := c.Get("userID")
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
