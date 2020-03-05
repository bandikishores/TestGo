package service

import (
	"fmt"
	"strings"
	"time"

	"bandi.com/main/pkg/data"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
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

	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		val := md[strings.ToLower("X-Custom-orgname")]
		if val == nil {
			fmt.Println("Header X-Custom-orgname not found ")
		} else {
			fmt.Println("Value of X-Custom-orgname Retrieved was : ", val[0])
		}
	} else {
		fmt.Errorf("Could not load metadata")
	}

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

// UpdateUser Updates existing user
func (us *UserService) UpdateUser(ctx context.Context, req *data.UpdateUserRequest) (*data.UpdateUserResponse, error) {
	fmt.Println("Updating user: ", req)

	_, exists := userCache[req.Name]

	if !exists {
		return nil, fmt.Errorf("user %s doesn't exist", req.Name)
	}

	userCache[req.Name] = req.User

	return &data.UpdateUserResponse{
		Name: req.Name,
	}, nil
}

// StreamUsers Streams existing user with manual delay for demo
func (us *UserService) StreamUsers(req *data.GetUserRequest, stream data.UserService_StreamUsersServer) error {
	fmt.Println("Getting user: ", req)

	value, exists := userCache[req.Name]

	if !exists {
		fmt.Printf(fmt.Sprintf("user %s doesn't exist", req.Name))
		return fmt.Errorf("user %s doesn't exist", req.Name)
	}

	stream.Send(&data.GetUserResponse{
		User: value,
	})

	time.Sleep(3000 * time.Millisecond)

	stream.Send(&data.GetUserResponse{
		User: &data.User{
			Name: "After Sleep, Dummy Data",
		},
	})

	stream.Context().Done()
	fmt.Println("Completed Getting user: ", req)

	return nil
}
