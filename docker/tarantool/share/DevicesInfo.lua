devsInfo = box.schema.space.create('DevicesInfo', { if_not_exists = true })

devsInfo:format({
    { name = 'id', type = 'unsigned' },
    { name = 'deviceId', type = 'unsigned' },
    { name = 'deviceNumber', type = 'string' },
    { name = 'serial', type = 'string' },
    { name = 'imei', type = 'string' },
})

box.schema.sequence.create('devs_info_sequence', { min = 1, start = 1, if_not_exists = true })
devsInfo:create_index('id', { sequence = 'devs_info_sequence', if_not_exists = true })
devsInfo:create_index('device_id', { unique = false, type = 'TREE', parts = { 'device_id' }, if_not_exists = true })
devsInfo:create_index('imei', { unique = false, type = 'TREE', parts = { 'imei' }, if_not_exists = true })

function getDeviceIdByImei(space, imei)
    return box.space['DeviceInfo']:iterator(imei, {iterator='EQ'})
end