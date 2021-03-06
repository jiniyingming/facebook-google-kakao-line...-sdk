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

// Adv2PlacementMaterial represents a row from 'aypcddg.adv2_placement_material'.
type Adv2PlacementMaterial struct {
	ID           int            `json:"id"`            // id
	ImgURL       sql.NullString `json:"img_url"`       // img_url
	Width        sql.NullInt64  `json:"width"`         // width
	Height       sql.NullInt64  `json:"height"`        // height
	Target       sql.NullString `json:"target"`        // target
	Operator     sql.NullInt64  `json:"operator"`      // operator
	OperatorName sql.NullString `json:"operator_name"` // operator_name
	Start        mysql.NullTime `json:"start"`         // start
	End          mysql.NullTime `json:"end"`           // end
	PlaceID      sql.NullInt64  `json:"place_id"`      // place_id
	Title        sql.NullString `json:"title"`         // title
	Deleted      sql.NullInt64  `json:"deleted"`       // deleted
	Sort         sql.NullInt64  `json:"sort"`          // sort
	Created      sql.NullInt64  `json:"created"`       // created
	Updated      sql.NullInt64  `json:"updated"`       // updated
	IsUnlimited  sql.NullInt64  `json:"is_unlimited"`  // is_unlimited
	MaterialID   sql.NullInt64  `json:"material_id"`   // material_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Adv2PlacementMaterial exists in the database.
func (am *Adv2PlacementMaterial) Exists() bool { //adv2_placement_material
	return am._exists
}

// Deleted provides information if the Adv2PlacementMaterial has been deleted from the database.
func (am *Adv2PlacementMaterial) Deleted() bool {
	return am._deleted
}

// Get table name
func GetAdv2PlacementMaterialTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "adv2_placement_material", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Adv2PlacementMaterial to the database.
func (am *Adv2PlacementMaterial) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if am._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`img_url, width, height, target, operator, operator_name, start, end, place_id, title, deleted, sort, created, updated, is_unlimited, material_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID)
	} else {
		res, err = dbConn.Exec(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID)
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
	am.ID = int(id)
	am._exists = true

	return nil
}

// Update updates the Adv2PlacementMaterial in the database.
func (am *Adv2PlacementMaterial) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if am._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`img_url = ?, width = ?, height = ?, target = ?, operator = ?, operator_name = ?, start = ?, end = ?, place_id = ?, title = ?, deleted = ?, sort = ?, created = ?, updated = ?, is_unlimited = ?, material_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID, am.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID, am.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, am.ImgURL, am.Width, am.Height, am.Target, am.Operator, am.OperatorName, am.Start, am.End, am.PlaceID, am.Title, am.Deleted, am.Sort, am.Created, am.Updated, am.IsUnlimited, am.MaterialID, am.ID)
	}
	return err
}

// Save saves the Adv2PlacementMaterial to the database.
func (am *Adv2PlacementMaterial) Save(ctx context.Context) error {
	if am.Exists() {
		return am.Update(ctx)
	}

	return am.Insert(ctx)
}

// Delete deletes the Adv2PlacementMaterial from the database.
func (am *Adv2PlacementMaterial) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if am._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, am.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, am.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, am.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	am._deleted = true

	return nil
}

// Adv2PlacementMaterialByID retrieves a row from 'aypcddg.adv2_placement_material' as a Adv2PlacementMaterial.
//
// Generated from index 'adv2_placement_material_id_pkey'.
func Adv2PlacementMaterialByID(ctx context.Context, id int, key ...interface{}) (*Adv2PlacementMaterial, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAdv2PlacementMaterialTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, img_url, width, height, target, operator, operator_name, start, end, place_id, title, deleted, sort, created, updated, is_unlimited, material_id ` +
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
	am := Adv2PlacementMaterial{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&am.ID, &am.ImgURL, &am.Width, &am.Height, &am.Target, &am.Operator, &am.OperatorName, &am.Start, &am.End, &am.PlaceID, &am.Title, &am.Deleted, &am.Sort, &am.Created, &am.Updated, &am.IsUnlimited, &am.MaterialID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&am.ID, &am.ImgURL, &am.Width, &am.Height, &am.Target, &am.Operator, &am.OperatorName, &am.Start, &am.End, &am.PlaceID, &am.Title, &am.Deleted, &am.Sort, &am.Created, &am.Updated, &am.IsUnlimited, &am.MaterialID)
		if err != nil {
			return nil, err
		}
	}

	return &am, nil
}

// Adv2PlacementMaterialsByPlaceIDIsUnlimited retrieves a row from 'aypcddg.adv2_placement_material' as a Adv2PlacementMaterial.
//
// Generated from index 'idx_placeid_unlimited'.
func Adv2PlacementMaterialsByPlaceIDIsUnlimited(ctx context.Context, placeID sql.NullInt64, isUnlimited sql.NullInt64, key ...interface{}) ([]*Adv2PlacementMaterial, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAdv2PlacementMaterialTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, img_url, width, height, target, operator, operator_name, start, end, place_id, title, deleted, sort, created, updated, is_unlimited, material_id ` +
		`FROM ` + tableName +
		` WHERE place_id = ? AND is_unlimited = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, placeID, isUnlimited)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, placeID, isUnlimited)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, placeID, isUnlimited)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*Adv2PlacementMaterial, 0)
	for queryData.Next() {
		am := Adv2PlacementMaterial{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&am.ID, &am.ImgURL, &am.Width, &am.Height, &am.Target, &am.Operator, &am.OperatorName, &am.Start, &am.End, &am.PlaceID, &am.Title, &am.Deleted, &am.Sort, &am.Created, &am.Updated, &am.IsUnlimited, &am.MaterialID)
		if err != nil {
			return nil, err
		}

		res = append(res, &am)
	}

	return res, nil
}
