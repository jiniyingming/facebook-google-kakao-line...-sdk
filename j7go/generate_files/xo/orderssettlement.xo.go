// Package xo contains the types for schema 'aypcddg'.
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

// OrdersSettlement represents a row from 'aypcddg.orders_settlement'.
type OrdersSettlement struct {
	ID            int64           `json:"id"`              // id
	OrdersID      sql.NullInt64   `json:"orders_id"`       // orders_id
	DoStatus      sql.NullInt64   `json:"do_status"`       // do_status
	DoDatetime    mysql.NullTime  `json:"do_datetime"`     // do_datetime
	OrdersMoney   sql.NullFloat64 `json:"orders_money"`    // orders_money
	AfterMoney    sql.NullFloat64 `json:"after_money"`     // after_money
	AfterID       sql.NullInt64   `json:"after_id"`        // after_id
	RealMoney     sql.NullFloat64 `json:"real_money"`      // real_money
	CreatedAt     mysql.NullTime  `json:"created_at"`      // created_at
	UpdateAt      mysql.NullTime  `json:"update_at"`       // update_at
	UserID        sql.NullInt64   `json:"user_id"`         // user_id
	Shipping      sql.NullFloat64 `json:"shipping"`        // shipping
	RealShipping  sql.NullFloat64 `json:"real_shipping"`   // real_shipping
	FactoryBiilID sql.NullInt64   `json:"factory_biil_id"` // factory_biil_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OrdersSettlement exists in the database.
func (os *OrdersSettlement) Exists() bool { //orders_settlement
	return os._exists
}

// Deleted provides information if the OrdersSettlement has been deleted from the database.
func (os *OrdersSettlement) Deleted() bool {
	return os._deleted
}

// Get table name
func GetOrdersSettlementTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "orders_settlement", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the OrdersSettlement to the database.
func (os *OrdersSettlement) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if os._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrdersSettlementTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`orders_id, do_status, do_datetime, orders_money, after_money, after_id, real_money, created_at, update_at, user_id, shipping, real_shipping, factory_biil_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID)
	} else {
		res, err = dbConn.Exec(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID)
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
	os.ID = int64(id)
	os._exists = true

	return nil
}

// Update updates the OrdersSettlement in the database.
func (os *OrdersSettlement) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if os._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrdersSettlementTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`orders_id = ?, do_status = ?, do_datetime = ?, orders_money = ?, after_money = ?, after_id = ?, real_money = ?, created_at = ?, update_at = ?, user_id = ?, shipping = ?, real_shipping = ?, factory_biil_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID, os.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID, os.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, os.OrdersID, os.DoStatus, os.DoDatetime, os.OrdersMoney, os.AfterMoney, os.AfterID, os.RealMoney, os.CreatedAt, os.UpdateAt, os.UserID, os.Shipping, os.RealShipping, os.FactoryBiilID, os.ID)
	}
	return err
}

// Save saves the OrdersSettlement to the database.
func (os *OrdersSettlement) Save(ctx context.Context) error {
	if os.Exists() {
		return os.Update(ctx)
	}

	return os.Insert(ctx)
}

// Delete deletes the OrdersSettlement from the database.
func (os *OrdersSettlement) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if os._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrdersSettlementTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, os.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, os.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, os.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	os._deleted = true

	return nil
}

// OrdersSettlementByID retrieves a row from 'aypcddg.orders_settlement' as a OrdersSettlement.
//
// Generated from index 'orders_settlement_id_pkey'.
func OrdersSettlementByID(ctx context.Context, id int64, key ...interface{}) (*OrdersSettlement, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrdersSettlementTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, orders_id, do_status, do_datetime, orders_money, after_money, after_id, real_money, created_at, update_at, user_id, shipping, real_shipping, factory_biil_id ` +
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
	os := OrdersSettlement{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&os.ID, &os.OrdersID, &os.DoStatus, &os.DoDatetime, &os.OrdersMoney, &os.AfterMoney, &os.AfterID, &os.RealMoney, &os.CreatedAt, &os.UpdateAt, &os.UserID, &os.Shipping, &os.RealShipping, &os.FactoryBiilID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&os.ID, &os.OrdersID, &os.DoStatus, &os.DoDatetime, &os.OrdersMoney, &os.AfterMoney, &os.AfterID, &os.RealMoney, &os.CreatedAt, &os.UpdateAt, &os.UserID, &os.Shipping, &os.RealShipping, &os.FactoryBiilID)
		if err != nil {
			return nil, err
		}
	}

	return &os, nil
}
