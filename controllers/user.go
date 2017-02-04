package controllers

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"

	"zqc/middlewares"
	"zqc/services"
)

type UserInfoParams struct {
	Id string `valid:"objectidhex"`
}

func UserInfo(c echo.Context) (err error) {
	cc := c.(*middlewares.Context)
	params := UserInfoParams{
		Id: cc.FormValue("id"),
	}
	if ok, err := valid.ValidateStruct(params); !ok {
		return services.NewServiceError(services.ErrCodeInvalidParams, err.Error())
	}
	id, err := services.ParseObjectId(params.Id)
	if err != nil {
		return services.NewServiceError(services.ErrCodeInvalidParams, err.Error())
	}

	user, err := services.GetUser(id)
	if err != nil {
		return err
	}

	return ResponseJSON(http.StatusOK, Response{
		Data: map[string]interface{}{
			"user": user,
		},
	}, cc)
}

type UserInfosParams struct {
	Ids string `valid:"stringlength(24|2400)"`
}

func UserInfos(c echo.Context) (err error) {
	cc := c.(*middlewares.Context)
	params := UserInfosParams{
		Ids: cc.FormValue("ids"),
	}
	if ok, err := valid.ValidateStruct(params); !ok {
		return services.NewServiceError(services.ErrCodeInvalidParams, err.Error())
	}
	ids, err := services.ParseObjectIds(params.Ids)
	if err != nil {
		return services.NewServiceError(services.ErrCodeInvalidParams, err.Error())
	}

	users, err := services.GetUsers(ids)
	if err != nil {
		return err
	}

	return ResponseJSON(http.StatusOK, Response{
		Data: map[string]interface{}{
			"users": users,
		},
	}, cc)
}
