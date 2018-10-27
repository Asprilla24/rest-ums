package repository

import (
	"time"
)

type User struct {
	ID                     string    `json:"id"`
	AccessFailedCount      int       `json:"accessFailedCount"`
	EmailAddress           string    `json:"emailAddress"`
	EmailConfirmationCode  string    `json:"emailConfirmationCode"`
	IsActive               bool      `json:"isActive"`
	IsEmailConfirmed       bool      `json:"isEmailConfirmed"`
	IsPhoneNumberConfirmed bool      `json:"isPhoneNumberConfirmed"`
	LastLoginTime          time.Time `json:"lastLoginTime"`
	Password               string    `json:"-"`
	PasswordResetCode      string    `json:"passwordResetCode"`
	PhoneNumber            string    `json:"phoneNumber"`
	FirstName              string    `json:"firstName"`
	LastName               string    `json:"lastName"`
	TenantID               string    `json:"tenantId"`
	Username               string    `json:"username"`
	Gender                 string    `json:"gender"`
	LastUpdated            time.Time `json:"lastUpdated"`
}

//GetUser : Function to get user by email or username
func (u *User) GetUser(emailOrUsername string) (User, error) {
	var result User

	db := connect()
	defer db.Close()

	query, err := db.Query(
		`SELECT * FROM AppUsers WHERE EmailAddress = ? OR Username = ?`,
		emailOrUsername, emailOrUsername)
	if err != nil {
		return result, err
	}
	defer query.Close()

	for query.Next() {
		isActive, isEmailConfirmed, isPhoneNumberConfirmed := 0, 0, 0
		if err := query.Scan(
			&result.ID,
			&result.AccessFailedCount,
			&result.EmailAddress,
			&result.EmailConfirmationCode,
			&isActive,
			&isEmailConfirmed,
			&isPhoneNumberConfirmed,
			&result.LastLoginTime,
			&result.Password,
			&result.PasswordResetCode,
			&result.PhoneNumber,
			&result.FirstName,
			&result.LastName,
			&result.TenantID,
			&result.Username,
			&result.Gender,
			&result.LastUpdated,
		); err != nil {
			return result, err
		}
		result.IsActive = isActive == 1
		result.IsEmailConfirmed = isEmailConfirmed == 1
		result.IsPhoneNumberConfirmed = isPhoneNumberConfirmed == 1
	}

	return result, nil
}
