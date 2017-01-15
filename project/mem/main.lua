NodeLineChart = Node("LineChart")
Data = {}
LoopEventSig = SetInterval(200, function()
    table.insert(Data, base.GetMemAlloc())
    NodeLineChart:SetValue(table.concat(Data, ","))
end)
