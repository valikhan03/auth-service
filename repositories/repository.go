package repositories

import (
	"auth-service/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddUser(user *models.User) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM tb_users WHERE email=$1)`
	err := r.db.Get(&exists, query, user.Email)
	if err != nil {
		return err
	}

	if exists {
		return models.USER_ALREADY_EXISTS_ERR
	}

	query = `
		INSERT INTO tb_users (id, full_name, email, password, phone_num, birth_date)
			VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = r.db.Exec(query, user.ID, user.FullName, user.Email, user.Password, user.PhoneNum, user.BirthDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUser(email string) (*models.User, error) {
	query := `SELECT * FROM tb_users WHERE email=$1`
	var user models.User
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}


type username struct{
	Fullname string
}

func (r *Repository) GetUserName(id int64) (string, error) {
	var u username
	err := r.db.Get(&u, "SELECT full_name FROM tb_users WHERE id=$1", id)
	if err != nil{
		return "", err
	}

	return u.Fullname, nil
}


func (r *Repository) IsCreator(auctionid string, userid int64) bool {
	var v_userid int64
	query := `select organizer_id from tb_auctions where id=$1`
	rows, err := r.db.DB.Query(query, auctionid)
	if err != nil {
		return false
	}

	for rows.Next() {
		err = rows.Scan(v_userid)
		if err != nil{
			return false
		}
	}

	if v_userid == userid {
		return true
	}

	return false
}


func(r *Repository) IsParticipant(auctionid string, userid int64) bool {
	res := false
	query := `select 
		case when exists(select 1 from tb_participants where auction_id=$1 and user_id=$2) 
			then TRUE 
			else FALSE 
		end
		from dual`
	rows, err := r.db.DB.Query(query, auctionid, userid)
	if err != nil{
		return res
	}
	for rows.Next() {
		rows.Scan(res)
	}

	return res
}