package script

import (
	"database/sql"
	"in/utils"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	lua "github.com/yuin/gopher-lua"
)

func (p *Script) _getDBPointerFromUserData(L *lua.LState, lu *lua.LUserData) *sql.DB {
	if nil == lu || nil == lu.Value {
		return nil
	}

	db, ok := lu.Value.(*sql.DB)
	if false == ok || nil == db {
		return nil
	}

	return db
}

func (p *Script) _getDBRowsPointerFromUserData(L *lua.LState, lu *lua.LUserData) *sql.Rows {
	if nil == lu || nil == lu.Value {
		return nil
	}

	rows, ok := lu.Value.(*sql.Rows)
	if false == ok || nil == rows {
		return nil
	}

	return rows
}

func (p *Script) _getDBResultPointerFromUserData(L *lua.LState, lu *lua.LUserData) sql.Result {
	if nil == lu || nil == lu.Value {
		return nil
	}

	rows, ok := lu.Value.(sql.Result)
	if false == ok || nil == rows {
		return nil
	}

	return rows
}

func (p *Script) OpenDB(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	dbpath := filepath.Join(GlobalOption.ResBaseDir, "project", GlobalOption.ProjectName, L.ToString(1))
	db, err := sql.Open("sqlite3", dbpath)
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	luaDB := L.NewUserData()
	luaDB.Value = db
	L.Push(luaDB)
	return 1
}

func (p *Script) CloseDB(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		return 0
	}

	db := p._getDBPointerFromUserData(L, L.ToUserData(1))
	if nil != db {
		db.Close()
	}
	return 0
}

func (p *Script) DBQuery(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 2 {
		L.Push(lua.LNil)
		return 1
	}

	db := p._getDBPointerFromUserData(L, L.ToUserData(1))
	if nil == db {
		L.Push(lua.LNil)
		return 1
	}
	query := L.ToString(2)

	rows, err := db.Query(query)
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	luaDBRows := L.NewUserData()
	luaDBRows.Value = rows
	L.Push(luaDBRows)
	return 1
}

func (p *Script) DBRowsNext(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	rows := p._getDBRowsPointerFromUserData(L, L.ToUserData(1))
	if nil == rows || false == rows.Next() {
		L.Push(lua.LNil)
		return 1
	}

	columns, _ := rows.Columns()
	_stringArray := make([]string, len(columns))
	dbRet := make(map[string]*string, len(columns))
	var dbRetForScan []interface{}
	table := L.NewTable()
	for k, column := range columns {
		dbRet[column] = &_stringArray[k]
		dbRetForScan = append(dbRetForScan, &_stringArray[k])
	}

	err := rows.Scan(dbRetForScan...)
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	for k, v := range dbRet {
		L.SetField(table, k, lua.LString(*v))
	}

	L.Push(table)
	return 1
}

func (p *Script) DBRowsClose(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		return 0
	}

	rows := p._getDBRowsPointerFromUserData(L, L.ToUserData(1))
	if nil == rows {
		return 0
	}

	rows.Close()
	return 0
}

func (p *Script) DBExec(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 2 {
		L.Push(lua.LNil)
		return 1
	}

	db := p._getDBPointerFromUserData(L, L.ToUserData(1))
	if nil == db {
		L.Push(lua.LNil)
		return 1
	}
	query := L.ToString(2)

	ret, err := db.Exec(query)
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	luaDBResult := L.NewUserData()
	luaDBResult.Value = ret
	L.Push(luaDBResult)
	return 1
}

func (p *Script) DBResultLastInsertId(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	dbResult := p._getDBResultPointerFromUserData(L, L.ToUserData(1))
	if nil == dbResult {
		L.Push(lua.LNil)
		return 1
	}

	ret, err := dbResult.LastInsertId()
	if nil != err {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNumber(ret))
	}
	return 1
}

func (p *Script) DBResultRowsAffected(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	dbResult := p._getDBResultPointerFromUserData(L, L.ToUserData(1))
	if nil == dbResult {
		L.Push(lua.LNil)
		return 1
	}

	ret, err := (dbResult).RowsAffected()
	if nil != err {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNumber(ret))
	}
	return 1

}
