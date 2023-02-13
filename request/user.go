package request

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email"`
	Password string `form:"password" binding:"required"`
	Code     string `form:"code"`
}

type SendMailVerify struct {
	Email string `form:"email" binding:"required"`
}

type ChangeUsername struct {
	OldUsername string `form:"oldUsername" binding:"required"`
	NewUsername string `form:"newUsername" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type ChangePassword struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required"`
}

type BanUser struct {
	Username string `form:"username" binding:"required"`
}

type DeleteUser struct {
	Username string `form:"username" binding:"required"`
}
