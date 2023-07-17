package services

import "goredis/repositories"

type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(productRepo repositories.ProductRepository) CatalogService {
	return catalogService{productRepo}
}

func (s catalogService) GetProducts() (products []Product, err error) {
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

	return products, nil
}
