package services

import(
	"auth-service/models"
	pb "auth-service/pb/auth-service"
	//"time"
)

func ValidateSignUp(req *pb.SignUpRequest) (*models.User, error) {
	// birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	// if err != nil {
	// 	return nil, err
	// }
	user := models.User{
		ID:        4,
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  req.Password,
		PhoneNum:  req.PhoneNum,
	}

	return &user, nil
}










