package.path = package.path .. ";" .. base.ResBaseDir .. "/lua/script/?.lua;"

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
