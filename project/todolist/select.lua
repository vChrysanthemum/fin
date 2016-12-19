function RefreshNodeSelectTasksOptions()
    NodeSelectTasks:SelectClearOptions()
    for k, v in pairs(GTasks) do
        NodeSelectTasks:SelectAppendOption(v.Id, v.Name)
        GCurrentTask = v
    end

    if nil == GCurrentTask then
        SetCurrentTask(nil)
    else 
        NodeSelectTasks:SetValue(GCurrentTask.Id)
        SetCurrentTask(GCurrentTask.Id)
    end

    RefreshNodeSelectTasksBorderLabel()
end

function RefreshNodeSelectTasksBorderLabel()
    NodeSelectTasks:SetAttribute("borderlabel", string.format(" 任务总量:%d ", TableLength(GTasks)))
end
