package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceAttrInfo struct {
	Id            int64     `json:"id"`
	Device_sn     string    `json:"device_sn"`
    Attr_code     int32     `json:"attr_code"`
	Attr_name     string    `json:"attr_name"`
	Attr_permission     string    `json:"attr_permission"`
	Attr_value_ctrl    string     `json:"attr_value_ctrl"`
	Is_control    int32     `json:"is_control"`
	Attr_value_cur    string     `json:"attr_value_cur"`
	Gateway_sn    string    `json:"gateway_sn"`
}

func NewDeviceAttrInfo(f *DeviceAttrInfoPostForm, t time.Time) *DeviceAttrInfo {
	deviceattrinfo := DeviceAttrInfo{
		Device_sn:     f.Device_sn,
		Attr_code:     f.Attr_code,
		Attr_name:     f.Attr_name,
		Attr_permission:     f.Attr_permission,
		Attr_value_ctrl:    f.Attr_value_ctrl,
		Is_control:    f.Is_control,
		Attr_value_cur:    f.Attr_value_cur,
		Gateway_sn:  f.Gateway_sn}

	return &deviceattrinfo
}

func (r *DeviceAttrInfo) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_deviceattrinfo(device_sn, attr_code, attr_name, attr_permission, attr_value_ctrl, is_control, attr_value_cur, gateway_sn) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Device_sn, r.Attr_code, r.Attr_name, r.Attr_permission, r.Attr_value_ctrl, r.Is_control, r.Attr_value_cur, r.Gateway_sn); err != nil {
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

func (r *DeviceAttrInfo) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, attr_code, attr_name, attr_permission, attr_value_ctrl, is_control, attr_value_cur, gateway_sn FROM dev_deviceattrinfo WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId           sql.NullInt64 
	var tmpDevice_sn    sql.NullString 
    var tmpAttr_code    int32 
	var tmpAttr_name    sql.NullString
	var tmpAttr_permission    sql.NullString
    var tmpAttr_value_ctrl   sql.NullString  
    var tmpIs_control   int32 
    var tmpAttr_value_cur   sql.NullString 
	var tmpGateway_sn   sql.NullString 


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpAttr_code, &tmpAttr_name, &tmpAttr_permission, &tmpAttr_value_ctrl, &tmpIs_control, &tmpAttr_value_cur,
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
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}

	r.Attr_code = tmpAttr_code
	if tmpAttr_name.Valid {
		r.Attr_name = tmpAttr_name.String
	}
	if tmpAttr_permission.Valid {
		r.Attr_permission = tmpAttr_permission.String
	}
	if tmpAttr_value_ctrl.Valid {
		r.Attr_value_ctrl = tmpAttr_value_ctrl.String
	}
	r.Is_control = tmpIs_control
	if tmpAttr_value_cur.Valid {
		r.Attr_value_cur = tmpAttr_value_cur.String
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	return 0, nil
}

func (r *DeviceAttrInfo) FindByAttrSN(device_sn string, attr_code int32) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, attr_code, attr_name, attr_permission, attr_value_ctrl, is_control, attr_value_cur, gateway_sn FROM dev_deviceattrinfo WHERE attr_code = ? and device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(attr_code, device_sn)

	var tmpId           sql.NullInt64 
	var tmpDevice_sn    sql.NullString 

    var tmpAttr_code    int32 
	var tmpAttr_name    sql.NullString
	var tmpAttr_permission    sql.NullString
    var tmpAttr_value_ctrl   sql.NullString 
    var tmpIs_control   int32 
    var tmpAttr_value_cur   sql.NullString 
	var tmpGateway_sn   sql.NullString 


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpAttr_code, &tmpAttr_name,  &tmpAttr_permission, &tmpAttr_value_ctrl, &tmpIs_control, &tmpAttr_value_cur,
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
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}

	r.Attr_code = tmpAttr_code
	if tmpAttr_name.Valid {
		r.Attr_name = tmpAttr_name.String
	}
	if tmpAttr_permission.Valid {
		r.Attr_permission = tmpAttr_permission.String
	}
	if tmpAttr_value_ctrl.Valid {
		r.Attr_value_ctrl = tmpAttr_value_ctrl.String
	}
	r.Is_control = tmpIs_control
	if tmpAttr_value_cur.Valid {
		r.Attr_value_cur = tmpAttr_value_cur.String
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	return 0, nil
}

func (r *DeviceAttrInfo) ClearPass() {
	
}

func GetAllDeviceAttrInfos(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceAttrInfo, err error) {
	sqlStr := "SELECT id, device_sn, attr_code, attr_name, attr_permission, attr_value_ctrl, is_control, attr_value_cur, gateway_sn FROM dev_deviceattrinfo"
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

	records = make([]DeviceAttrInfo, 0, limit)
	for rows.Next() {

		var tmpId           sql.NullInt64 
		var tmpDevice_sn    sql.NullString 
		var tmpAttr_code    int32 
		var tmpAttr_name    sql.NullString
	    var tmpAttr_permission    sql.NullString
		var tmpAttr_value_ctrl   sql.NullString 
		var tmpIs_control   int32 
		var tmpAttr_value_cur   sql.NullString 
		var tmpGateway_sn   sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpDevice_sn, &tmpAttr_code, &tmpAttr_name, &tmpAttr_permission, &tmpAttr_value_ctrl, &tmpIs_control, &tmpAttr_value_cur,
		&tmpGateway_sn); err != nil {
			return nil, err
		}

        r := DeviceAttrInfo{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}

		r.Attr_code = tmpAttr_code
		if tmpAttr_name.Valid {
			r.Attr_name = tmpAttr_name.String
		}
		if tmpAttr_permission.Valid {
			r.Attr_permission = tmpAttr_permission.String
		}
		if tmpAttr_value_ctrl.Valid {
			r.Attr_value_ctrl = tmpAttr_value_ctrl.String
		}
		r.Is_control = tmpIs_control
		if tmpAttr_value_cur.Valid {
			r.Attr_value_cur = tmpAttr_value_cur.String
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

func (r *DeviceAttrInfo) UpdateById(id int64, f *DeviceAttrInfoPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Attr_code != -1 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_permission) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_permission = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_permission, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_value_ctrl) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_value_ctrl = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_value_ctrl, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Is_control != -1 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET is_control = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Is_control, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_value_cur) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_value_cur = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_value_cur, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET gateway_sn = ? WHERE id = ?")
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

func (r *DeviceAttrInfo) UpdateByAttrSN(device_sn string, attr_code int32, f *DeviceAttrInfoPutForm) (code int, err error) {
	db := mymysql.Conn()


	if len(f.Attr_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_name = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_name, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_permission) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_permission = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_permission, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_value_ctrl) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_value_ctrl = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_value_ctrl, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Is_control != -1 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET is_control = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Is_control, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Attr_value_cur) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET attr_value_cur = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_value_cur, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_deviceattrinfo SET gateway_sn = ? WHERE attr_code = ? and device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, attr_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *DeviceAttrInfo) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_deviceattrinfo WHERE id = ?")
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

func (r *DeviceAttrInfo) DeleteByAttrSN(device_sn string, attr_code int32) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_deviceattrinfo WHERE attr_code = ? and device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(attr_code, device_sn)
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

func (r *DeviceAttrInfo) DeleteByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_deviceattrinfo WHERE device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(device_sn)
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
