package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(productSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{productSrv, redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "handlers::GetProducts"

	if productsJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(productsJson)
	}

	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if data, err := json.Marshal(products); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*15).Err()
	}

	res := fiber.Map{
		"status":   "ok",
		"products": products,
	}
	fmt.Println("databse")
	return c.JSON(res)
}
