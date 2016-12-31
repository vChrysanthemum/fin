_RobotEngineer = {}
local _mtRobotEngineer = {__index = _RobotEngineer} 

function NewRobotEngineer()
    local RobotEngineer = setmetatable({}, _mtRobotEngineer)

    return RobotEngineer
end
