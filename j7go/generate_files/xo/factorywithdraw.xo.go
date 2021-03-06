// Package xo contains the types for schema 'ddg_local'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// FactoryWithdraw represents a row from 'ddg_local.factory_withdraw'.
type FactoryWithdraw struct {
	ID           int64           `json:"id"`            // id
	Fid          sql.NullInt64   `json:"fid"`           // fid
	AccountName  sql.NullString  `json:"account_name"`  // account_name
	AccountNo    sql.NullString  `json:"account_no"`    // account_no
	BankName     sql.NullString  `json:"bank_name"`     // bank_name
	ImgURL       sql.NullString  `json:"img_url"`       // img_url
	CreatedAt    mysql.NullTime  `json:"created_at"`    // created_at
	UpdateAt     mysql.NullTime  `json:"update_at"`     // update_at
	DoStatus     sql.NullInt64   `json:"do_status"`     // do_status
	CurrentMoney sql.NullFloat64 `json:"current_money"` // current_money
	PreMoney     sql.NullFloat64 `json:"pre_money"`     // pre_money
	PostMoney    sql.NullFloat64 `json:"post_money"`    // post_money
	DoTime       mysql.NullTime  `json:"do_time"`       // do_time
	Remark       sql.NullString  `json:"remark"`        // remark

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryWithdraw exists in the database.
func (fw *FactoryWithdraw) Exists() bool { //factory_withdraw
	return fw._exists
}

// Deleted provides information if the FactoryWithdraw has been deleted from the database.
func (fw *FactoryWithdraw) Deleted() bool {
	return fw._deleted
}

// Get table name
func GetFactoryWithdrawTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "factory_withdraw", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryWithdraw to the database.
func (fw *FactoryWithdraw) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fw._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryWithdrawTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`id, fid, account_name, account_no, bank_name, img_url, created_at, update_at, do_status, current_money, pre_money, post_money, do_time, remark` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fw.ID, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, fw.ID, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark)
	} else {
		res, err = dbConn.Exec(sqlstr, fw.ID, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark)
	}

	if err != nil {
		return err
	}

	// set existence
	fw._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	fw.ID = int64(id)
	fw._exists = true

	return nil
}

// Update updates the FactoryWithdraw in the database.
func (fw *FactoryWithdraw) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fw._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryWithdrawTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fid = ?, account_name = ?, account_no = ?, bank_name = ?, img_url = ?, created_at = ?, update_at = ?, do_status = ?, current_money = ?, pre_money = ?, post_money = ?, do_time = ?, remark = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark, fw.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark, fw.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fw.Fid, fw.AccountName, fw.AccountNo, fw.BankName, fw.ImgURL, fw.CreatedAt, fw.UpdateAt, fw.DoStatus, fw.CurrentMoney, fw.PreMoney, fw.PostMoney, fw.DoTime, fw.Remark, fw.ID)
	}
	return err
}

// Save saves the FactoryWithdraw to the database.
func (fw *FactoryWithdraw) Save(ctx context.Context) error {
	if fw.Exists() {
		return fw.Update(ctx)
	}

	return fw.Insert(ctx)
}

// Delete deletes the FactoryWithdraw from the database.
func (fw *FactoryWithdraw) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fw._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryWithdrawTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fw.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fw.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fw.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	fw._deleted = true

	return nil
}

// FactoryWithdrawByID retrieves a row from 'ddg_local.factory_withdraw' as a FactoryWithdraw.
//
// Generated from index 'factory_withdraw_id_pkey'.
func FactoryWithdrawByID(ctx context.Context, id int64, key ...interface{}) (*FactoryWithdraw, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryWithdrawTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, account_name, account_no, bank_name, img_url, created_at, update_at, do_status, current_money, pre_money, post_money, do_time, remark ` +
		`FROM ` + tableName +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, id)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fw := FactoryWithdraw{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&fw.ID, &fw.Fid, &fw.AccountName, &fw.AccountNo, &fw.BankName, &fw.ImgURL, &fw.CreatedAt, &fw.UpdateAt, &fw.DoStatus, &fw.CurrentMoney, &fw.PreMoney, &fw.PostMoney, &fw.DoTime, &fw.Remark)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&fw.ID, &fw.Fid, &fw.AccountName, &fw.AccountNo, &fw.BankName, &fw.ImgURL, &fw.CreatedAt, &fw.UpdateAt, &fw.DoStatus, &fw.CurrentMoney, &fw.PreMoney, &fw.PostMoney, &fw.DoTime, &fw.Remark)
		if err != nil {
			return nil, err
		}
	}

	return &fw, nil
}
