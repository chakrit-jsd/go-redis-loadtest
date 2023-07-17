package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productRepo, redisClient}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "services::GetProducts"

	productsJson, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		if json.Unmarshal([]byte(productsJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	productsDB, err := s.productRepo.GetProduct()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	if err = s.redisClient.Set(context.Background(), key, string(data), time.Second*15).Err(); err != nil {
		return nil, err
	}

	fmt.Println("database")
	return products, nil
}
