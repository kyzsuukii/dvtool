package types

type AuthValidator struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ActionValidator struct {
	Name string `form:"shell" binding:"required"`
}
