GCurrentTask = {}
GTasksLastId = nil
GTasks = nil
ReadGTasksFromDB()

NodeTerminalTask = Node("TerminalTask")
NodeSelectTasks = Node("SelectTasks")
NodeGaugeProgress = Node("GaugeProgress")

RefreshNodeSelectTasksOptions()

NodeTerminalTaskCommandSig = NodeTerminalTask:TerminalRegisterCommandHandle(function(nodePointer, command)
    NodeTerminalTask:TerminalClearLines()

    local commandArr = StringSplit(command, " ")
    if nil == commandArr or 0 == TableLength(commandArr) then
        return
    end

    if "new" == commandArr[1] or "n" == commandArr[1] then
        CommandNewTask(nodePointer, command)

    elseif "delete" == commandArr[1] or "d" == commandArr[1] then
        CommandDeleteTask()

    elseif "percent" == commandArr[1] or "p" == commandArr[1] then
        CommandEditTaskPercent(nodePointer, command)

    elseif "rename" == commandArr[1] or "r" == commandArr[1] then
        CommandRenameTask(nodePointer, command)

    elseif "select" == commandArr[1] or "s" == commandArr[1] then
        NodeSelectTasks:SetActive()
        NodeSelectTasksKeyPressSkipEnter = true

    elseif "clear" == commandArr[1] or "c" == commandArr[1] then
        CommandClearGTasks(nodePointer)

    elseif "quit" == commandArr[1] or "q" == commandArr[1] then
        FlushGTasksToDB()
        Quit()
    end

    FlushGTasksToDB()
end)

NodeSelectTasksKeyPressSkipEnter = false
NodeSelectTasksKeyPressKeySig = NodeSelectTasks:RegisterKeyPressHandler(function(nodePointer, KeyStr)
    if "<enter>" == KeyStr and true == NodeSelectTasksKeyPressSkipEnter then
        NodeSelectTasksKeyPressSkipEnter = false
        return
    end

    SetCurrentTask(NodeSelectTasks:GetValue())

    if "<enter>" == KeyStr then
        NodeTerminalTask:SetActive()
    end
end)
