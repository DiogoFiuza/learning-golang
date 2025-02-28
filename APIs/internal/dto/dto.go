package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
