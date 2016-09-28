package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceAttr struct {
	Id            int64     `json:"id"`
	Attr_name     string    `json:"attr_name"`
	Attr_code     int32     `json:"attr_code"`
	Datatype      string    `json:"datatype"`

}

func NewDeviceAttr(f *DeviceAttrPostForm, t time.Time) *DeviceAttr {
	deviceattr := DeviceAttr{
		Attr_name:       f.Attr_name,
		Attr_code:      f.Attr_code,
		Datatype:       f.Datatype}

	return &deviceattr
}

func (r *DeviceAttr) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_deviceattr(attr_name, attr_code, datatype) VALUES(?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Attr_name, r.Attr_code, r.Datatype); err != nil {
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

func (r *DeviceAttr) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, attr_name,r attr_code, datatype FROM dev_deviceattr WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId                 sql.NullInt64 
	var tmpAttr_name          sql.NullString   
	var tmpAttr_code          int32 	
	var tmpDatatype           sql.NullString 


	if err := row.Scan(&tmpId, &tmpAttr_name, &tmpAttr_code, 
		&tmpDatatype); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpAttr_name.Valid {
		r.Attr_name = tmpAttr_name.String
	}
	r.Attr_code = tmpAttr_code

	if tmpDatatype.Valid {
		r.Datatype = tmpDatatype.String
	}

	return 0, nil
}

func (r *DeviceAttr) FindByAttrCode(attrCode int32) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, attr_name, attr_code, datatype FROM dev_deviceattr WHERE attr_code = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(attrCode)

	var tmpId                 sql.NullInt64 
	var tmpAttr_name          sql.NullString   
	var tmpAttr_code          int32 	
	var tmpDatatype           sql.NullString 


	if err := row.Scan(&tmpId, &tmpAttr_name, &tmpAttr_code, 
		&tmpDatatype); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpAttr_name.Valid {
		r.Attr_name = tmpAttr_name.String
	}
	r.Attr_code = tmpAttr_code

	if tmpDatatype.Valid {
		r.Datatype = tmpDatatype.String
	}

	return 0, nil
}

func (r *DeviceAttr) ClearPass() {
	
}

func GetAllDeviceAttrs(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceAttr, err error) {
	sqlStr := "SELECT id, attr_name, attr_code, datatype FROM dev_deviceattr"
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

	records = make([]DeviceAttr, 0, limit)
	for rows.Next() {
		var tmpId            sql.NullInt64 
		var tmpAttr_name          sql.NullString   
		var tmpAttr_code     int32 	
		var tmpDatatype          sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpAttr_name, &tmpAttr_code,
		&tmpDatatype); err != nil {
			return nil, err
		}

        r := DeviceAttr{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpAttr_name.Valid {
			r.Attr_name = tmpAttr_name.String
		}
		r.Attr_code = tmpAttr_code

		if tmpDatatype.Valid {
			r.Datatype = tmpDatatype.String
		}

	
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *DeviceAttr) UpdateById(id int64, f *DeviceAttrPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Attr_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattr SET attr_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Attr_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattr SET attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Datatype) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattr SET datatype = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Datatype, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *DeviceAttr) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_deviceattr WHERE id = ?")
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
