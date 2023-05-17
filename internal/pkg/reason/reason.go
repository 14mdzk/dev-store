package reason

var (
	EntityNotFound      = ":entity not found"
	EntityCreateSuccess = "successfully create :entity"
	EntityCreateFailed  = "failed to create :entity"
	EntityUpdateSuccess = "successfully update :entity"
	EntityUpdateFailed  = "failed to update :entity"
	EntityDeleteSuccess = "successfully delete :entity"
	EntityDeleteFailed  = "failed to delete :entity"

	CategoryNotFound      = "category not found"
	CategoryCreateSuccess = "successfully create category"
	CategoryCreateFailed  = "failed to create category"
	CategoryUpdateSuccess = "successfully update category"
	CategoryUpdateFailed  = "failed to update category"
	CategoryDeleteSuccess = "successfully delete category"
	CategoryDeleteFailed  = "failed to delete category"

	ProductNotFound      = "product not found"
	ProductCreateSuccess = "successfully create product"
	ProductCreateFailed  = "failed to create product"
	ProductUpdateSuccess = "successfully update product"
	ProductUpdateFailed  = "failed to update product"
	ProductDeleteSuccess = "successfully delete product"
	ProductDeleteFailed  = "failed to delete product"

	RegistrationSuccess = "registration success"
	RegistrationFailed  = "failed to register"
	UserAlreadyExist    = "user already exists"
	LoginFailed         = "failed login, please make sure your email and password match"

	RefreshTokenFailed = "failed to refresh token, please check your token"
	Unauthorized       = "doesn't have authorization to access this resource"
	InvalidAccess      = "you don't have access to this resource"

	InternalServerError = "internal server error"
	RequestFormError    = "request format is invalid"
)
