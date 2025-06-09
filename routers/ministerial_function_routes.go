package routers

import (
	"church_ministerial_functions/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupMinisterialFunctionRoutes(app *fiber.App, mfHandler *handler.MinisterialFunctionHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Post("/ministerial-functions", mfHandler.CreateMinisterialFunction)
		v1.Get("/ministerial-functions/:id", mfHandler.GetMinisterialFunctionByID)
		v1.Get("/ministerial-functions", mfHandler.GetAllMinisterialFunctions)
		v1.Put("/ministerial-functions/:id", mfHandler.UpdateMinisterialFunction)
		v1.Delete("/ministerial-functions/:id", mfHandler.DeleteMinisterialFunction)
	}
}
