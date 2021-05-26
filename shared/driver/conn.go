// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package driver

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// Conn is a DB driver conn implementation
// which will just pass-through all queries to its
// underlying connection.
type Conn struct {
	conn *sql.Conn
}

// driverConn is a super-interface combining the basic
// driver.Conn interface with some new additions later.
type driverConn interface {
	driver.Conn
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Execer
	driver.ExecerContext
	driver.Queryer
	driver.QueryerContext
	driver.Pinger
}

var (
	// Compile-time check to ensure Conn implements the interface.
	_ driverConn = &Conn{}
)

func (c *Conn) Begin() (tx driver.Tx, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		tx, err = innerConn.(driver.Conn).Begin()
		return err
	})
	return tx, err
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (tx driver.Tx, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		tx, err = innerConn.(driver.ConnBeginTx).BeginTx(ctx, opts)
		return err
	})
	return tx, err
}

func (c *Conn) Prepare(q string) (stmt driver.Stmt, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		stmt, err = innerConn.(driver.Conn).Prepare(q)
		return err
	})
	return stmt, err
}

func (c *Conn) PrepareContext(ctx context.Context, q string) (stmt driver.Stmt, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		stmt, err = innerConn.(driver.ConnPrepareContext).PrepareContext(ctx, q)
		return err
	})
	return stmt, err
}

func (c *Conn) Exec(q string, args []driver.Value) (res driver.Result, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		res, err = innerConn.(driver.Execer).Exec(q, args)
		return err
	})
	return res, err
}

func (c *Conn) ExecContext(ctx context.Context, q string, args []driver.NamedValue) (res driver.Result, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		res, err = innerConn.(driver.ExecerContext).ExecContext(ctx, q, args)
		return err
	})
	return res, err
}

func (c *Conn) Query(q string, args []driver.Value) (rows driver.Rows, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		rows, err = innerConn.(driver.Queryer).Query(q, args)
		return err
	})
	return rows, err
}

func (c *Conn) QueryContext(ctx context.Context, q string, args []driver.NamedValue) (rows driver.Rows, err error) {
	err = c.conn.Raw(func(innerConn interface{}) error {
		rows, err = innerConn.(driver.QueryerContext).QueryContext(ctx, q, args)
		return err
	})
	return rows, err
}

func (c *Conn) Ping(ctx context.Context) error {
	return c.conn.Raw(func(innerConn interface{}) error {
		return innerConn.(driver.Pinger).Ping(ctx)
	})
}

func (c *Conn) Close() error {
	return c.conn.Raw(func(innerConn interface{}) error {
		return innerConn.(driver.Conn).Close()
	})
}
