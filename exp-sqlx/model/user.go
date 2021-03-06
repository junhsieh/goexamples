package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	IDUser   uint   `db:"IDUser"`
	Username string `db:"Username"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return fmt.Errorf("Username must be greater than or equal to three characters.")
	}

	return nil
}

func (u *User) Load() error {
	var err error

	sq := "SELECT "
	sq += "  IDUser "
	sq += ", Username "
	sq += "FROM user "
	sq += "WHERE "
	sq += "Username = ? "

	err = db.Get(u, sq, u.Username)

	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func (u *User) Save() error {
	var err error

	_, err = db.NamedExec(`INSERT INTO user (Username) VALUES (:Username)`,
		map[string]interface{}{
			"Username": u.Username,
		})

	return err
}

func ListUser() ([]User, error) {
	var err error
	userArr := make([]User, 0)

	sq := "SELECT "
	sq += "  IDUser "
	sq += ", Username "
	sq += "FROM user "

	err = db.Select(&userArr, sq)

	if err == sql.ErrNoRows {
		return userArr, nil
	} else if err != nil {
		return nil, err
	}

	return userArr, nil
}
