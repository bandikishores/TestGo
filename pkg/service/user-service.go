package service

import (
	"fmt"

	"bandi.com/main/pkg/data"
	"golang.org/x/net/context"
)

var userCache map[string]*data.User = make(map[string]*data.User)

// UserService helps manage users
type UserService struct {
	userName string
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(ctx context.Context, req *data.CreateUserRequest) (*data.CreateUserResponse, error) {
	fmt.Println("Creating user: ", req)

	_, exists := userCache[req.User.Name]

	if exists {
		return nil, fmt.Errorf("user %s already exists", req.User.Name)
	}

	userCache[req.User.Name] = req.User

	return &data.CreateUserResponse{
		Name: req.User.Name,
	}, nil
}

// GetUser Gets existing user
func (us *UserService) GetUser(ctx context.Context, req *data.GetUserRequest) (*data.GetUserResponse, error) {
	fmt.Println("Getting user: ", req)

	value, exists := userCache[req.Name]

	if !exists {
		return nil, fmt.Errorf("user %s doesn't exist", req.Name)
	}

	return &data.GetUserResponse{
		User: value,
	}, nil
}

// DeleteUser Deletes existing user
func (us *UserService) DeleteUser(ctx context.Context, req *data.DeleteUserRequest) (*data.DeleteUserResponse, error) {
	fmt.Println("Deleting user: ", req)

	_, exists := userCache[req.Name]

	if !exists {
		return nil, fmt.Errorf("user %s doesn't exist", req.Name)
	}

	delete(userCache, req.Name)

	return &data.DeleteUserResponse{
		Name: req.Name,
	}, nil
}
