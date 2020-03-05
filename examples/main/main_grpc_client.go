package main

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"google.golang.org/grpc/metadata"

	pb "bandi.com/main/pkg/data"
)

// Host - name of the grpc server endpoint
const Host = "localhost"

// Port - the port where grpc server is running on
const Port = 18091

var entry *logrus.Entry = logrus.NewEntry(logrus.StandardLogger())

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := CreateUsers(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Name: "kishore",
		},
	})
	if err != nil {
		entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred ")
		// return
	}

	resp, err := StreamUsers(ctx, &pb.GetUserRequest{
		Name: "kishore",
	})
	if err != nil {
		entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred ")
		return
	}
	entry.Infof("Got Stream Response : %v", resp)

	getResp, err := GetUsers(ctx, &pb.GetUserRequest{
		Name: "kishore",
	})
	if err != nil {
		entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred ")
		return
	}
	entry.Infof("Got Get Response : %v", getResp)
}

// Execute - Executes the passed method by injecting grpc Connection of grpc to the caller.
type Execute func(*grpc.ClientConn) (interface{}, error)

// SafeExecute - Establishes a connection to grpc vault, executes the function passed by injecting client
// and Closes the connection
func SafeExecute(executeFunction Execute) (interface{}, error) {
	conn, err := getgrpcConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return executeFunction(conn)
}

// getgrpcConnection - Establishes a grpc Connection to grpc Service
func getgrpcConnection() (*grpc.ClientConn, error) {
	entry.Infof("Starting grpc Client connection")

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	endPoint := Host + ":" + strconv.Itoa(Port)

	conn, err := grpc.Dial(endPoint, opts...)
	if err != nil {
		entry.WithError(err).Errorln("Error occurred while trying to connect to grpc Service")
		return nil, err
	}

	return conn, nil
}

// GetUsers - Get Users from grpc server
func GetUsers(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	getUserFunction := func(conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewUserServiceClient(conn)

		md := metadata.Pairs("X-Custom-orgname", "This is my Custom Header O_o Weird")
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		resp, err := client.GetUser(ctx, req)
		if err != nil {
			entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred while trying to Get Users Table")
			return nil, err
		}
		return resp, nil
	}
	resp, err := SafeExecute(getUserFunction)

	if err != nil {
		return nil, err
	}

	getUserResponse, ok := resp.(*pb.GetUserResponse)
	if !ok {
		return nil, errors.New("Could not cast response to get user")
	}
	return getUserResponse, nil
}

// CreateUsers - Create Users from grpc server
func CreateUsers(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	createUserFunction := func(conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewUserServiceClient(conn)
		resp, err := client.CreateUser(ctx, req)
		if err != nil {
			entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred while trying to Create Users Table")
			return nil, err
		}
		return resp, nil
	}
	resp, err := SafeExecute(createUserFunction)

	if err != nil {
		return nil, err
	}

	createUserResponse, ok := resp.(*pb.CreateUserResponse)
	if !ok {
		return nil, errors.New("Could not cast response to list of users")
	}
	return createUserResponse, nil
}

// StreamUsers - Gets Streams from grpc server
func StreamUsers(ctx context.Context, req *pb.GetUserRequest) ([]*pb.GetUserResponse, error) {
	getUsersFunction := func(conn *grpc.ClientConn) (interface{}, error) {
		client := pb.NewUserServiceClient(conn)
		usersStream, err := client.StreamUsers(ctx, req)
		if err != nil {
			entry.WithField(logrus.ErrorKey, err).Errorln("Error occurred while trying to Get Users Table")
			return nil, err
		}

		var userResponses = make([]*pb.GetUserResponse, 1)

		for {
			userResponse, err := usersStream.Recv()
			if err == io.EOF {
				return userResponses, nil
			}
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Received Error : "))
			}
			userResponses = append(userResponses, userResponse)
		}
	}

	resp, err := SafeExecute(getUsersFunction)

	if err != nil {
		return nil, err
	}

	usersList, ok := resp.([]*pb.GetUserResponse)
	if !ok {
		return nil, errors.New("Could not cast response to list of users")
	}
	return usersList, nil
}
