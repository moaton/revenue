package users

import (
	"backend/internal/dto"
	"backend/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usersRoutes struct {
	usecase interfaces.Usecase
}

func NewUsersRoutes(routerGroup *gin.RouterGroup, usecase interfaces.Usecase) {
	router := &usersRoutes{
		usecase: usecase,
	}

	group := routerGroup.Group("/users")

	group.POST("", router.Login)
	group.POST("/sign-up", router.SignUp)
}

// SignUp returns bool
// @Tags user
// @Summary Retrieve bool
// @Description Retrieve bool
// @Param request body dto.SignUpRequest true "request"
// @Success 200 {object} dto.Response "Successful operation"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /users/sign-up [post]
func (r *usersRoutes) SignUp(c *gin.Context) {

	var body dto.SignUpRequest
	err := c.ShouldBind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: "Неверный формат данных",
		})
		return
	}

	err = r.usecase.Users().SignUp(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.Response{
		Payload: dto.SignUpResponse{
			IsSucceed: true,
		},
	})
}

// Login returns user token
// @Tags user
// @Summary Retrieve user token
// @Description Retrieve user token
// @Param request body dto.LoginRequest true "request"
// @Success 200 {object} dto.Response "Successful operation"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /users [post]
func (r *usersRoutes) Login(c *gin.Context) {

	var body dto.LoginRequest
	err := c.ShouldBind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: "Неверный формат данных",
		})
		return
	}

	token, err := r.usecase.Users().Login(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			ErrorCode:   "BAD_REQUEST",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &dto.Response{
		Payload: dto.LoginResponse{
			Token: token,
		},
	})
}
