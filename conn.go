// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go_ibm_db

import (
	"database/sql/driver"
	"runtime"
	"unsafe"
	"fmt"
	"github.com/ibmdb/go_ibm_db/api"
	trc "github.com/ibmdb/go_ibm_db/log2"
)

type Conn struct {
	h  api.SQLHDBC
	tx *Tx
}

func (d *Driver) Open(dsn string) (driver.Conn, error) {
	trc.Trace1("conn.go: Open() - ENTRY")
	trc.Trace1(fmt.Sprintf("dsn = %s", dsn))

	var out api.SQLHANDLE
	ret := api.SQLAllocHandle(api.SQL_HANDLE_DBC, api.SQLHANDLE(d.h), &out)
	if IsError(ret) {
		return nil, NewError("SQLAllocHandle", d.h)
	}
	h := api.SQLHDBC(out)
	drv.Stats.updateHandleCount(api.SQL_HANDLE_DBC, 1)

	b := api.StringToUTF16(dsn)
	if runtime.GOOS == "zos" {
		ret = api.SQLDriverConnect(h, 0,
			(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQLSMALLINT(2*len(b)), // odbc api on zos doesn't handle null terminated strings, the exact size is passed
			nil, 0, nil, api.SQL_DRIVER_NOPROMPT)
	} else {
		ret = api.SQLDriverConnect(h, 0,
			(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQLSMALLINT(len(b)),
			nil, 0, nil, api.SQL_DRIVER_NOPROMPT)
	}
	if IsError(ret) {
		defer releaseHandle(h)
		return nil, NewError("SQLDriverConnect", h)
	}
	trc.Trace1("conn.go: Open() - EXIT")
	return &Conn{h: h}, nil
}

func (c *Conn) Close() error {
	trc.Trace1("conn.go: Close() - ENTRY")

	ret := api.SQLDisconnect(c.h)
	if IsError(ret) {
		return NewError("SQLDisconnect", c.h)
	}
	h := c.h
	c.h = api.SQLHDBC(api.SQL_NULL_HDBC)
	trc.Trace1("conn.go: Close() - EXIT")
	return releaseHandle(h)
}

//Query method executes the statement with out prepare if no args provided, and a driver.ErrSkip otherwise (handled by sql.go to execute usual preparedStmt)
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	trc.Trace1("conn.go: Query() - ENTRY")
	trc.Trace1(fmt.Sprintf("query = %s", query))

	if len(args) > 0 {
		// Not implemented for queries with parameters
		return nil, driver.ErrSkip
	}
	var out api.SQLHANDLE
	var os *ODBCStmt
	ret := api.SQLAllocHandle(api.SQL_HANDLE_STMT, api.SQLHANDLE(c.h), &out)
	if IsError(ret) {
		return nil, NewError("SQLAllocHandle", c.h)
	}
	h := api.SQLHSTMT(out)
	drv.Stats.updateHandleCount(api.SQL_HANDLE_STMT, 1)
	b := api.StringToUTF16(query)
	ret = api.SQLExecDirect(h,
		(*api.SQLWCHAR)(unsafe.Pointer(&b[0])), api.SQL_NTS)
	if IsError(ret) {
		defer releaseHandle(h)
		return nil, NewError("SQLExecDirectW", h)
	}
	ps, err := ExtractParameters(h)
	if err != nil {
		defer releaseHandle(h)
		return nil, err
	}
	os = &ODBCStmt{
		h:          h,
		Parameters: ps,
		usedByRows: true}
	err = os.BindColumns()
	if err != nil {
		return nil, err
	}
	trc.Trace1("conn.go: Query() - EXIT")
	return &Rows{os: os}, nil
}
