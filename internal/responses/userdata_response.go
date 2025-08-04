package responses

import "github.com/vimudakorn/internal/domain"

type UserUserdataResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AdminUserdataResponse struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Role   string         `json:"role"`
	Cart   *domain.Cart   `json:"cart"`
	Orders []domain.Order `json:"orders"`
}
