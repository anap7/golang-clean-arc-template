package routes

import (
	"net/http"
	"github.com/anap7/golang-clean-arc-template/app/controller"
)


var transactionRoutes = []Route{
	{
		URI:    "/transaction",
		Method: http.MethodPost,
		Function: controller.Create,
	},
}