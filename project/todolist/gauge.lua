function RefreshGaugeProgress()
    if nil == GCurrentTask then
        NodeGaugeProgress:SetAttribute("borderlabel", "")
        NodeGaugeProgress:SetAttribute("percent", tostring(0))
        return
    end

    NodeGaugeProgress:SetAttribute("borderlabel", string.format(" %s ", GCurrentTask.Name))
    NodeGaugeProgress:SetAttribute("percent", tostring(math.ceil(GCurrentTask.Percent * 100)))

    local percent = tonumber(GCurrentTask.Percent)

    if percent < 0.3 then
        NodeGaugeProgress:SetAttribute("barcolor", "red")
        NodeGaugeProgress:SetAttribute("percentcolor_highlighted", "black")
    elseif percent < 0.6 then
        NodeGaugeProgress:SetAttribute("barcolor", "yellow")
        NodeGaugeProgress:SetAttribute("percentcolor_highlighted", "black")
    elseif  percent < 0.9 then
        NodeGaugeProgress:SetAttribute("barcolor", "green")
        NodeGaugeProgress:SetAttribute("percentcolor_highlighted", "black")
    else
        NodeGaugeProgress:SetAttribute("barcolor", "blue")
        NodeGaugeProgress:SetAttribute("percentcolor_highlighted", "white")
    end
end
