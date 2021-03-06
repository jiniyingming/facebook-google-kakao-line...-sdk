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

// SettlementWalletOrder represents a row from 'ddg_local.settlement_wallet_order'.
type SettlementWalletOrder struct {
	ID                 int64           `json:"id"`                   // id
	AssetID            sql.NullInt64   `json:"asset_id"`             // asset_id
	SettlementWalletID sql.NullInt64   `json:"settlement_wallet_id"` // settlement_wallet_id
	OrdersID           sql.NullInt64   `json:"orders_id"`            // orders_id
	OrdersMoney        sql.NullFloat64 `json:"orders_money"`         // orders_money
	NeedMoney          sql.NullFloat64 `json:"need_money"`           // need_money
	RealMoney          sql.NullFloat64 `json:"real_money"`           // real_money
	DoStatus           sql.NullInt64   `json:"do_status"`            // do_status
	CreatedAt          mysql.NullTime  `json:"created_at"`           // created_at
	UpdateAt           mysql.NullTime  `json:"update_at"`            // update_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the SettlementWalletOrder exists in the database.
func (swo *SettlementWalletOrder) Exists() bool { //settlement_wallet_order
	return swo._exists
}

// Deleted provides information if the SettlementWalletOrder has been deleted from the database.
func (swo *SettlementWalletOrder) Deleted() bool {
	return swo._deleted
}

// Get table name
func GetSettlementWalletOrderTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "settlement_wallet_order", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the SettlementWalletOrder to the database.
func (swo *SettlementWalletOrder) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if swo._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSettlementWalletOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`asset_id, settlement_wallet_id, orders_id, orders_money, need_money, real_money, do_status, created_at, update_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt)
	} else {
		res, err = dbConn.Exec(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt)
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
	swo.ID = int64(id)
	swo._exists = true

	return nil
}

// Update updates the SettlementWalletOrder in the database.
func (swo *SettlementWalletOrder) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if swo._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSettlementWalletOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`asset_id = ?, settlement_wallet_id = ?, orders_id = ?, orders_money = ?, need_money = ?, real_money = ?, do_status = ?, created_at = ?, update_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt, swo.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt, swo.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, swo.AssetID, swo.SettlementWalletID, swo.OrdersID, swo.OrdersMoney, swo.NeedMoney, swo.RealMoney, swo.DoStatus, swo.CreatedAt, swo.UpdateAt, swo.ID)
	}
	return err
}

// Save saves the SettlementWalletOrder to the database.
func (swo *SettlementWalletOrder) Save(ctx context.Context) error {
	if swo.Exists() {
		return swo.Update(ctx)
	}

	return swo.Insert(ctx)
}

// Delete deletes the SettlementWalletOrder from the database.
func (swo *SettlementWalletOrder) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if swo._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSettlementWalletOrderTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, swo.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, swo.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, swo.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	swo._deleted = true

	return nil
}

// SettlementWalletOrderByID retrieves a row from 'ddg_local.settlement_wallet_order' as a SettlementWalletOrder.
//
// Generated from index 'settlement_wallet_order_id_pkey'.
func SettlementWalletOrderByID(ctx context.Context, id int64, key ...interface{}) (*SettlementWalletOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSettlementWalletOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, asset_id, settlement_wallet_id, orders_id, orders_money, need_money, real_money, do_status, created_at, update_at ` +
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
	swo := SettlementWalletOrder{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&swo.ID, &swo.AssetID, &swo.SettlementWalletID, &swo.OrdersID, &swo.OrdersMoney, &swo.NeedMoney, &swo.RealMoney, &swo.DoStatus, &swo.CreatedAt, &swo.UpdateAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&swo.ID, &swo.AssetID, &swo.SettlementWalletID, &swo.OrdersID, &swo.OrdersMoney, &swo.NeedMoney, &swo.RealMoney, &swo.DoStatus, &swo.CreatedAt, &swo.UpdateAt)
		if err != nil {
			return nil, err
		}
	}

	return &swo, nil
}
