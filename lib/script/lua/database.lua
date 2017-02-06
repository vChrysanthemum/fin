local database = {}

local _DB = {}
local _mtDB = {__index = _DB}
local _DBRows = {}
local _mtDBRows = {__index = _DBRows}
local _DBResult = {}
local _mtDBResult = {__index = _DBResult}

function database.OpenDB(dbpath)
    local dbPointer
    dbPointer = base.OpenDB(dbpath)
    local ret = setmetatable({}, _mtDB)
    ret.dbPointer = dbPointer
    return ret
end 

function _DB.Query(self, query)
    local dbRowsPointer = base.DBQuery(self.dbPointer, query)
    local retType = type(dbRowsPointer)
    if "userdata" ~= retType then 
        return ret
    end

    local ret = setmetatable({}, _mtDBRows)
    ret.dbRowsPointer = dbRowsPointer
    return ret
end

function _DB.Exec(self, query)
    local dbResultPointer = base.DBExec(self.dbPointer, query)
    local retType = type(dbResultPointer)
    if "userdata" ~= retType then 
        return dbResultPointer
    end

    local ret = setmetatable({}, _mtDBResult)
    ret.dbResultPointer = dbResultPointer
    return ret
end

function _DB.QuoteSQL(self, sql)
    local ret = string.gsub(sql, "\\", "\\\\")
    ret = string.gsub(ret, "'", "\\'")
    return ret
end

function _DBRows.FetchOne(self)
    return base.DBRowsNext(self.dbRowsPointer)
end

function _DBRows.Close(self)
    return base.DBRowsClose(self.dbRowsPointer)
end

function _DBResult.LastInsertID(self)
    return base.DBResultLastInsertID(self.dbResultPointer)
end

function _DBResult.RowsAffected(self)
    return base.DBResultRowsAffected(self.dbResultPointer)
end

return database
