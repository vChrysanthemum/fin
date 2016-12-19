local json = require("json")

function SetCurrentTask(id)
    if nil == id then
        GCurrentTask = nil
    else 
        GCurrentTask = GetTaskFromGTasksById(id)
        if nil == GCurrentTask then
            return
        end
    end

    RefreshGaugeProgress()
end

function GetTaskFromGTasksById(id)
    for k, v in pairs(GTasks) do
        if v.Id == id then
            return v
        end
    end
    return nil
end

function FlushGTasksToDB()
    WriteContentToFile("main.db", json.encode({
        GTasksLastId = GTasksLastId,
        GTasks       = GTasks,
    }))
end

function ReadGTasksFromDB()
    GTasksLastId = 0
    GTasks = {}
    local ret = json.decode(ReadContentFromFile("main.db"))
    if "table" == type(ret) then
        GTasksLastId = ret.GTasksLastId
        GTasks = ret.GTasks
    end
    if nil == ret.GTasks then
        GTasks = {}
    end
end

function CommandNewTask(nodePointer, command)
    local commandArr = StringSplit(command, " ")
    if nil == commandArr or TableLength(commandArr) < 2 then
        return
    end

    table.remove(commandArr, 1)
    local taskName = table.concat(commandArr, " ")

    GTasksLastId = GTasksLastId + 1
    GCurrentTask = {
        Id       = tostring(GTasksLastId),
        Name     = taskName,
        StartAt  = TimeNow(),
        Percent  = 0,
    }
    table.insert(GTasks, GCurrentTask)

    NodeSelectTasks:SelectAppendOption(GCurrentTask.Id, GCurrentTask.Name)
    NodeSelectTasks:SetValue(GCurrentTask.Id)

    SetCurrentTask(GCurrentTask.Id)
    RefreshNodeSelectTasksBorderLabel()

    NodeTerminalTask:TerminalClearLines()
end

function CommandDeleteTask(nodePointer)
    if nil == GCurrentTask then
        return
    end

    for k, v in pairs(GTasks) do
        if v.Id == GCurrentTask.Id then
            table.remove(GTasks, k)
            break
        end
    end
    GCurrentTask = nil

    RefreshNodeSelectTasksOptions()
    RefreshGaugeProgress()
end

function CommandEditTaskPercent(nodePointer, command)
    local commandArr = StringSplit(command, " ")
    if nil == commandArr or TableLength(commandArr) < 2 then
        return
    end

    if nil == GCurrentTask then
        return
    end
    GCurrentTask.Percent = commandArr[2]
    RefreshGaugeProgress()
end

function CommandRenameTask(nodePointer, command)
    if nil == GCurrentTask then
        return
    end

    local commandArr = StringSplit(command, " ")
    if nil == commandArr or TableLength(commandArr) < 2 then
        return
    end

    table.remove(commandArr, 1)
    local taskName = table.concat(commandArr, " ")
    GCurrentTask.Name = taskName
    NodeSelectTasks:SelectSetOptionData(GCurrentTask.Id, GCurrentTask.Name)

    SetCurrentTask(GCurrentTask.Id)
end

function CommandClearGTasks(nodePointer)
    GTasksLastId = 0
    GTasks = {}
    GCurrentTask = nil
    NodeSelectTasks:SelectClearOptions()
    RefreshNodeSelectTasksBorderLabel()
    RefreshGaugeProgress()
end
