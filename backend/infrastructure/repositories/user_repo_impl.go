package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	entities "social-network/core/entities"
	repositories "social-network/core/repositories"

	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *sql.DB
}

// GetName implements repositories.UserRepo.
func (u *userRepo) GetName(ctx context.Context, id string) (string, error) {
	row := u.db.QueryRowContext(ctx, `SELECT username FROM users WHERE id = ?`, id)
	var Username string
	err := row.Scan(&Username)
	if err != nil {
		return "", err
	}
	fmt.Println("errr",err)
	return Username, nil
}

func NewUserRepo(db *sql.DB) repositories.UserRepo {
	return &userRepo{db: db}
}

// ____________________________________________________________________________________
func (r *userRepo) Connect(ctx context.Context, input, password string) (int, error) {
	// select the correct statment
	var stmt, mailOrNick string
	if strings.Contains(input, "@") {
		mailOrNick = "mail"
		stmt = `SELECT id FROM users WHERE email = ? AND password = ?`
	} else {
		mailOrNick = "username"
		stmt = `SELECT id FROM users WHERE username = ? AND password = ?`
	}
	var id int
	// this statment work also stmt = `SELECT id FROM users WHERE`+mailOrNick+`= ? AND password = ?` (update the if stmt)
	err := r.db.QueryRowContext(ctx, stmt, input, password).Scan(&id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New(mailOrNick + " or password invalid")
	} else if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *userRepo) Save(ctx context.Context, u *entities.User) error {
	query := `
      INSERT INTO users 
        (username,email, password, first_name, last_name, date_of_birth, city, profile_pic, about_me)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.ExecContext(ctx, query,
		u.Username, u.Email, u.Password, u.FirstName, u.LastName,
		u.DateOfBirth, u.City, u.ProfilePic, u.AboutMe,
	)

	return err
}

func (r *userRepo) Get(ctx context.Context, id string) entities.User {
	query := `
		SELECT id ,username, email, first_name, last_name, date_of_birth, 
			   city ,profile_pic, about_me, createdAt
		FROM users 
		WHERE id = ?
	`
	var user entities.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.City, &user.ProfilePic, &user.AboutMe, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("utilisateur avec l'ID %s introuvable", id)
			return entities.User{}
		}
		return entities.User{}
	}

	return user
}

// better add hash compare to delivery or app layer
func (r *userRepo) Hash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func (r *userRepo) Compare(hashed string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
