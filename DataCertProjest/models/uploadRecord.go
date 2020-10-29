package models

import "DataCertProjest/db_mysql"

type UploadRecord struct {
	Id        int
	FileName  string
	FileSize  int64
	FileCert  string //认证号
	FileTitle string
	CertTime  int64
	Phone     string //对应的用户的phone
}

func (u UploadRecord)SaveRecord()(int64 , error)  {
	rs ,err :=db_mysql.Db.Exec("insert into upload_record(file_name, file_size, file_cert, file_title, cert_time, phone) "+
		"values(?,?,?,?,?,?) ",
		u.FileName,
		u.FileSize,
		u.FileCert, u.FileTitle,
		u.CertTime, u.Phone)
	if err !=nil {
		return -1,err
	}
	id ,err :=rs.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil

}
func QueryRecordByPhone(phone string) ([]UploadRecord, error) {
	rs, err := db_mysql.Db.Query(" select id, file_name, file_size, file_cert, file_title, cert_time, phone from upload_record where phone = ?", phone)
	if err != nil {
		return nil, err
	}
	records := make([]UploadRecord, 0)
	for rs.Next() {
		var record UploadRecord
		err := rs.Scan(&record.Id, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime, &record.Phone)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}