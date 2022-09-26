package services

import (
	"auth-service/models"
	pb "auth-service/pb/auth-service"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	service :=  &AuthService{
		repository: &testRepository{},
		jwt: &models.JWTWrapper{
			SecretKey: "test-secret-key",
			Issuer: "test-issuer",
			ExpirationTime: 10 * time.Minute,
		},
	}
	pb.RegisterAuthServiceServer(server, service)

	go func ()  {
		if err := server.Serve(listener); err != nil{
			log.Fatal(err)
		}
	}()

	return func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSignUp(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req1 := pb.SignUpRequest{
		FullName: "test-full-name",
		Email: "test@email.com",
		Password: "test-password",
		PhoneNum: "+79991232525",
		BirthDate: "1999-10-10",
	}

	res, err := client.SignUp(context.Background(), &req1)
	require.NoError(t, err)
	require.Empty(t, res.Error)
	require.Equal(t, 201, int(res.Status))

	req2 := pb.SignUpRequest{
		FullName: "test-full-name",
		Email: "test@email.com",
		Password: "test-password",
		PhoneNum: "+79991232525",
		BirthDate: "1999-10-10-10:10",
	}

	res, err = client.SignUp(context.Background(), &req2)
	require.Error(t, err)
	//require.Equal(t, 400, int(res.Status))
}


func TestSignIn(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := pb.SignInRequest{
		Email: "test@email.com",
		Password: "test-password",
	}

	res, err := client.SignIn(context.Background(), &req)
	require.NoError(t, err)
	require.Empty(t, res.Error)
	require.NotEmpty(t, res.Token)
	require.Equal(t, 200, int(res.Status))

	fmt.Println(res.Token)
}


func TestValidate(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZW1haWwuY29tIiwiZXhwIjoxNjYwMDU1MTMzLCJpc3MiOiJ0ZXN0LWlzc3VlciJ9.iX5cVMOBmYtvu446SXIkdVrW8HpuGWgVqdLSJ7Gvb4Bw-bk7Xf2pWD6ITlgWS6dn5gsTeRM0ZAP7OWiPb9HKKw"
	req := pb.ValidateRequest{
		Token: token,
	}

	res, err := client.Validate(context.Background(), &req)
	require.NoError(t, err)
	require.Empty(t, res.Error)
	require.Equal(t, 123, int(res.UserID))
	require.Equal(t, 200, int(res.Status))
}

type testRepository struct{}

func (tr *testRepository) AddUser(user *models.User) error {
	err := errors.New("invalid user")
	if user.Email == ""{
		return err
	}

	if user.FullName == "" {
		return err
	}

	if user.BirthDate.Format("2006-10-12") == "" {
		return err
	}

	if user.PhoneNum == "" {
		return err
	}

	return nil
}

func (tr *testRepository) GetUser(email string) (*models.User, error) {
	return &models.User{
		ID:        123,
    	FullName:  "test full name",
    	Email: "test@email.com",
    	Password: "test-password",
    	PhoneNum: "test-phone-num",
    	BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}, nil
}
