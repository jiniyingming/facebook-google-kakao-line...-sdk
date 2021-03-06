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

// WithdrawRecord represents a row from 'ddg_local.withdraw_record'.
type WithdrawRecord struct {
	ID         int64          `json:"id"`          // id
	WithdrawID sql.NullInt64  `json:"withdraw_id"` // withdraw_id
	Content    sql.NullString `json:"content"`     // content
	Title      sql.NullString `json:"title"`       // title
	CreatedAt  mysql.NullTime `json:"created_at"`  // created_at
	UpdateAt   mysql.NullTime `json:"update_at"`   // update_at
	AssetID    int64          `json:"asset_id"`    // asset_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WithdrawRecord exists in the database.
func (wr *WithdrawRecord) Exists() bool { //withdraw_record
	return wr._exists
}

// Deleted provides information if the WithdrawRecord has been deleted from the database.
func (wr *WithdrawRecord) Deleted() bool {
	return wr._deleted
}

// Get table name
func GetWithdrawRecordTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "withdraw_record", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the WithdrawRecord to the database.
func (wr *WithdrawRecord) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if wr._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWithdrawRecordTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`withdraw_id, content, title, created_at, update_at, asset_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID)
	} else {
		res, err = dbConn.Exec(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID)
	}

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wr.ID = int64(id)
	wr._exists = true

	return nil
}

// Update updates the WithdrawRecord in the database.
func (wr *WithdrawRecord) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWithdrawRecordTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`withdraw_id = ?, content = ?, title = ?, created_at = ?, update_at = ?, asset_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID, wr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID, wr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wr.WithdrawID, wr.Content, wr.Title, wr.CreatedAt, wr.UpdateAt, wr.AssetID, wr.ID)
	}
	return err
}

// Save saves the WithdrawRecord to the database.
func (wr *WithdrawRecord) Save(ctx context.Context) error {
	if wr.Exists() {
		return wr.Update(ctx)
	}

	return wr.Insert(ctx)
}

// Delete deletes the WithdrawRecord from the database.
func (wr *WithdrawRecord) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wr._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWithdrawRecordTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wr.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	wr._deleted = true

	return nil
}

// WithdrawRecordByID retrieves a row from 'ddg_local.withdraw_record' as a WithdrawRecord.
//
// Generated from index 'withdraw_record_id_pkey'.
func WithdrawRecordByID(ctx context.Context, id int64, key ...interface{}) (*WithdrawRecord, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetWithdrawRecordTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, withdraw_id, content, title, created_at, update_at, asset_id ` +
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
	wr := WithdrawRecord{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&wr.ID, &wr.WithdrawID, &wr.Content, &wr.Title, &wr.CreatedAt, &wr.UpdateAt, &wr.AssetID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&wr.ID, &wr.WithdrawID, &wr.Content, &wr.Title, &wr.CreatedAt, &wr.UpdateAt, &wr.AssetID)
		if err != nil {
			return nil, err
		}
	}

	return &wr, nil
}
