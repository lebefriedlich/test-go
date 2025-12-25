package handler

import (
	"context"
	nethttp "net/http"

	"github.com/gofiber/fiber/v2"

	"my-go-project/internal/repository"
)

// RegisterMasterCategoryMerchantRoutes sets up master_category_merchant endpoints.
func RegisterMasterCategoryMerchantRoutes(app *fiber.App, repo *repository.MasterCategoryMerchantRepository) {
	app.Get("/master-category-merchants", func(c *fiber.Ctx) error {
		ctx := context.Background()
		records, err := repo.List(ctx)
		if err != nil {
			return fiber.NewError(nethttp.StatusInternalServerError, err.Error())
		}
		return c.JSON(records)
	})
}
