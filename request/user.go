package request

type Login struct {
	Email    string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Email    string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Code     string `form:"code"`
}

type SendMailVerify struct {
	Email string `form:"email" binding:"required"`
}

type ChangeUsername struct {
	OldEmail string `form:"oldEmail" binding:"required"`
	NewEmail string `form:"newEmail" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ChangePassword struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required"`
}

type BanUser struct {
	Email string `form:"username" binding:"required"`
}

type DeleteUser struct {
	Email string `form:"username" binding:"required"`
}
