package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Device struct {
	Id           int64    `json:"id"`
	Device_sn    string   `json:"device_sn"`
	Type_code    int32    `json:"type_code"`
	Type_name    string   `json:"type_name"`
	Device_model string   `json:"device_model"`
	Device_ver   string   `json:"device_ver"`
	Protocol     string   `json:"protocol"`
	Device_name     string   `json:"device_name"`
	Gateway_sn   string   `json:"gateway_sn"`
	Is_online    int32    `json:"is_online"`
	Activetime  time.Time `json:"activetime"`
}
type Device2 struct {
	Id           int64    `json:"id,omitempty"`
	Device_sn    string   `json:"device_sn"`
	Type_code    int32    `json:"type_code"`
	Type_name    string   `json:"type_name"`
	Device_model string   `json:"device_model"`
	Device_ver   string   `json:"device_ver"`
	Protocol     string   `json:"protocol"`
	Device_name  string   `json:"device_name"`
	Gateway_sn   string   `json:"gateway_sn"`
	Is_online    int32    `json:"is_online"`
	Activetime  time.Time `json:"activetime"`
	DeviceAttrInfoInfo []DeviceAttrInfo `json:"deviceattrinfos"`
}


func NewDevice(f *DevicePostForm, t time.Time) *Device {
	device := Device{
		Device_sn:      f.Device_sn,
		Type_code:      f.Type_code,
		Type_name:      f.Type_name,
		Device_model:   f.Device_model,
		Device_ver:     f.Device_ver,
		Protocol:       f.Protocol,
		Device_name:    f.Device_name,
		Gateway_sn:     f.Gateway_sn,
		Activetime:  t}

	return &device
}

func (r *Device) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_device(device_sn, type_code, type_name, dev_model, dev_ver, protocol, dev_name, gateway_sn, is_online, activetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Device_sn, r.Type_code, r.Type_name, r.Device_model, r.Device_ver, r.Protocol, r.Device_name, r.Gateway_sn, r.Is_online, r.Activetime); err != nil {
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

func (r *Device) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, type_code, type_name, dev_model, dev_ver, protocol, dev_name, gateway_sn, is_online, activetime FROM dev_device WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId    sql.NullInt64 
	var tmpDevice_sn    sql.NullString   
	var tmpType_code    int32
    var tmpType_name    sql.NullString      
    var tmpDevice_model sql.NullString   
	var tmpDevice_ver   sql.NullString   
	var tmpProtocol     sql.NullString   
	var tmpDevice_name  sql.NullString   
	var tmpGateway_sn   sql.NullString   
	var tmpIs_online    int32    
	var tmpActivetime   mysql.NullTime 
	

	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpType_code, &tmpType_name, &tmpDevice_model, &tmpDevice_ver, &tmpProtocol, &tmpDevice_name, &tmpGateway_sn, &tmpIs_online, 
		&tmpActivetime); err != nil {
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
    r.Type_code = tmpType_code
	if tmpType_name.Valid {
		r.Type_name = tmpType_name.String
	}
	if tmpDevice_model.Valid {
		r.Device_model = tmpDevice_model.String
	}
	if tmpDevice_ver.Valid {
		r.Device_ver = tmpDevice_ver.String
	}
	if tmpProtocol.Valid {
		r.Protocol = tmpProtocol.String
	}
	if tmpDevice_name.Valid {
		r.Device_name = tmpDevice_name.String
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
    r.Is_online = tmpIs_online
	if tmpActivetime.Valid {
		r.Activetime = tmpActivetime.Time
	}

	return 0, nil
}

func (r *Device) FindByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, type_code, type_name, dev_model, dev_ver, protocol, dev_name, gateway_sn, is_online, activetime FROM dev_device WHERE device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(device_sn)

	var tmpId    sql.NullInt64 
	var tmpDevice_sn    sql.NullString   
	var tmpType_code    int32
    var tmpType_name    sql.NullString      
    var tmpDevice_model sql.NullString   
	var tmpDevice_ver   sql.NullString   
	var tmpProtocol     sql.NullString   
	var tmpDevice_name  sql.NullString   
	var tmpGateway_sn   sql.NullString   
	var tmpIs_online    int32    
	var tmpActivetime   mysql.NullTime 
	
	

	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpType_code, &tmpType_name, &tmpDevice_model, &tmpDevice_ver, &tmpProtocol, &tmpDevice_name, &tmpGateway_sn, &tmpIs_online, 
		&tmpActivetime); err != nil {
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
    r.Type_code = tmpType_code
	if tmpType_name.Valid {
		r.Type_name = tmpType_name.String
	}
	if tmpDevice_model.Valid {
		r.Device_model = tmpDevice_model.String
	}
	if tmpDevice_ver.Valid {
		r.Device_ver = tmpDevice_ver.String
	}
	if tmpProtocol.Valid {
		r.Protocol = tmpProtocol.String
	}
	if tmpDevice_name.Valid {
		r.Device_name = tmpDevice_name.String
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
    r.Is_online = tmpIs_online
	if tmpActivetime.Valid {
		r.Activetime = tmpActivetime.Time
	}

	return 0, nil
}

func (r *Device) ClearPass() {
	
}

func (r *Device2) ClearPass() {
	
}

func GetAllDevices(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []Device2, err error) {
	sqlStr := "SELECT id, device_sn, type_code, type_name, dev_model, dev_ver, protocol, dev_name, gateway_sn, is_online, activetime FROM dev_device"
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

	records = make([]Device2, 0, limit)
	for rows.Next() {
		var tmpId    sql.NullInt64 
		var tmpDevice_sn    sql.NullString   
		var tmpType_code    int32   
		var tmpType_name    sql.NullString 
		var tmpDevice_model sql.NullString   
		var tmpDevice_ver   sql.NullString   
		var tmpProtocol     sql.NullString   
		var tmpDevice_name         sql.NullString   
		var tmpGateway_sn   sql.NullString   
		var tmpIs_online    int32    
		var tmpActivetime   mysql.NullTime 
		
		if err := rows.Scan(&tmpId, &tmpDevice_sn, &tmpType_code, &tmpType_name, &tmpDevice_model, &tmpDevice_ver, &tmpProtocol, &tmpDevice_name, &tmpGateway_sn, &tmpIs_online, 
			&tmpActivetime); err != nil {
			return nil, err
		}

		r := Device2{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
        r.Type_code = tmpType_code
		if tmpType_name.Valid {
			r.Type_name = tmpType_name.String
		}
		if tmpDevice_model.Valid {
			r.Device_model = tmpDevice_model.String
		}
		if tmpDevice_ver.Valid {
			r.Device_ver = tmpDevice_ver.String
		}
		if tmpProtocol.Valid {
			r.Protocol = tmpProtocol.String
		}
		if tmpDevice_name.Valid {
			r.Device_name = tmpDevice_name.String
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
        r.Is_online = tmpIs_online
		if tmpActivetime.Valid {
			r.Activetime = tmpActivetime.Time
		}
		

	   var queryVal map[string]string = make(map[string]string)
	   var queryOp map[string]string = make(map[string]string)
	   var queryOd map[string]string = make(map[string]string)
	   queryVal["device_sn"] = r.Device_sn
	   queryOp["device_sn"] = "="
		r.DeviceAttrInfoInfo, err = GetAllDeviceAttrInfos(queryVal, queryOp, queryOd,
		20, 0)
		
	    if err != nil {
		    beego.Error("GetAllDeviceAttrInfo:", err)
	    }

		
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *Device) UpdateById(id int64, f *DevicePutForm) (code int, err error) {
	db := mymysql.Conn()

	if f.Type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET type_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Type_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET type_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Device_model) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_model = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_model, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_ver = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_ver, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Protocol) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET protocol = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Protocol, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET gateway_sn = ? WHERE id = ?")
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

func (r *Device) UpdateByDeviceSN(device_sn string, f *DevicePutForm) (code int, err error) {
	db := mymysql.Conn()

	if f.Type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET type_code = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_code, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Type_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET type_name = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Type_name, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Device_model) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_model = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_model, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_ver = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_ver, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Protocol) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET protocol = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Protocol, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET dev_name = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_name, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_device SET gateway_sn = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *Device) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_device WHERE id = ?")
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

func (r *Device) DeleteByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_device WHERE device_sn = ?")
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
