q = box.schema.space.create('CallsQueue', { if_not_exists =true })

q:format({
{ name = 'id', type = 'integer' },
{ name = 'status', type = 'string' },
{ name = 'ttl', type = 'integer', is_nullable = true},
{ name = 'attempts', type = 'integer', is_nullable = true},

{ name = 'number', type = 'string'},
{ name = 'device_id', type = 'unsigned'},
{ name = 'platform', type = 'string'},
{ name = 'carrier', type = 'string'},
{ name = 'custom', type = 'string'},
{ name = 'priority', type = 'string'},
})


box.schema.sequence.create('dcCacheSequence',{min=1, start=1, if_not_exists = true})
q:create_index('id', {sequence='dcCacheSequence', if_not_exists = true})

setBaseIndex('DCCache')

q:create_index('device_id', {
unique = false,
type = 'TREE',
parts = { 'device_id' },
if_not_exists = true,
})

q:create_index('priority_platform_carrier_status', {
unique = false,
type = 'TREE',
parts = { 'priority', 'platform', 'carrier', 'status' },
if_not_exists = true,
})

q:create_index('priority', {
unique = false,
type = 'TREE',
parts = { 'priority' },
if_not_exists = true,
})


