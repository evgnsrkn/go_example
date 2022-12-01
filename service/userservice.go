package service

import (
	"go_example/db"
	"go_example/model"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	users = []model.User{}
	user  model.User
	lock  = sync.Mutex{}
)

func Index(c echo.Context) error {
	rows := db.DbManager().Find(&users)

	return c.JSON(http.StatusOK, rows)
}

func GetById(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	//Get id from param
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Try to find user
	err = db.DbManager().First(&user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func GetAll(c echo.Context) error {
	//Lock goroutine
	lock.Lock()
	defer lock.Unlock()

	//Find all
	db.DbManager().Find(&users)

	return c.JSON(http.StatusOK, users)
}

func Create(c echo.Context) error {
	//Lock goroutine
	lock.Lock()
	defer lock.Unlock()

	//Try to bind request body to User model
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//If binding was successed, make new User
	user = model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	//Push user to database
	err := db.DbManager().Create(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func Update(c echo.Context) error {
	//Lock goroutine
	lock.Lock()
	defer lock.Unlock()

	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db.DbManager().Save(&user)

	return c.JSON(http.StatusOK, user)

}

func Delete(c echo.Context) error {
	//Lock goroutine
	lock.Lock()
	defer lock.Unlock()

	//Parse uuid
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Delete user
	db.DbManager().Delete(users, id)

	return c.NoContent(http.StatusNoContent)

}
