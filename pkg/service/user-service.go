package service

import (
	"fmt"
	"strings"
	"time"

	constants "bandi.com/TestGo/data"
	"bandi.com/TestGo/pkg/data"
	myErrors "bandi.com/TestGo/pkg/error"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Error - sdf
type Error struct {
	Message         string `json:"message,omitempty"`
	Code            int32  `json:"code,omitempty"`
	Type            string `json:"type,omitempty"`
	DetailedMessage string `json:"detailed_message,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

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

	/*	if !exists {
		return nil, errors.New(fmt.Sprintf("user %s doesn't exist", req.Name))
	}*/

	if !exists {

		desc := "The username Doesn't exist, please give a valid username"

		customProtoError := &data.Error{
			Message:         "Check username",
			Code:            1404,
			Type:            "Skyflow",
			DetailedMessage: desc,
		}

		/*	st := status.Newf(codes.NotFound, "user %s doesn't exist", req.Name)
			st, err := st.WithDetails(customProtoError)
			if err != nil {
				panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
			}
			return nil, st.Err()*/

		/*v := &errdetails.BadRequest_FieldViolation{
			Field:       "username",
			Description: desc,
		}
		customProtoError := &errdetails.BadRequest{}
		customProtoError.FieldViolations = append(customProtoError.FieldViolations, v)
		*/
		//if err != nil {
		//	panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		//}

		customError := myErrors.NewCustomErrorf(codes.NotFound, "user %s doesn't exist", req.Name)
		v := &rpc.BadRequest_FieldViolation{
			Field:       "username",
			Description: desc,
		}
		br := &rpc.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		customError = customError.WithDetails(br)
		customError = customError.WithDetails(customProtoError)
		return nil, customError

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
		st := status.Newf(codes.NotFound, "user %s doesn't exist", req.Name)
		desc := "The username Doesn't exist, please give a valid username"

		customProtoError := &data.Error{
			Message:         "Check username",
			Code:            1404,
			Type:            "Skyflow",
			DetailedMessage: desc,
		}

		/*v := &errdetails.BadRequest_FieldViolation{
			Field:       "username",
			Description: desc,
		}
		customProtoError := &errdetails.BadRequest{}
		customProtoError.FieldViolations = append(customProtoError.FieldViolations, v)
		*/
		st, err := st.WithDetails(customProtoError)

		if err != nil {
			panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
		}
		return st.Err()
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
