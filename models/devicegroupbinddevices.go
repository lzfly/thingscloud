package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceGroupBindDevice struct {
	Id            int64     `json:"id"`
	Group_name    string    `json:"group_name"`
	Device_sn     string    `json:"device_sn"`
	Device_name   string    `json:"device_name"`
	Type_code     int32     `json:"type_code"`
	Attr_code     int32     `json:"attr_code"`
	Attr_value    string    `json:"attr_value"`
	Gateway_sn    string    `json:"gateway_sn"`
	Username      string    `json:"username"`
}

func NewDeviceGroupBindDevice(f *DeviceGroupBindDevicePostForm, t time.Time) *DeviceGroupBindDevice {
	devicectrl := DeviceGroupBindDevice{
	    Group_name:    f.Group_name,
		Device_sn:     f.Device_sn,
		Device_name:   f.Device_name,
		Type_code:     f.Type_code,
		Attr_code:     f.Attr_code,
		Attr_value:    f.Attr_value,
		Gateway_sn:  f.Gateway_sn,
		Username:     f.Username}

	return &devicectrl
}

func (r *DeviceGroupBindDevice) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_devicegroup_bind_device(group_name, device_sn, device_name, type_code, attr_code, attr_value, gateway_sn, username) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Group_name, r.Device_sn, r.Device_name, r.Type_code, r.Attr_code, r.Attr_value, r.Gateway_sn, r.Username); err != nil {
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

func (r *DeviceGroupBindDevice) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, group_name, device_sn, device_name, type_code, attr_code, attr_value, gateway_sn, username FROM dev_devicegroup_bind_device WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId           sql.NullInt64 
    var tmpGroup_name   sql.NullString 
	var tmpDevice_sn    sql.NullString 
	var tmpDevice_name  sql.NullString 
    var tmpType_code    int32
    var tmpAttr_code    int32 
    var tmpAttr_value   sql.NullString 
	var tmpGateway_sn   sql.NullString 
	var tmpUsername     sql.NullString 


	if err := row.Scan(&tmpId, &tmpGroup_name, &tmpDevice_sn, &tmpDevice_name, &tmpType_code, &tmpAttr_code, &tmpAttr_value, &tmpGateway_sn, &tmpUsername); err != nil {
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
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}
	if tmpDevice_name.Valid {
		r.Device_name = tmpDevice_name.String
	}
	
	r.Type_code = tmpType_code
	r.Attr_code = tmpAttr_code
	if tmpAttr_value.Valid {
		r.Attr_value = tmpAttr_value.String
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpUsername.Valid {
		r.Username = tmpUsername.String
	}
	return 0, nil
}


func GetAllDeviceGroupBindDevices(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceGroupBindDevice, err error) {
	sqlStr := "SELECT id, group_name, device_sn, device_name, type_code, attr_code, attr_value, gateway_sn, username FROM dev_devicegroup_bind_device"
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

	records = make([]DeviceGroupBindDevice, 0, limit)
	for rows.Next() {

		var tmpId           sql.NullInt64 
		var tmpGroup_name   sql.NullString 
		var tmpDevice_sn    sql.NullString 
		var tmpDevice_name  sql.NullString 
		var tmpType_code    int32
		var tmpAttr_code    int32 
		var tmpAttr_value   sql.NullString 
		var tmpGateway_sn   sql.NullString 
		var tmpUsername     sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpGroup_name, &tmpDevice_sn, &tmpDevice_name, &tmpType_code, &tmpAttr_code, &tmpAttr_value, &tmpGateway_sn, &tmpUsername); err != nil {
			return nil, err
		}

        r := DeviceGroupBindDevice{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpGroup_name.Valid {
			r.Group_name = tmpGroup_name.String
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpDevice_name.Valid {
			r.Device_name = tmpDevice_name.String
		}
		
		r.Type_code = tmpType_code
		r.Attr_code = tmpAttr_code
		if tmpAttr_value.Valid {
			r.Attr_value = tmpAttr_value.String
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
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

func (r *DeviceGroupBindDevice) UpdateById(id int64, f *DeviceGroupBindDevicePutForm) (code int, err error) {
	db := mymysql.Conn()

	
	if len(f.Group_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET group_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Group_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Device_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET device_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET type_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Attr_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_value) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET attr_value = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_value, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET gateway_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Username) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicegroup_bind_device SET username = ? WHERE id = ?")
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

func (r *DeviceGroupBindDevice) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_devicegroup_bind_device WHERE id = ?")
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

func (r *DeviceGroupBindDevice) DeleteByGroupName(username string, groupname string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_devicegroup_bind_device WHERE username = ? and group_name = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(username, groupname)
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
