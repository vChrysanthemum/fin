local unixsock = {}

local _UnixSock = {}
local _mtUnixSock = {__index = _UnixSock}
local _UnixSockRows = {}
local _mtUnixSockRows = {__index = _UnixSockRows}
local _UnixSockResult = {}
local _mtUnixSockResult = {__index = _UnixSockResult}

function unixsock.NewUnixSockClient(unixSockClientPath)
    local ret = setmetatable({}, _mtUnixSock)
    ret.unixSockClientPointer = base.NewUnixSockClient(unixSockClientPath)
    return ret
end

function _UnixSock.Get(self, urlStr)
    return base.UnixSockGet(self.unixSockClientPointer, urlStr)
end 

return unixsock
