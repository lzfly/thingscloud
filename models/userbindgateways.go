package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type UserBindGateway struct {
	Id            int64     `json:"id"`
	Username      string     `json:"username"`
	Gateway_sn    string    `json:"gateway_sn"`
	Is_master     int32     `json:"is_master"`
}

func NewUserBindGateway(f *UserBindGatewayPostForm, t time.Time) *UserBindGateway {
	userbindgateway := UserBindGateway{
		Username:       f.Username,
		Gateway_sn:    f.Gateway_sn,
		Is_master:  f.Is_master}

	return &userbindgateway
}

func (r *UserBindGateway) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_user_bind_gateway(username, gateway_sn, is_master) VALUES(?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Username, r.Gateway_sn, r.Is_master); err != nil {
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

func (r *UserBindGateway) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, gateway_sn, is_master FROM dev_user_bind_gateway WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId            sql.NullInt64 
	var tmpUsername      sql.NullString
	var tmpGateway_sn    sql.NullString 
	var tmpIs_master     int32


	if err := row.Scan(&tmpId, &tmpUsername, &tmpGateway_sn, 
		&tmpIs_master); err != nil {
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
    r.Is_master = tmpIs_master

	return 0, nil
}

func (r *UserBindGateway) FindByUsernameGatewaySN(name string, gateway_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, gateway_sn, is_master FROM dev_user_bind_gateway WHERE username = ? and gateway_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(name, gateway_sn)

	var tmpId            sql.NullInt64 
	var tmpUsername      sql.NullString
	var tmpGateway_sn    sql.NullString 
	var tmpIs_master     int32


	if err := row.Scan(&tmpId, &tmpUsername, &tmpGateway_sn, 
		&tmpIs_master); err != nil {
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
    r.Is_master = tmpIs_master

	return 0, nil
}

func (r *UserBindGateway) ClearPass() {
	
}

func GetAllUserBindGateways(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []UserBindGateway, err error) {
	sqlStr := "SELECT id, username, gateway_sn, is_master FROM dev_user_bind_gateway"
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

	records = make([]UserBindGateway, 0, limit)
	for rows.Next() {

		var tmpId            sql.NullInt64 
		var tmpUsername      sql.NullString 	
		var tmpGateway_sn    sql.NullString 
		var tmpIs_master     int32
		
		if err := rows.Scan(&tmpId, &tmpUsername, &tmpGateway_sn, 
		&tmpIs_master); err != nil {
			return nil, err
		}

        r := UserBindGateway{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpUsername.Valid {
			r.Username = tmpUsername.String
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
		r.Is_master = tmpIs_master
		
			records = append(records, r)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

	return records, nil
}

func (r *UserBindGateway) UpdateById(id int64, f *UserBindGatewayPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Username) > 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_gateway SET username = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Username, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_gateway SET gateway_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Is_master != 0 {
		st, err1 := db.Prepare("UPDATE dev_user_bind_gateway SET is_master = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Is_master, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *UserBindGateway) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_user_bind_gateway WHERE id = ?")
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
