package routers

import (
	"devbook/src/controllers"
	"net/http"
)

var userRouters = []Router{
	//? Rota de usuários
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
