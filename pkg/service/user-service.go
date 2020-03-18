package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	constants "bandi.com/main/data"
	"bandi.com/main/pkg/data"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		fmt.Println("Could not load metadata")
	}

	if ctx.Value(constants.OrgID) != nil {
		fmt.Println("Value of ORGID:", ctx.Value(constants.OrgID).(string))
	} else {
		fmt.Println("Could not load OrgID")
	}

	value, exists := userCache[req.Name]

	if !exists {
		return nil, errors.New(fmt.Sprintf("user %s doesn't exist", req.Name))
	}

	if false {
		st := status.New(codes.NotFound, fmt.Sprintf("user %s doesn't exist", req.Name))
		desc := "The username Doesn't exist, please give a valid username"

		/*
			customProtoError := &data.Error{
				Message:         "Check username",
				Code:            1404,
				Type:            "Skyflow",
				DetailedMessage: desc,
			}
		*/
		v := &errdetails.BadRequest_FieldViolation{
			Field:       "username",
			Description: desc,
		}
		customProtoError := &errdetails.BadRequest{}
		customProtoError.FieldViolations = append(customProtoError.FieldViolations, v)

		st, err := st.WithDetails(customProtoError)

		if err != nil {
			panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}
		return nil, st.Err()
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
	fmt.Println("Streaming user: ", req)

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
	fmt.Println("Completed Streaming user: ", req)

	return nil
}
