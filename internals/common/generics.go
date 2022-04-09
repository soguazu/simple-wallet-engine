package common

import "github.com/soguazu/boilerplate_golang/pkg/utils"

// GetByIDRequest DTO to get address by id
type GetByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// GetAllResponse DTO get all companies
type GetAllResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    utils.Pagination `json:"data"`
}
