package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Mail         string    `json:"mail"`
	Registertime time.Time `json:"registertime"`
	State        int32     `json:"state"`
}

func NewUser(f *UserPostForm, t time.Time) *User {
	user := User{
		Username:     f.Username,
		Password:     f.Password,
		Phone:     f.Phone,
		Mail:     f.Mail,
		Registertime:  t,
		State:  f.State}

	return &user
}

func (r *User) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO usr_user(username, password, phone, mail, registertime, state) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Username, r.Password, r.Phone, r.Mail, r.Registertime, r.State); err != nil {
		if e, ok := err.(*mysql.MySQLError); ok {
			//Duplicate key
			if e.Number == 1062 {
				return ErrDupRows, err
			} else {
				return ErrDatabase, err
			}
		} else {
			return ErrDatabase, err
		}
	} else {
		//r.Id, _ = result.LastInsertId()
	}

	return 0, nil
}

func (r *User) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, password, phone, mail, registertime, state FROM usr_user WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId       sql.NullInt64
	var tmpUsername sql.NullString
	var tmpPassword sql.NullString
	var tmpPhone    sql.NullString
	var tmpMail     sql.NullString
	var tmpRegistertime mysql.NullTime
	var tmpState    int32
	if err := row.Scan(&tmpId, &tmpUsername, &tmpPassword, &tmpPhone, &tmpMail, &tmpRegistertime,
		&tmpState); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpUsername.Valid {
		r.Username = tmpUsername.String
	}
	if tmpPassword.Valid {
		r.Password = tmpPassword.String
	}
	if tmpPhone.Valid {
		r.Phone = tmpPhone.String
	}
	if tmpMail.Valid {
		r.Mail = tmpMail.String
	}
	if tmpRegistertime.Valid {
		r.Registertime = tmpRegistertime.Time
	}
	r.State = tmpState

	return 0, nil
}

func (r *User) FindByUserName(name string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, password, phone, mail, registertime, state FROM usr_user WHERE username = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(name)

	var tmpId       sql.NullInt64
	var tmpUsername sql.NullString
	var tmpPassword sql.NullString
	var tmpPhone    sql.NullString
	var tmpMail     sql.NullString
	var tmpRegistertime mysql.NullTime
	var tmpState    int32
	if err := row.Scan(&tmpId, &tmpUsername, &tmpPassword, &tmpPhone, &tmpMail, &tmpRegistertime,
		&tmpState); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpUsername.Valid {
		r.Username = tmpUsername.String
	}
	if tmpPassword.Valid {
		r.Password = tmpPassword.String
	}
	if tmpPhone.Valid {
		r.Phone = tmpPhone.String
	}
	if tmpMail.Valid {
		r.Mail = tmpMail.String
	}
	if tmpRegistertime.Valid {
		r.Registertime = tmpRegistertime.Time
	}
	r.State = tmpState

	return 0, nil
}

func (r *User) ClearPass() {
	r.Password = ""
}

func GetAllUsers(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []User, err error) {
	sqlStr := "SELECT id, username, password, phone, mail, registertime, state FROM usr_user"
	if len(queryVal) > 0 {
		sqlStr += " WHERE "
		first := true
		for k, v := range queryVal {
			if !first {
				sqlStr += " AND "
			} else {
				first = false
			}

			sqlStr += k
			sqlStr += " "
			sqlStr += queryOp[k]
			sqlStr += " '"
			sqlStr += v
			sqlStr += "'"
		}
	}
	if len(order) > 0 {
		sqlStr += " ORDER BY "
		first := true
		for k, v := range order {
			if !first {
				sqlStr += ", "
			} else {
				first = false
			}

			sqlStr += k
			sqlStr += " "
			sqlStr += v
		}
	}
	sqlStr += " LIMIT " + fmt.Sprintf("%d", limit)
	if offset > 0 {
		sqlStr += " OFFSET " + fmt.Sprintf("%d", offset)
	}
	beego.Debug("sqlStr:", sqlStr)

	db := mymysql.Conn()

	st, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records = make([]User, 0, limit)
	for rows.Next() {
		var tmpId       sql.NullInt64
		var tmpUsername sql.NullString
		var tmpPassword sql.NullString
		var tmpPhone    sql.NullString
		var tmpMail     sql.NullString
		var tmpRegistertime mysql.NullTime
		var tmpState    int32
		if err := rows.Scan(&tmpId, &tmpUsername, &tmpPassword, &tmpPhone, &tmpMail, &tmpRegistertime,
			&tmpState); err != nil {
			return nil, err
		}

		r := User{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpUsername.Valid {
			r.Username = tmpUsername.String
		}
		if tmpPassword.Valid {
			r.Password = tmpPassword.String
		}
		if tmpPhone.Valid {
			r.Phone = tmpPhone.String
		}
		if tmpMail.Valid {
			r.Mail = tmpMail.String
		}
		if tmpRegistertime.Valid {
			r.Registertime = tmpRegistertime.Time
		}
		r.State = tmpState
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *User) UpdateById(id int64, f *UserPutForm) (code int, err error) {
	db := mymysql.Conn()
	
	if len(f.Password) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET password = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Password, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Phone) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET phone = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Phone, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Mail) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET mail = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Mail, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.State != 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET state = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.State, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *User) UpdateByUserName(name string, f *UserPutForm) (code int, err error) {
	db := mymysql.Conn()
	
	if len(f.Password) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET password = ? WHERE username = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Password, name)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Phone) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET phone = ? WHERE username = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Phone, name)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Mail) > 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET mail = ? WHERE username = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Mail, name)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.State != 0 {
		st, err1 := db.Prepare("UPDATE usr_user SET state = ? WHERE username = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.State, name)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *User) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM usr_user WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(id)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	} else {
		return ErrNotFound, nil
	}
}

func (r *User) DeleteByUserName(name string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM usr_user WHERE username = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(name)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	} else {
		return ErrNotFound, nil
	}
}