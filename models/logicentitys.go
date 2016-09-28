package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type LogicEntity struct {
	Id          int64    `json:"id"`
	Username    string   `json:"username"`
	Gateway_sn  string   `json:"gateway_sn"`
	If_device_sn   string   `json:"if_device_sn"`
	If_device_name   string   `json:"if_device_name"`
	If_type_code   int32    `json:"if_type_code"`
	If_attr_code   int32    `json:"if_attr_code"`
	If_operate_code      string   `json:"operate_code"`
	If_attr_value   string    `json:"if_attr_value"`
	If_attr_value2   string    `json:"if_attr_value2"`
	Th_device_sn   string   `json:"th_device_sn"`
	Th_device_name   string   `json:"th_device_name"`
	Th_type_code   int32    `json:"th_type_code"`
	Th_attr_code   int32    `json:"th_attr_code"`
	Th_attr_value   string    `json:"th_attr_value"`
}

func NewLogicEntity(f *LogicEntityPostForm, t time.Time) *LogicEntity {
	logicentity := LogicEntity{
		Username:      f.Username,
		Gateway_sn:      f.Gateway_sn,
		If_device_sn:      f.If_device_sn,
		If_device_name:      f.If_device_name,
		If_type_code:   f.If_type_code,
		If_attr_code:   f.If_attr_code,
		If_operate_code:     f.If_operate_code,
		If_attr_value:   f.If_attr_value,
		If_attr_value2:   f.If_attr_value2,
		Th_device_sn:      f.Th_device_sn,
		Th_device_name:      f.Th_device_name,
		Th_type_code:   f.Th_type_code,
		Th_attr_code:   f.Th_attr_code,
		Th_attr_value:   f.Th_attr_value}

	return &logicentity
}

func (r *LogicEntity) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_logicentity(username, gateway_sn, if_device_sn, if_device_name, if_type_code, if_attr_code, if_operate_code, if_attr_value, if_attr_value2, th_device_sn, th_device_name, th_type_code, th_attr_code, th_attr_value) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Username, r.Gateway_sn, r.If_device_sn, r.If_device_name, r.If_type_code, r.If_attr_code, r.If_operate_code, r.If_attr_value,  r.If_attr_value2, r.Th_device_sn, r.Th_device_name, r.Th_type_code, r.Th_attr_code, r.Th_attr_value); err != nil {
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

func (r *LogicEntity) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, username, gateway_sn, if_device_sn, if_device_name, if_type_code, if_attr_code, if_operate_code, if_attr_value, if_attr_value2, th_device_sn, th_device_name, th_type_code, th_attr_code, th_attr_value FROM dev_logicentity WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId           sql.NullInt64 
	var tmpUsername     sql.NullString   
	var tmpGateway_sn   sql.NullString
    var tmpIf_device_sn    sql.NullString   
    var tmpIf_device_name    sql.NullString 	
    var tmpIf_type_code    int32 
    var tmpIf_attr_code    int32  
    var tmpIf_operate_code    sql.NullString
    var tmpIf_attr_value    sql.NullString	
    var tmpIf_attr_value2    sql.NullString	
    var tmpTh_device_sn    sql.NullString   
    var tmpTh_device_name    sql.NullString 	
    var tmpTh_type_code    int32 
    var tmpTh_attr_code    int32  
    var tmpTh_attr_value    sql.NullString		 
	

	if err := row.Scan(&tmpId, &tmpUsername, &tmpGateway_sn, &tmpIf_device_sn, &tmpIf_device_name, &tmpIf_type_code, &tmpIf_attr_code, &tmpIf_operate_code, &tmpIf_attr_value, &tmpIf_attr_value2, &tmpTh_device_sn, &tmpTh_device_name, &tmpTh_type_code, &tmpTh_attr_code, &tmpTh_attr_value); err != nil {
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
	if tmpIf_device_sn.Valid {
		r.If_device_sn = tmpIf_device_sn.String
	}
	if tmpIf_device_name.Valid {
		r.If_device_name = tmpIf_device_name.String
	}
	
	r.If_type_code = tmpIf_type_code
    r.If_attr_code = tmpIf_attr_code
	
	if tmpIf_operate_code.Valid {
		r.If_operate_code = tmpIf_operate_code.String
	}
	if tmpIf_attr_value.Valid {
		r.If_attr_value = tmpIf_attr_value.String
	}
	
	if tmpIf_attr_value2.Valid {
		r.If_attr_value2 = tmpIf_attr_value2.String
	}
	
	if tmpTh_device_sn.Valid {
		r.Th_device_sn = tmpTh_device_sn.String
	}
	if tmpTh_device_name.Valid {
		r.Th_device_name = tmpTh_device_name.String
	}
	
	r.Th_type_code = tmpTh_type_code
    r.Th_attr_code = tmpTh_attr_code
	
	if tmpTh_attr_value.Valid {
		r.Th_attr_value = tmpTh_attr_value.String
	}

	return 0, nil
}


func GetAllLogicEntitys(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []LogicEntity, err error) {
	sqlStr := "SELECT id, username, gateway_sn, if_device_sn, if_device_name, if_type_code, if_attr_code, if_operate_code, if_attr_value, if_attr_value2, th_device_sn, th_device_name, th_type_code, th_attr_code, th_attr_value FROM dev_logicentity"
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

	records = make([]LogicEntity, 0, limit)
	for rows.Next() {
	
		var tmpId           sql.NullInt64 
		var tmpUsername     sql.NullString   
		var tmpGateway_sn   sql.NullString
		var tmpIf_device_sn    sql.NullString   
		var tmpIf_device_name    sql.NullString 	
		var tmpIf_type_code    int32 
		var tmpIf_attr_code    int32  
		var tmpIf_operate_code    sql.NullString
		var tmpIf_attr_value    sql.NullString	
		var tmpIf_attr_value2    sql.NullString	
		var tmpTh_device_sn    sql.NullString   
		var tmpTh_device_name    sql.NullString 	
		var tmpTh_type_code    int32 
		var tmpTh_attr_code    int32  
		var tmpTh_attr_value    sql.NullString	
		
		if err := rows.Scan(&tmpId, &tmpUsername, &tmpGateway_sn, &tmpIf_device_sn, &tmpIf_device_name, &tmpIf_type_code, &tmpIf_attr_code, &tmpIf_operate_code, &tmpIf_attr_value, &tmpIf_attr_value2, &tmpTh_device_sn, &tmpTh_device_name, &tmpTh_type_code, &tmpTh_attr_code, &tmpTh_attr_value); err != nil {
			return nil, err
		}

		r := LogicEntity{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpUsername.Valid {
			r.Username = tmpUsername.String
		}

		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
		if tmpIf_device_sn.Valid {
			r.If_device_sn = tmpIf_device_sn.String
		}
		if tmpIf_device_name.Valid {
			r.If_device_name = tmpIf_device_name.String
		}
		
		r.If_type_code = tmpIf_type_code
		r.If_attr_code = tmpIf_attr_code
		
		if tmpIf_operate_code.Valid {
			r.If_operate_code = tmpIf_operate_code.String
		}
		if tmpIf_attr_value.Valid {
			r.If_attr_value = tmpIf_attr_value.String
		}
		
		if tmpIf_attr_value2.Valid {
			r.If_attr_value2 = tmpIf_attr_value2.String
		}
		
		if tmpTh_device_sn.Valid {
			r.Th_device_sn = tmpTh_device_sn.String
		}
		if tmpTh_device_name.Valid {
			r.Th_device_name = tmpTh_device_name.String
		}
		
		r.Th_type_code = tmpTh_type_code
		r.Th_attr_code = tmpTh_attr_code
		
		if tmpTh_attr_value.Valid {
			r.Th_attr_value = tmpTh_attr_value.String
		}
		
		
		records = append(records, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *LogicEntity) UpdateById(id int64, f *LogicEntityPutForm) (code int, err error) {
	db := mymysql.Conn()
	
	if len(f.Username) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET username = ? WHERE id = ?")
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
		st, err1 := db.Prepare("UPDATE dev_logicentity SET gateway_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gateway_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.If_device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.If_device_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_device_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_device_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.If_type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_type_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_type_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.If_attr_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.If_operate_code) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_operate_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_operate_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.If_attr_value) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_attr_value = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_attr_value, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.If_attr_value2) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET if_attr_value2 = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.If_attr_value2, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Th_device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET th_device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Th_device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Th_device_name) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET th_device_name = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Th_device_name, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Th_type_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET Th_type_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Th_type_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if f.Th_attr_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET Th_attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Th_attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	
	if len(f.Th_attr_value) > 0 {
		st, err1 := db.Prepare("UPDATE dev_logicentity SET Th_attr_value = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Th_attr_value, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil
}

func (r *LogicEntity) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_logicentity WHERE id = ?")
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

