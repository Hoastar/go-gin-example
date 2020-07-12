package models


type Auth struct {
	ID int `gorm:"primary_Key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var (
		auth Auth
		err error
	)
	err = db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if auth.ID > 0 {
		return true, nil
	}
	return false, err
}