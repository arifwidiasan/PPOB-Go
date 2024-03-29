package domain

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)

	CreateProductType(product_type model.Product_type) error
	GetAllProductType() []model.Product_type
	GetProductTypeByID(id int) (product_type model.Product_type, err error)
	UpdateProductTypeByID(id int, product_type model.Product_type) error
	DeleteProductTypeByID(id int) error

	CreateOperator(operator model.Operator) error
	GetAllOperator() []model.Operator
	GetOperatorByID(id int) (operator model.Operator, err error)
	UpdateOperatorByID(id int, operator model.Operator) error
	DeleteOperatorByID(id int) error

	CreateProduct(product model.Product) error
	GetProductByName(name string) (product model.Product, err error)
	UpdateProductByID(id int, product model.Product) error
	GetAllProduct() []model.Product
	GetProductByID(id int) (product model.Product, err error)
	DeleteProductByID(id int) error
	GetProductByProductType(id int) ([]model.Product, error)
	GetProductByOperator(id int) ([]model.Product, error)

	CreateUser(user model.User) error
	GetAllUser() []model.User
	GetUserByID(id int) (user model.User, err error)
	UpdateUserByID(id int, user model.User) error
	DeleteUserByID(id int) error
	CheckLoginUser(input string) (user model.User, err error)
	GetUserByUsername(username string) (user model.User, err error)

	CreatePaymentMethod(payment_method model.Payment_method) error
	GetAllPaymentMethod() []model.Payment_method
	GetPaymentMethodByID(id int) (payment_method model.Payment_method, err error)
	UpdatePaymentMethodByID(id int, payment_method model.Payment_method) error
	DeletePaymentMethodByID(id int) error

	CreateTransaction(transaction model.Transaction) error
	GetTransactionByCodeTransaction(code_transaction string) (transaction model.Transaction, err error)
	UpdateTransactionByID(id int, transaction model.Transaction) error
	GetAllTransaction() []model.Transaction
	GetTransactionByID(id int) (transaction model.Transaction, err error)
	DeleteTransactionByID(id int) error
	GetAllUserTransaction(id int) []model.Transaction

	CreateVirtualAccount(virtual_account model.Virtual_account) error
	CheckVirtualAccount(user_id, payment_method_id int) (virtual_account model.Virtual_account, err error)
	UpdateVirtualAccountByID(id int, virtual_account model.Virtual_account) error
	GetVirtualAccountByExternalID(external_id string) (virtual_account model.Virtual_account, err error)

	CreateCallbackPayment(callback_payment model.Callback_payment) error
	GetAllCallbackPayment() []model.Callback_payment

	GetCountTransaction() (count int64, err error)
	GetCountTransactionSuccess() (count int64, err error)
	GetCountProduct() (count int64, err error)
	GetCountUser() (count int64, err error)
}

type AdapterService interface {
	LoginAdmin(username, password string) (string, int)
	GetAdminByUsernameService(username string) (model.Admin, error)

	ClaimToken(bearer *jwt.Token) string

	CreateProductTypeService(product_type model.Product_type) error
	GetAllProductTypeService() []model.Product_type
	GetProductTypeByIDService(id int) (model.Product_type, error)
	UpdateProductTypeByIDService(id int, product_type model.Product_type) error
	DeleteProductTypeByIDService(id int) error

	CreateOperatorService(operator model.Operator) error
	GetAllOperatorService() []model.Operator
	GetOperatorByIDService(id int) (model.Operator, error)
	UpdateOperatorByIDService(id int, operator model.Operator) error
	DeleteOperatorByIDService(id int) error

	CreateProductService(product model.Product) error
	GetAllProductService() []model.Product
	GetProductByIDService(id int) (model.Product, error)
	UpdateProductByIDService(id int, product model.Product) error
	DeleteProductByIDService(id int) error
	GetProductByProductTypeService(id int) ([]model.Product, error)
	GetProductByOperatorService(id int) ([]model.Product, error)

	CreateUserService(user model.User) error
	GetAllUserService() []model.User
	GetUserByIDService(id int) (model.User, error)
	UpdateUserByIDService(id int, user model.User) error
	DeleteUserByIDService(id int) error
	LoginUser(input, password string) (string, int)
	GetUserByUsernameService(username string) (model.User, error)

	CreatePaymentMethodService(payment_method model.Payment_method) error
	GetAllPaymentMethodService() []model.Payment_method
	GetPaymentMethodByIDService(id int) (model.Payment_method, error)
	UpdatePaymentMethodByIDService(id int, payment_method model.Payment_method) error
	DeletePaymentMethodByIDService(id int) error

	CreateTransactionService(transaction model.Transaction) error
	GetAllTransactionService() []model.Transaction
	GetTransactionByIDService(id int) (model.Transaction, error)
	UpdateTransactionByIDService(id int, transaction model.Transaction) error
	DeleteTransactionByIDService(id int) error
	GetAllUserTransactionService(id int) []model.Transaction
	GetTransactionByCodeTransactionService(code_transaction string) (model.Transaction, error)

	CreateVirtualAccountService(transaction model.Transaction) error
	GetVirtualAccountService(external_id string) (model.Virtual_account, error)

	CreateCallbackPaymentService(callback_payment model.Callback_payment) error
	GetAllCallbackPaymentService() []model.Callback_payment

	SendEmailService(email string, transaction model.Transaction, virtual_account model.Virtual_account) error

	SimulatePaymentService(code_transaction string, price int) error

	GetCountTransactionService() (int64, error)
	GetCountTransactionSuccessService() (int64, error)
	GetCountProductService() (int64, error)
	GetCountUserService() (int64, error)
}
