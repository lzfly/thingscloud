package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type UserBindDevice struct {
	Id            int64     `json:"id"`
	Username      string     `json:"username"`
	Device_sn     string    `json:"device_sn"`
	Gateway_sn    string    `json:"gateway_sn"`
}

func NewUserBindDevice(f *UserBindDevicePostForm, t time.Time) *UserBindDevice {
	userbinddevice := UserBindDevice{
		Username:       f.Username,
		Device_sn:    f.Device_sn,
		Gateway_sn:  f.Gateway_sn}

	return &userbinddevice
}

func (r *UserBindDevice) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_user_bind_device(username, device_sn, gateway_sn) VALUES(?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Username, r.Device_sn, r.Gateway_sn); err != nil {
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

func (r *UserBindDevice) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, device_sn, gateway_sn FROM dev_user_bind_device WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId            sql.NullInt64 
	var tmpUsername      sql.NullString 	
	var tmpDevice_sn     sql.NullString 
	var tmpGateway_sn    sql.NullString 


	if err := row.Scan(&tmpId, &tmpUsername, &tmpDevice_sn, 
		&tmpGateway_sn); err != nil {
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
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}

	return 0, nil
}

func (r *UserBindDevice) FindByUsernameDeviceSN(username string, device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, device_sn, gateway_sn FROM dev_user_bind_device WHERE username = ? and device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(username, device_sn)

	var tmpId            sql.NullInt64 
	var tmpUsername      sql.NullString 	
	var tmpDevice_sn     sql.NullString 
	var tmpGateway_sn    sql.NullString 


	if err := row.Scan(&tmpId, &tmpUsername, &tmpDevice_sn, 
		&tmpGateway_sn); err != nil {
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
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}

	return 0, nil
}

func (r *UserBindDevice) ClearPass() {
	
}

func GetAllUserBindDevices(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []UserBindDevice, err error) {
	sqlStr := "SELECT id, username, device_sn, gateway_sn FROM dev_user_bind_device"
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

	records = make([]UserBindDevice, 0, limit)
	for rows.Next() {

		var tmpId            sql.NullInt64 
		var tmpUsername      sql.NullString  	
		var tmpDevice_sn     sql.NullString 
		var tmpGateway_sn    sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpUsername, &tmpDevice_sn, 
		&tmpGateway_sn); err != nil {
			return nil, err
		}

        r := UserBindDevice{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpUsername.Valid {
			r.Username = tmpUsername.String
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
		
			records = append(records, r)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

	return records, nil
}

func (r *UserBindDevice) UpdateById(id int64, f *UserBindDevicePutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Username) > 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_device SET username = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Username, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_device SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_device SET gateway_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *UserBindDevice) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_user_bind_device WHERE id = ?")
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
