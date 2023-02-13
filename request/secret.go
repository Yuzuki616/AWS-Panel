package request

type AddSecret struct {
	Name   string `form:"name" binding:"required"`
	Id     string `form:"id" binding:"required"`
	Secret string `form:"secret" binding:"required"`
}

type GetSecret struct {
	Name string `form:"name" binding:"required"`
}

type DelSecret struct {
	Name string `form:"name" binding:"required"`
}
