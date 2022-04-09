package common

// GetByIDRequest DTO to get address by id
type GetByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}
