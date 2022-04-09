package ports

import (
	"github.com/soguazu/boilerplate_golang/internals/core/domain"
)

// RequestDTO declaring input DTO
type RequestDTO interface {
	domain.Company
}
