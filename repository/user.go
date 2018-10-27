package repository

type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
}

func (u *User) GetUser(email string, password string) User {
	var result User

	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT uid, username, email, role_id FROM Users WHERE email = ? AND password = ?", email, password)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&result.UID, &result.Username, &result.Email, &result.RoleID); err != nil {
			panic(err.Error())
		}
	}

	return result
}
