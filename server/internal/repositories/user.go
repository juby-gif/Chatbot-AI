package repositories

import (
	"context"
	// "fmt"
	"database/sql"
	"time"

	"github.com/bcinnovationlabs/Apps/Chatbot-AI/server/data/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateNewUser(ctx context.Context, userId string, firstName string, middleName string, lastName string, username string, email string, password string, checkedStatus bool, age string, gender string, dob string, address string, city string, province string, country string, zip string, phone string, weight string, height string, bmi string, bodyMassIndexValue string, bloodGroup string, underlyingHealthIssues string, otherHealthIssues string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO users (user_id, first_name,middle_name,last_name,username,email,password,checked_status,age,gender,dob,address,city,province,country,zip,phone,weight,height,bmi,body_mass_index_value,blood_group,underlying_health_issues,other_health_issues) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		userId,
		firstName,
		middleName,
		lastName,
		username,
		email,
		password,
		checkedStatus,
		age,
		gender,
		dob,
		address,
		city,
		province,
		country,
		zip,
		phone,
		weight,
		height,
		bmi,
		bodyMassIndexValue,
		bloodGroup,
		underlyingHealthIssues,
		otherHealthIssues,
		// images,
	)
	return err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := "SELECT user_id, first_name,middle_name,last_name,username,email,password,checked_status,age,gender,dob,address,city,province,country,zip,phone,weight,height,bmi,body_mass_index_value,blood_group,underlying_health_issues,other_health_issues FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&m.UserId,
		&m.FirstName,
		&m.MiddleName,
		&m.LastName,
		&m.Username,
		&m.Email,
		&m.Password,
		&m.CheckedStatus,
		&m.Age,
		&m.Gender,
		&m.Dob,
		&m.Address,
		&m.City,
		&m.Province,
		&m.Country,
		&m.Zip,
		&m.Phone,
		&m.Weight,
		&m.Height,
		&m.BMI,
		&m.BodyMassIndexValue,
		&m.BloodGroup,
		&m.UnderlyingHealthIssues,
		&m.OtherHealthIssues,
		// &m.Images,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE users SET first_name = $1,middle_name = $2,last_name = $3,username = $4,email = $5,checked_status = $6,age = $7,gender = $8,dob = $9,address = $10,city = $11,province = $12,country = $13,zip = $14,phone = $15,weight = $16,height = $17,bmi = $18,body_mass_index_value = $19,blood_group = $20,underlying_health_issues = $21,other_health_issues = $22 WHERE user_id = $23"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		&m.FirstName,
		&m.MiddleName,
		&m.LastName,
		&m.Username,
		&m.Email,
		&m.CheckedStatus,
		&m.Age,
		&m.Gender,
		&m.Dob,
		&m.Address,
		&m.City,
		&m.Province,
		&m.Country,
		&m.Zip,
		&m.Phone,
		&m.Weight,
		&m.Height,
		&m.BMI,
		&m.BodyMassIndexValue,
		&m.BloodGroup,
		&m.UnderlyingHealthIssues,
		&m.OtherHealthIssues,
		// &m.Images,
		&m.UserId,
	)
	return err
}
