package models

import (
	"thingscloud/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceReport struct {
	Id            int64     `json:"id"`
	Device_sn     string    `json:"device_sn"`
	Attr_code     int32     `json:"attr_code"`
	Attr_value    int64     `json:"attr_value"`
	Gateway_sn    string    `json:"gateway_sn"`
	ReportTime    time.Time `json:"reporttime"`
}

func NewDeviceReport(f *DeviceReportPostForm, t time.Time) *DeviceReport {
	devicereport := DeviceReport{
		Device_sn:     f.Device_sn,
		Attr_code:     f.Attr_code,
		Attr_value:    f.Attr_value,
		Gateway_sn:  f.Gateway_sn,
		ReportTime:  t}

	return &devicereport
}

func (r *DeviceReport) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_devicereport(device_sn, attr_code, attr_value, gateway_sn, reporttime) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Device_sn, r.Attr_code, r.Attr_value, r.Gateway_sn, r.ReportTime); err != nil {
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

func (r *DeviceReport) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, attr_code, attr_value, gateway_sn, reporttime FROM dev_devicereport WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpId           sql.NullInt64 
	var tmpDevice_sn    sql.NullString 
    var tmpAttr_code    int32 
    var tmpAttr_value   sql.NullInt64 
	var tmpGateway_sn   sql.NullString 
	var tmpReportTime   mysql.NullTime


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpAttr_code, &tmpAttr_value, &tmpGateway_sn, 
		&tmpReportTime); err != nil {
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
	if tmpAttr_value.Valid {
		r.Attr_value = tmpAttr_value.Int64
	}
	if tmpGateway_sn.Valid {
		r.Gateway_sn = tmpGateway_sn.String
	}
	if tmpReportTime.Valid {
		r.ReportTime = tmpReportTime.Time
	}
	
	return 0, nil
}

func (r *DeviceReport) ClearPass() {
	
}

func GetAllDeviceReports(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceReport, err error) {
	sqlStr := "SELECT id, device_sn, attr_code, attr_value, gateway_sn, reporttime FROM dev_devicereport"
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

	records = make([]DeviceReport, 0, limit)
	for rows.Next() {

		var tmpId           sql.NullInt64 
		var tmpDevice_sn    sql.NullString 
		var tmpAttr_code    int32 
		var tmpAttr_value   sql.NullInt64 
		var tmpGateway_sn   sql.NullString 
		var tmpReportTime   mysql.NullTime
		
		if err := rows.Scan(&tmpId, &tmpDevice_sn, &tmpAttr_code, &tmpAttr_value, &tmpGateway_sn,
		&tmpReportTime); err != nil {
			return nil, err
		}

        r := DeviceReport{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		r.Attr_code = tmpAttr_code
		if tmpAttr_value.Valid {
			r.Attr_value = tmpAttr_value.Int64
		}
		if tmpGateway_sn.Valid {
			r.Gateway_sn = tmpGateway_sn.String
		}
		if tmpReportTime.Valid {
			r.ReportTime = tmpReportTime.Time
		}
		
			records = append(records, r)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

	return records, nil
}

func (r *DeviceReport) UpdateById(id int64, f *DeviceReportPutForm) (code int, err error) {
	db := mymysql.Conn()

	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_devicereport SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Attr_code != 0 {
		st, err1 := db.Prepare("UPDATE dev_devicereport SET attr_code = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Attr_code, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if f.Attr_value != 0 {
		st, err1 := db.Prepare("UPDATE dev_devicereport SET attr_value = ? WHERE id = ?")
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
		st, err1 := db.Prepare("UPDATE dev_devicereport SET gateway_sn = ? WHERE id = ?")
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

func (r *DeviceReport) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_devicereport WHERE id = ?")
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
