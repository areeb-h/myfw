package myfw

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Paginated struct for structured JSON response
type Paginated[T any] struct {
	Data       []T `json:"data"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// Paginate slices the data and automatically handles pagination
func Paginate[T any](items []T) Paginated[T] {
	c := fiber.Ctx{} // Dummy context to extract query params

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(items) {
		start = len(items)
	}
	if end > len(items) {
		end = len(items)
	}

	paginatedItems := items[start:end]
	totalPages := (len(items) + limit - 1) / limit

	return Paginated[T]{
		Data:       paginatedItems,
		Page:       page,
		Limit:      limit,
		TotalItems: len(items),
		TotalPages: totalPages,
	}
}
