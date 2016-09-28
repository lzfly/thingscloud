package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceGroup struct {
	Id            int64     `json:"id"`
	Group_name    string    `json:"group_name"`
	Username      string    `json:"username"`
	
}

func NewDeviceGroup(f *DeviceGroupPostForm, t time.Time) *DeviceGroup {
	devicegroup := DeviceGroup{
		Group_name:      f.Group_name,
		Username:  f.Username}

	return &devicegroup
}

func (r *DeviceGroup) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_devicegroup(group_name, username) VALUES(?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Group_name, r.Username); err != nil {
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

func (r *DeviceGroup) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, group_name, username FROM dev_devicegroup WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId            sql.NullInt64 
	var tmpGroup_name    sql.NullString   
	var tmpUsername      sql.NullString    
 
	

	if err := row.Scan(&tmpId, &tmpGroup_name, &tmpUsername); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpGroup_name.Valid {
		r.Group_name = tmpGroup_name.String
	}
	if tmpUsername.Valid {
		r.Username = tmpUsername.String
	}

	return 0, nil
}

func GetAllDeviceGroups(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceGroup, err error) {
	sqlStr := "SELECT id, group_name, username FROM dev_devicegroup"
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

	records = make([]DeviceGroup, 0, limit)
	for rows.Next() {
		var tmpId            sql.NullInt64 
		var tmpGroup_name    sql.NullString   
		var tmpUsername      sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpGroup_name, &tmpUsername); err != nil {
			return nil, err
		}

		r := DeviceGroup{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpGroup_name.Valid {
			r.Group_name = tmpGroup_name.String
		}
		if tmpUsername.Valid {
			r.Username = tmpUsername.String
		}
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *DeviceGroup) UpdateById(id int64, f *DeviceGroupPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Group_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup SET group_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Group_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Username) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup SET username = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Username, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *DeviceGroup) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_devicegroup WHERE id = ?")
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

