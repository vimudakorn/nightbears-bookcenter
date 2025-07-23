package responses

import "github.com/vimudakorn/internal/domain"

type UserUserdataResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AdmidUserdataResponse struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Role     string         `json:"role"`
	Phone    string         `json:"phone"`
	Cart     *domain.Cart   `json:"cart"`
	Orders   []domain.Order `json:"orders"`
}
