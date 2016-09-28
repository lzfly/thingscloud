package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Gateway struct {
	Id            int64     `json:"id"`
	Gateway_sn    string    `json:"gateway_sn"`
	Gw_model      string    `json:"gw_model"`
	Gw_type       string    `json:"gw_type"`
	Gw_mac        string    `json:"gw_mac"`
	Wifi_ssid     string    `json:"wifi_ssid"`
	Wifi_pwd      string    `json:"wifi_pwd"`
	Hw_ver        string    `json:"hw_ver"`
	Sw_ver        string    `json:"sw_ver"`
	State         int32     `json:"state"`
	Operate_cap   int32     `json:"operate_cap"`
	Activetime    time.Time `json:"activetime"`
	
}

func NewGateway(f *GatewayPostForm, t time.Time) *Gateway {
	gateway := Gateway{
		Gateway_sn:      f.Gateway_sn,
		Gw_model:        f.Gw_model,
		Gw_type:         f.Gw_type,
		Gw_mac:          f.Gw_mac,
		Wifi_ssid:       f.Wifi_ssid,
		Wifi_pwd:        f.Wifi_pwd,
		Hw_ver:          f.Hw_ver,
		Sw_ver:          f.Sw_ver,
		State:           f.State,
		Operate_cap:     f.Operate_cap,
		Activetime:  t}

	return &gateway
}

func (r *Gateway) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_gateway(gateway_sn, gw_model, gw_type, gw_mac, wifi_ssid, wifi_pwd, hw_ver, sw_ver, state, operate_cap, activetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Gateway_sn, r.Gw_model, r.Gw_type, r.Gw_mac, r.Wifi_ssid, r.Wifi_pwd, r.Hw_ver, r.Sw_ver, r.State, r.Operate_cap, r.Activetime); err != nil {
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

func (r *Gateway) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, gateway_sn, gw_model, gw_type, gw_mac, wifi_ssid, wifi_pwd, hw_ver, sw_ver, state, operate_cap, activetime FROM dev_gateway WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId            sql.NullInt64 
	var tmpGateway_sn    sql.NullString   
	var tmpGw_model      sql.NullString    
    var tmpGw_type       sql.NullString   
	var tmpGw_mac        sql.NullString   
	var tmpWifi_ssid     sql.NullString   
	var tmpWifi_pwd      sql.NullString   
	var tmpHw_ver        sql.NullString   
	var tmpSw_ver        sql.NullString 
	var tmpState         int32   
	var tmpOperate_cap   int32 	
	var tmpActivetime    mysql.NullTime 
	

	if err := row.Scan(&tmpId, &tmpGateway_sn, &tmpGw_model, &tmpGw_type, &tmpGw_mac, &tmpWifi_ssid, &tmpWifi_pwd, &tmpHw_ver, &tmpSw_ver, &tmpState, &tmpOperate_cap, 
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
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpGw_model.Valid {
		r.Gw_model = tmpGw_model.String
	}
	if tmpGw_type.Valid {
		r.Gw_type = tmpGw_type.String
	}
	if tmpGw_mac.Valid {
		r.Gw_mac = tmpGw_mac.String
	}
	if tmpWifi_ssid.Valid {
		r.Wifi_ssid = tmpWifi_ssid.String
	}
	if tmpWifi_pwd.Valid {
		r.Wifi_pwd = tmpWifi_pwd.String
	}
	if tmpHw_ver.Valid {
		r.Hw_ver = tmpHw_ver.String
	}
	if tmpSw_ver.Valid {
		r.Sw_ver = tmpSw_ver.String
	}
	r.State = tmpState
	r.Operate_cap = tmpOperate_cap
	if tmpActivetime.Valid {
		r.Activetime = tmpActivetime.Time
	}

	return 0, nil
}

func (r *Gateway) FindByGatewaySN(gateway_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, gateway_sn, gw_model, gw_type, gw_mac, wifi_ssid, wifi_pwd, hw_ver, sw_ver, state, operate_cap, activetime FROM dev_gateway WHERE gateway_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(gateway_sn)

	var tmpId            sql.NullInt64 
	var tmpGateway_sn    sql.NullString   
	var tmpGw_model      sql.NullString    
    var tmpGw_type       sql.NullString   
	var tmpGw_mac        sql.NullString   
	var tmpWifi_ssid     sql.NullString   
	var tmpWifi_pwd      sql.NullString   
	var tmpHw_ver        sql.NullString   
	var tmpSw_ver        sql.NullString 
	var tmpState         int32   
	var tmpOperate_cap   int32 	
	var tmpActivetime    mysql.NullTime 
	

	if err := row.Scan(&tmpId, &tmpGateway_sn, &tmpGw_model, &tmpGw_type, &tmpGw_mac, &tmpWifi_ssid, &tmpWifi_pwd, &tmpHw_ver, &tmpSw_ver, &tmpState, &tmpOperate_cap, 
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
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpGw_model.Valid {
		r.Gw_model = tmpGw_model.String
	}
	if tmpGw_type.Valid {
		r.Gw_type = tmpGw_type.String
	}
	if tmpGw_mac.Valid {
		r.Gw_mac = tmpGw_mac.String
	}
	if tmpWifi_ssid.Valid {
		r.Wifi_ssid = tmpWifi_ssid.String
	}
	if tmpWifi_pwd.Valid {
		r.Wifi_pwd = tmpWifi_pwd.String
	}
	if tmpHw_ver.Valid {
		r.Hw_ver = tmpHw_ver.String
	}
	if tmpSw_ver.Valid {
		r.Sw_ver = tmpSw_ver.String
	}
	r.State = tmpState
	r.Operate_cap = tmpOperate_cap
	if tmpActivetime.Valid {
		r.Activetime = tmpActivetime.Time
	}

	return 0, nil
}

func (r *Gateway) ClearPass() {
	
}

func GetAllGateways(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []Gateway, err error) {
	sqlStr := "SELECT id, gateway_sn, gw_model, gw_type, gw_mac, wifi_ssid, wifi_pwd, hw_ver, sw_ver, state, operate_cap, activetime FROM dev_gateway"
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

	records = make([]Gateway, 0, limit)
	for rows.Next() {
		var tmpId            sql.NullInt64 
		var tmpGateway_sn    sql.NullString   
		var tmpGw_model      sql.NullString    
		var tmpGw_type       sql.NullString   
		var tmpGw_mac        sql.NullString   
		var tmpWifi_ssid     sql.NullString   
		var tmpWifi_pwd      sql.NullString   
		var tmpHw_ver        sql.NullString   
		var tmpSw_ver        sql.NullString 
		var tmpState         int32   
		var tmpOperate_cap   int32 	
		var tmpActivetime    mysql.NullTime 
		
		if err := rows.Scan(&tmpId, &tmpGateway_sn, &tmpGw_model, &tmpGw_type, &tmpGw_mac, &tmpWifi_ssid, &tmpWifi_pwd, &tmpHw_ver, &tmpSw_ver, &tmpState, &tmpOperate_cap,  
			&tmpActivetime); err != nil {
			return nil, err
		}

		r := Gateway{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
		if tmpGw_model.Valid {
			r.Gw_model = tmpGw_model.String
		}
		if tmpGw_type.Valid {
			r.Gw_type = tmpGw_type.String
		}
		if tmpGw_mac.Valid {
			r.Gw_mac = tmpGw_mac.String
		}
		if tmpWifi_ssid.Valid {
			r.Wifi_ssid = tmpWifi_ssid.String
		}
		if tmpWifi_pwd.Valid {
			r.Wifi_pwd = tmpWifi_pwd.String
		}
		if tmpHw_ver.Valid {
			r.Hw_ver = tmpHw_ver.String
		}
		if tmpSw_ver.Valid {
			r.Sw_ver = tmpSw_ver.String
		}
		r.State = tmpState
		r.Operate_cap = tmpOperate_cap
		if tmpActivetime.Valid {
			r.Activetime = tmpActivetime.Time
		}
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *Gateway) UpdateById(id int64, f *GatewayPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Gw_model) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_model = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_model, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gw_type) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_type = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_type, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gw_mac) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_mac = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_mac, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Wifi_ssid) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET wifi_ssid = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Wifi_ssid, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Wifi_pwd) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET wifi_pwd = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Wifi_pwd, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Hw_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET hw_ver = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Hw_ver, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Sw_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET sw_ver = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Sw_ver, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.State != 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET state = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.State, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.Operate_cap != 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET operate_cap = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Operate_cap, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *Gateway) UpdateByGatewaySN(gateway_sn string, f *GatewayPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Gw_model) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_model = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_model, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gw_type) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_type = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_type, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Gw_mac) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET gw_mac = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gw_mac, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Wifi_ssid) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET wifi_ssid = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Wifi_ssid, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Wifi_pwd) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET wifi_pwd = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Wifi_pwd, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Hw_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET hw_ver = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Hw_ver, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Sw_ver) > 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET sw_ver = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Sw_ver, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.State != 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET state = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.State, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.Operate_cap != 0 {
		st, err1 := db.Prepare("UPDATE dev_gateway SET operate_cap = ? WHERE gateway_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Operate_cap, gateway_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *Gateway) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_gateway WHERE id = ?")
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

func (r *Gateway) DeleteByGatewaySN(gateway_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_gateway WHERE gateway_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(gateway_sn)
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
