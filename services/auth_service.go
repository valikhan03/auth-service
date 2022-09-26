package services

import (
	"auth-service/models"
	pb "auth-service/pb/auth-service"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type AuthService struct {
	repository IRepository
	jwt        *models.JWTWrapper
}

func NewService(repository IRepository) *AuthService {
	secret_key := os.Getenv("SECRET-KEY")
	exp_time := 24 * time.Hour
	
	return &AuthService{
		repository: repository,
		jwt: &models.JWTWrapper{
			SecretKey: secret_key,
			ExpirationTime: exp_time,
			Issuer: "auth service",
		},
	}
}

type IRepository interface {
	AddUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusBadRequest,
			Error: fmt.Sprintf("Invalid birth date: %v", err),
		}, err
	}

	user := models.User{
		ID:        4,
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  req.Password,
		PhoneNum:  req.PhoneNum,
		BirthDate: birthDate,
	}
	err = s.repository.AddUser(&user)
	if err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusInternalServerError,
			Error: "Internal Server Error. Try again.",
		}, err
	}

	return &pb.SignUpResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	var res pb.SignInResponse
	user, err := s.repository.GetUser(req.Email)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, err
	}

	user.CheckPassword(req.Password)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, err
	}

	token, err := s.jwt.GenerateToken(*user)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Error = models.ERR_500.Error()
		return &res, err
	}

	res.Status = http.StatusOK
	res.Token = token

	return &res, nil
}

func (s *AuthService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var res pb.ValidateResponse
	claims, err := s.jwt.ValidateToken(req.Token)
	if err != nil {
		res.Status = http.StatusUnauthorized
		res.Error = err.Error()
		return &res, err
	}

	user, err := s.repository.GetUser(claims.Email)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Error = models.ERR_500.Error()
		return &res, err
	}

	res.Status = http.StatusOK
	res.UserID = user.ID

	return &res, nil
}

