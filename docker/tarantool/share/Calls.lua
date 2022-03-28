dc = box.schema.space.create('DCCache', { if_not_exists =true })

dc:format({
{ name = 'id', type = 'integer' },
{ name = 'status', type = 'string' },
{ name = 'ttl', type = 'integer', is_nullable = true},
{ name = 'attempts', type = 'integer', is_nullable = true},

{ name = 'log', type = 'string'},
{ name = 'iso_call_end', type = 'string'},
{ name = 'originatingCarrier', type = 'string'},
{ name = 'screenshot', type = 'string'},
{ name = 'incoming_number', type = 'string'},
{ name = 'deviceFailed', type = 'boolean'},
{ name = 'incoming_call_status', type = 'boolean'},
{ name = 'from_num', type = 'string'},
{ name = 'to_num', type = 'string'},
{ name = 'call_end', type = 'string'},
{ name = 'makeCallStatusCode', type = 'unsigned'},
{ name = 'device_id', type = 'unsigned'},
{ name = 'MessageId', type = 'string'},
{ name = 'call_start', type = 'string'},
{ name = 'call_duration', type = 'unsigned'},
{ name = 'incoming_number_match', type = 'boolean'},
{ name = 'iso_call_start', type = 'string'},
{ name = 'cnam', type = 'string'},
{ name = 'text', type = 'string'},
{ name = 'text_recognized', type = 'boolean'},
{ name = 'custom', type = 'string', is_nullable = true},
})


box.schema.sequence.create('dcCacheSequence',{min=1, start=1, if_not_exists = true})
dc:create_index('id', {sequence='dcCacheSequence', if_not_exists = true})

setBaseIndex('DCCache')

dc:create_index('mid', {
unique = false,
type = 'TREE',
parts = { 'MessageId' },
if_not_exists = true,
})
