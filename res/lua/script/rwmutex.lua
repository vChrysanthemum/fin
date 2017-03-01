local rwmutex = {}

local _Mutex = {}
local _mtMutex = {__index = _Mutex}

function rwmutex.NewRWMutex(mutexpath)
    local mutexPointer
    mutexPointer = base.NewRWMutex()
    local ret = setmetatable({}, _mtMutex)
    ret.mutexPointer = mutexPointer
    return ret
end

function _Mutex.Lock(self)
    return base.RWMutexLock(self.mutexPointer)
end

function _Mutex.Unlock(self)
    return base.RWMutexUnlock(self.mutexPointer)
end

return rwmutex
