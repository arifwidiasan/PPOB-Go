package docs

type LoginAdminSuccess struct {
	Messages string `json:"messages" example:"success"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQyNzA2OTAsImlkIjoxLCJ1c2VybmFtZSI6ImFkbWluIn0.9Qif8_Ug6Uy_oxDXIuIui3nepCPFc1jJ6mO6wEhiuHE"`
}

type CreateProductTypeSuccess struct {
	Messages string `json:"messages" example:"success"`
	Name     string `json:"product_type_name" example:"Pulsa"`
}

type CreateProductTypeFail struct {
	Messages string `json:"messages" example:"error insert product type"`
}
