package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceType struct {
	Id            int64     `json:"id"`
	Type_name     string    `json:"type_name"`
	Type_code     int32     `json:"type_code"`
}

func NewDeviceType(f *DeviceTypePostForm, t time.Time) *DeviceType {
	devicetype := DeviceType{
		Type_name:       f.Type_name,
		Type_code:  f.Type_code}

	return &devicetype
}

func (r *DeviceType) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_devicetype(type_name, type_code) VALUES(?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Type_name, r.Type_code); err != nil {
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

func (r *DeviceType) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, type_name, type_code FROM dev_devicetype WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId            sql.NullInt64 
	var tmpType_name          sql.NullString   
	var tmpType_code     int32 	

	if err := row.Scan(&tmpId, &tmpType_name,  
		&tmpType_code); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpType_name.Valid {
		r.Type_name = tmpType_name.String
	}
	r.Type_code = tmpType_code

	return 0, nil
}

func (r *DeviceType) FindByTypeCode(typeCode int32) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, type_name, type_code FROM dev_devicetype WHERE type_code = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(typeCode)

	var tmpId            sql.NullInt64 
	var tmpType_name          sql.NullString   
	var tmpType_code     int32 	

	if err := row.Scan(&tmpId, &tmpType_name,  
		&tmpType_code); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpType_name.Valid {
		r.Type_name = tmpType_name.String
	}
	r.Type_code = tmpType_code

	return 0, nil
}

func (r *DeviceType) ClearPass() {
	
}

func GetAllDeviceTypes(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceType, err error) {
	sqlStr := "SELECT id, type_name, type_code FROM dev_devicetype"
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

	records = make([]DeviceType, 0, limit)
	for rows.Next() {
		var tmpId            sql.NullInt64 
		var tmpType_name          sql.NullString   
		var tmpType_code     int32 
		
		if err := rows.Scan(&tmpId, &tmpType_name,  
		&tmpType_code); err != nil {
			return nil, err
		}

		r := DeviceType{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpType_name.Valid {
			r.Type_name = tmpType_name.String
		}
		r.Type_code = tmpType_code
	
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *DeviceType) UpdateById(id int64, f *DeviceTypePutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Type_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicetype SET type_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_devicetype SET type_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *DeviceType) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_devicetype WHERE id = ?")
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
