function Log(...)
    base.Log(unpack(arg))
end

function SetInterval(tm, callback)
    return base.SetInterval(tm, callback)
end

function SetTimeout(tm, callback)
    return base.SetTimeout(tm, callback)
end

function SendCancelSig(sig)
    base.SendCancelSig(sig)
end

local _DB = {}
local _mtDB = {__index = _DB}
local _DBRows = {}
local _mtDBRows = {__index = _DBRows}
local _DBResult = {}
local _mtDBResult = {__index = _DBResult}

function OpenDB(dbpath)
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
        return ret
    end

    local ret = setmetatable({}, _mtDBResult)
    ret.dbResultPointer = dbResultPointer
    return ret
end

function _DBRows.FetchOne(self)
    return base.DBRowsNext(self.dbRowsPointer)
end

function _DBRows.Close(self)
    return base.DBRowsClose(self.dbRowsPointer)
end

function _DBResult.LastInsertId(self)
    return base.DBResultLastInsertId(self.dbResultPointer)
end

function _DBResult.RowsAffected(self)
    return base.DBResultRowsAffected(self.dbResultPointer)
end
