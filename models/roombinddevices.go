package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type RoomBindDevice struct {
	Id            int64     `json:"id"`
	Room_name    string    `json:"room_name"`
	Device_sn     string    `json:"device_sn"`
	Device_name   string    `json:"device_name"`
	Type_code     int32     `json:"type_code"`
	Gateway_sn    string    `json:"gateway_sn"`
	Username      string    `json:"username"`
}

func NewRoomBindDevice(f *RoomBindDevicePostForm, t time.Time) *RoomBindDevice {
	devicectrl := RoomBindDevice{
	    Room_name:    f.Room_name,
		Device_sn:     f.Device_sn,
		Device_name:   f.Device_name,
		Type_code:     f.Type_code,
		Gateway_sn:  f.Gateway_sn,
		Username:     f.Username}

	return &devicectrl
}

func (r *RoomBindDevice) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_room_bind_device(room_name, device_sn, device_name, type_code, gateway_sn, username) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Room_name, r.Device_sn, r.Device_name, r.Type_code, r.Gateway_sn, r.Username); err != nil {
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

func (r *RoomBindDevice) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, room_name, device_sn, device_name, type_code, gateway_sn, username FROM dev_room_bind_device WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId           sql.NullInt64 
    var tmpRoom_name   sql.NullString 
	var tmpDevice_sn    sql.NullString 
	var tmpDevice_name  sql.NullString 
    var tmpType_code    int32
	var tmpGateway_sn   sql.NullString 
	var tmpUsername     sql.NullString 


	if err := row.Scan(&tmpId, &tmpRoom_name, &tmpDevice_sn, &tmpDevice_name, &tmpType_code, &tmpGateway_sn, &tmpUsername); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
	if tmpRoom_name.Valid {
		r.Room_name = tmpRoom_name.String
	}
	if tmpDevice_sn.Valid {
		r.Device_sn = tmpDevice_sn.String
	}
	if tmpDevice_name.Valid {
		r.Device_name = tmpDevice_name.String
	}
	r.Type_code = tmpType_code
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpUsername.Valid {
		r.Username = tmpUsername.String
	}
	return 0, nil
}


func GetAllRoomBindDevices(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []RoomBindDevice, err error) {
	sqlStr := "SELECT id, room_name, device_sn, device_name, type_code, gateway_sn, username FROM dev_room_bind_device"
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

	records = make([]RoomBindDevice, 0, limit)
	for rows.Next() {

		var tmpId           sql.NullInt64 
		var tmpRoom_name   sql.NullString 
		var tmpDevice_sn    sql.NullString 
		var tmpDevice_name  sql.NullString 
		var tmpType_code    int32
		var tmpGateway_sn   sql.NullString 
		var tmpUsername     sql.NullString 
		
		if err := rows.Scan(&tmpId, &tmpRoom_name, &tmpDevice_sn, &tmpDevice_name, &tmpType_code, &tmpGateway_sn, &tmpUsername); err != nil {
			return nil, err
		}

        r := RoomBindDevice{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpRoom_name.Valid {
			r.Room_name = tmpRoom_name.String
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpDevice_name.Valid {
			r.Device_name = tmpDevice_name.String
		}
		r.Type_code = tmpType_code
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

func (r *RoomBindDevice) UpdateById(id int64, f *RoomBindDevicePutForm) (code int, err error) {
	db := mymysql.Conn()

	
	if len(f.Room_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_room_bind_device SET room_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Room_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_room_bind_device SET device_sn = ? WHERE id = ?")
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
	if len(f.Gateway_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_room_bind_device SET gateway_sn = ? WHERE id = ?")
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
		st, err1 := db.Prepare("UPDATE dev_room_bind_device SET username = ? WHERE id = ?")
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

func (r *RoomBindDevice) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_room_bind_device WHERE id = ?")
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
