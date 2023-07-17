package handlers

import (
	"goredis/services"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {

	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// simple use
	// c.JSON(products)

	res := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	return c.JSON(res)
}
