package services

import (
	"auth-service/models"
	pb "auth-service/pb/auth-service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/valikhan03/tool"
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
			SecretKey:      secret_key,
			ExpirationTime: exp_time,
			Issuer:         "auth service",
		},
	}
}

type IRepository interface {
	AddUser(user *models.User) error
	GetUser(email string) (*models.User, error)
	IsParticipant(auctionid string, userid int64) bool
	IsCreator(auctionid string, userid int64) bool
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var res pb.SignUpResponse
	user, err := ValidateSignUp(req)
	if err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("Validation failed: %v", err),
		}, nil
	}

	err = user.HashPassword()
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Error = models.ERR_500.Error()
		return &res, err
	}
	err = s.repository.AddUser(user)
	if err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusInternalServerError,
			Error:  "Internal Server Error. Try again.",
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
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, err
	}

	user.CheckPassword(req.Password)
	if err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Error = models.INVALID_USER_CREDS_ERR.Error()
		return &res, err
	}

	token, err := s.jwt.GenerateToken(*user)
	if err != nil {
		log.Println(err)
		res.Status = http.StatusInternalServerError
		res.Error = models.ERR_500.Error()
		return &res, err
	}

	res.Status = http.StatusOK
	res.Token = token
	res.User_ID = int32(user.ID)

	return &res, nil
}

func (s *AuthService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var res pb.ValidateResponse
	claims, err := s.jwt.ValidateToken(req.Token)
	if err != nil {
		log.Printf("ValidateToken: %v", err)
		res.Status = http.StatusUnauthorized
		res.Error = err.Error()
		return &res, err
	}

	user, err := s.repository.GetUser(claims.Email)
	if err != nil {
		log.Printf("GetUser: %v", err)
		res.Status = http.StatusInternalServerError
		res.Error = models.ERR_500.Error()
		return &res, err
	}

	switch req.Service {
	case tool.AUC_SRVC:
		if !s.repository.IsCreator(req.GetAUC_SRVC_DATA().AuctionId, user.ID) {
			res.Status = http.StatusUnauthorized
			res.Error = "Not enough rights"
			return &res, err
		}
	case tool.RUN_SRVC:
		if !s.repository.IsParticipant(req.GetRUN_SRVC_DATA().AuctionId, user.ID) {
			res.Status = http.StatusUnauthorized
			res.Error = "Not enough rights"
			return &res, err
		}
	case tool.MNG_SRVC:
		if !s.repository.IsCreator(req.GetAUC_SRVC_DATA().AuctionId, user.ID) {
			res.Status = http.StatusUnauthorized
			res.Error = "Not enough rights"
			return &res, err
		}
	}

	res.Status = http.StatusOK
	res.UserID = user.ID
	res.Email = user.Email
	res.FullName = user.FullName

	return &res, nil
}
