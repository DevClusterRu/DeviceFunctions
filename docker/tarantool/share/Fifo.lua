function take(space, from, to)
    local found = box.space[space].index.status:pairs({from},{ iterator = 'EQ' }):nth(1)
    if found then
        return box.space[space]:update( {found.id}, {{'=', 2, to }})
    end
    return
end

function getJob(space, priority, dev_id, platform, carrier)
    for _, tuple in
        box.space[space].index.device_id:pairs(dev_id,{iterator = "REQ"}) do
        if (tuple.status == 's') then
            return box.space[space]:update( {tuple.id}, {{'=', 2, 't' }})
        end
    end

    local tuple = box.space[space].index.priority_platform_carrier_status:pairs({priority, platform, carrier, "s"},{iterator = "REQ"}):nth(1)
    if tuple then
        return box.space[space]:update( {tuple.id}, {{'=', 2, 't' }})
    end
end
