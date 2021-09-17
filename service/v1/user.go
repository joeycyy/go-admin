package v1

import "go-admin/store"

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create()
	Delete()
	DeleteListCollection()
	Update()
	Get()
	List()
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{
		store: srv.store,
	}
}

func (u *userService) Create() {

}
func (u *userService) Delete() {

}
func (u *userService) DeleteListCollection() {

}
func (u *userService) Update() {

}
func (u *userService) Get() {

}
func (u *userService) List() {

}
