package user

import (
	srvv1 "go-admin/service/v1"
	"go-admin/store"
)

// defines UserHandler used to return a user handle resource
type UserController struct {
	srv srvv1.Service
}

// NewUserHandler return a UserHandler used to handle service resource
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
