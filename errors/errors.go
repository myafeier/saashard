// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNoMasterConn  = errors.New("no master connection")
	ErrNoSlaveConn   = errors.New("no slave connection")
	ErrNoDefaultNode = errors.New("no default node")
	ErrNoMasterDB    = errors.New("no master database")
	ErrNoSlaveDB     = errors.New("no slave database")
	ErrNoDatabase    = errors.New("no database")
	ErrNoIdleConn    = errors.New("exceed max conn num")

	ErrNoDataHost = errors.New("no data host")
	ErrNoDataNode = errors.New("no data node")
	ErrNoSchema   = errors.New("no schema")
	ErrNoTable    = errors.New("no table")

	ErrNoRouteNode = errors.New("no route node")
	ErrNoPlan      = errors.New("statement have no plan")
	ErrNoStatement = errors.New("no statement to execute")

	ErrSchemaNotExists   = errors.New("schema not exists")
	ErrDatabaseNotExists = errors.New("database not exists")
	ErrTableNotExists    = errors.New("table not exists")

	ErrMasterDown    = errors.New("master is down")
	ErrSlaveDown     = errors.New("slave is down")
	ErrDatabaseClose = errors.New("database is close")
	ErrConnIsNil     = errors.New("connection is nil")
	ErrBadConn       = errors.New("connection was bad")
	ErrIgnoreSQL     = errors.New("ignore this sql")

	ErrAddressNull     = errors.New("address is nil")
	ErrInvalidArgument = errors.New("argument is invalid")
	ErrInvalidCharset  = errors.New("charset is invalid")
	ErrCmdUnsupport    = errors.New("command unsupport")

	ErrSelectInInsert   = errors.New("select in insert not allowed")
	ErrTransInMulti     = errors.New("transaction in multi node")
	ErrExecInMulti      = errors.New("execute in multi node")
	ErrColsLenNotMatch  = errors.New("insert or replace cols and values length not match")
	ErrInsertColumnsKey = errors.New("no shard key in insert column list")
	ErrInsertValuesKey  = errors.New("no shard key or key has different values in insert values list")
	ErrUpdateKey        = errors.New("shard key in update expression")
	ErrWhereOrJoinOnKey = errors.New("no shard key or key has different values in where or join on expression")

	ErrMustPositiveIntegerInModShard = errors.New("shard key must positive integer when use mod shard algorithm")

	ErrStmtConvert      = errors.New("statement fail to convert")
	ErrExprConvert      = errors.New("expr fail to convert")
	ErrConnNotEqual     = errors.New("the length of conns not equal sqls")
	ErrKeyOutOfRange    = errors.New("shard key not in key range")
	ErrDateIllegal      = errors.New("date format illegal")
	ErrDateRangeIllegal = errors.New("date range format illegal")
	ErrDateRangeCount   = errors.New("date range count is not equal")
	ErrSlaveExist       = errors.New("slave has exist")
	ErrSlaveNotExist    = errors.New("slave has not exist")

	ErrMalformPacket = errors.New("Malform packet error")
	ErrTxDone        = errors.New("Transaction has already been committed or rolled back")
)

type SqlError struct {
	Code    uint16
	Message string
	State   string
}

func (e *SqlError) Error() string {
	return fmt.Sprintf("ERROR %d (%s): %s", e.Code, e.State, e.Message)
}

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}
