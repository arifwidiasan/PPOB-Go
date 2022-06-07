package docs

type CreateProductType struct {
	Name string `json:"name" example:"Pulsa"`
}

type NewAdminLogin struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin123"`
}
