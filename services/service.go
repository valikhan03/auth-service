package services

import (
	"auth-service/models"
	"auth-service/pb"
	"context"
	"os"
	"time"
)

type Service struct {
	repository IRepository
	jwt        *models.JWTWrapper
}

func NewService(repository IRepository) *Service {
	secret_key := os.Getenv("")
	exp_time := 24 * time.Hour
	return &Service{
		repository: repository,
		jwt: &models.JWTWrapper{
			SecretKey: secret_key,
			ExpirationTime: int64(exp_time),
			Issuer: "auth service",
		},
	}
}

type IRepository interface {
	AddUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

func (s *Service) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var res pb.SignUpResponse

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		//ADD HERE ERROR HANDLER
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
		return nil, err
	}

	return &res, err
}

func (s *Service) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	var res pb.SignInResponse
	user, err := s.repository.GetUser(req.Email)
	if err != nil {
		res.Status = "error"
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, nil
	}

	user.CheckPassword(req.Password)
	if err != nil {
		res.Status = "error"
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, nil
	}

	token, err := s.jwt.GenerateToken(user.Email)
	if err != nil {
		res.Status = "error"
		res.Error = models.ERR_500.Error()
		return &res, nil
	}

	res.Status = "signed in"
	res.Token = token

	return &res, nil
}

func (s *Service) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var res pb.ValidateResponse
	claims, err := s.jwt.ValidateToken(req.Token)
	if err != nil {
		res.Status = "error"
		res.Error = err.Error()
		return &res, nil
	}

	user, err := s.repository.GetUser(claims.Email)
	if err != nil {
		res.Status = "error"
		res.Error = models.ERR_500.Error()
		return &res, nil
	}

	res.Status = "validated"
	res.UserID = user.ID

	return &res, nil
}

