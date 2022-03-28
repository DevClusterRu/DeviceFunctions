devS = box.schema.space.create('DevStatus', { if_not_exists =true })

devS:format({
{ name = 'id', type = 'unsigned' },
{ name = 'status', type = 'string' },
{ name = 'ttl', type = 'unsigned', is_nullable = true},
{ name = 'attempts', type = 'unsigned', is_nullable = true},

{ name = 'serial', type = 'string'},
{ name = 'deviceId', type = 'unsigned'},
{ name = 'devCallStatus', type = 'string'},
{ name = 'incomingNumber', type = 'string'},
{ name = 'deviceNumber', type = 'string'},
{ name = 'imei', type = 'string'},
{ name = 'callSession', type = 'string'},
{ name = 'platform', type = 'string'},
{ name = 'carrier', type = 'string'},
{ name = 'hiya', type = 'boolean'},
{ name = 'jobId', type = 'unsigned'},
{ name = 'priority', type = 'string'},
{ name = 'model', type = 'string'},
{ name = 'attestation', type = 'string'},
{ name = 'deviceImageUrl', type = 'string', is_nullable = true},
{ name = 'originationCarrier', type = 'string'},
{ name = 'active', type = 'unsigned'},
})

box.schema.sequence.create('devSSequence',{min=1, start=1, if_not_exists = true})
devS:create_index('id', {sequence='devSSequence', if_not_exists = true})

setBaseIndex('DevStatus')


devS:create_index('serial', {unique = false,type = 'TREE',parts = { 'serial' },if_not_exists = true})
devS:create_index('deviceId', {unique = false,type = 'TREE',parts = { 'deviceId' },if_not_exists = true})
devS:create_index('devCallStatus', {unique = false,type = 'TREE',parts = { 'devCallStatus' },if_not_exists = true})
devS:create_index('incomingNumber', {unique = false,type = 'TREE',parts = { 'incomingNumber' },if_not_exists = true})
devS:create_index('deviceNumber', {unique = false,type = 'TREE',parts = { 'deviceNumber' },if_not_exists = true})
devS:create_index('callSession', {unique = false,type = 'TREE',parts = { 'callSession' },if_not_exists = true})
devS:create_index('platform_carrier', {unique = false,type = 'TREE',parts = { 'platform', 'carrier' },if_not_exists = true})
devS:create_index('jobId', {unique = false,type = 'TREE',parts = { 'jobId' },if_not_exists = true})
