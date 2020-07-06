package server

import (
	"context"

	"github.com/benchhub/benchhub/bhpb"
)

var _ bhpb.UserServiceServer = (*UserService)(nil)

type UserService struct {
	bhpb.UnimplementedUserServiceServer
}

func newUserService() (*UserService, error) {
	return &UserService{}, nil
}

func (u UserService) GetUser(ctx context.Context, name *bhpb.IdOrName) (*bhpb.User, error) {
	log.Infof("GetUser %d %s", name.Id, name.Name)
	// TODO: fetch from database
	return &bhpb.User{
		Id:       15,
		Name:     "at15",
		FullName: "Pinglei Guo",
		Email:    "at15@outlook.com",
	}, nil
}
