package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/amburskui/httpserver/internal/application/userservice"
	"github.com/amburskui/httpserver/internal/domain"
)

func registerRoute(log *logrus.Logger, service *userservice.Service) http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())

	e.HTTPErrorHandler = customErrorHandler
	e.GET("/health", healthHandler)

	g := e.Group("/api/v1")

	userHandler := userHandler{log: log, s: service}

	u := g.Group("/user")
	u.POST("", userHandler.Create)
	u.GET("/:id", userHandler.Get)
	u.PUT("/:id", userHandler.Update)
	u.DELETE("/:id", userHandler.Delete)

	return e.Server.Handler
}

func healthHandler(c echo.Context) error {
	data := map[string]string{"status": "OK"}

	return c.JSON(http.StatusOK, data)
}

type userHandler struct {
	s   *userservice.Service
	log *logrus.Logger
}

func (u *userHandler) Create(c echo.Context) error {
	var body struct {
		Username  string `json:"username"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	user, err := u.s.Create(body.Username, body.FirstName, body.LastName, body.Email, body.Phone)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (u *userHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := domain.UserIdentity(id)

	user, err := u.s.Get(userID)
	if err != nil {
		return fmt.Errorf("user by id %d %w", userID, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) Update(c echo.Context) error {
	var body struct {
		Username  string `json:"username"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := domain.UserIdentity(id)

	user, err := u.s.Update(userID, body.Username, body.FirstName, body.LastName, body.Email, body.Phone)
	if err != nil {
		return fmt.Errorf("user by id %d %w", userID, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := domain.UserIdentity(id)

	if err := u.s.Delete(userID); err != nil {
		return fmt.Errorf("user by id %d %w", userID, err)
	}

	return c.NoContent(http.StatusNoContent)
}
