package script

import (
	"fmt"
	"in/script"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func TestDB(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	dbpath := "test.db"
	os.Remove(filepath.Join(GlobalOption.ResBaseDir, "project", GlobalOption.ProjectName, dbpath))

	var (
		err error
		s   script.Script
	)

	L := lua.NewState()
	s.RegisterScript(L)
	luaBase := L.NewTable()
	L.SetGlobal("base", luaBase)
	s.RegisterBaseTable(L, luaBase)

	err = L.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/script/core.lua"))
	assert.Nil(t, err)

	content := fmt.Sprintf(`
	local database = require("database")
	local db = database.OpenDB("%s")
	local dbRet = db:Exec([[
		create table b_test (
			test_id integer primary key not null,
			data varchar(64)
		);
	]])
	dbRet = db:Exec("insert into b_test (data) values('testdata1');")
	dbRet = db:Exec("insert into b_test (data) values('testdata2');")
	local rows = db:Query("select test_id, data from b_test")
	local rowRet, rowRetType
	local num = 0
	repeat
		rowRet = rows:FetchOne()
		rowRetType = type(rowRet)
		if "table" ~= rowRetType then
			break
		end
		num = num + 1
	until false 
	if num ~= 2 then
		panic()
	end
	`, dbpath)
	err = L.DoString(content)
	assert.Nil(t, err)

	os.Remove(filepath.Join(GlobalOption.ResBaseDir, "project", GlobalOption.ProjectName, dbpath))
}
