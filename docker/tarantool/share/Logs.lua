lg = box.schema.space.create('Logs', { if_not_exists =true })

lg:format({
    { name = 'id', type = 'integer' },
    { name = 'status', type = 'string' },
    { name = 'ttl', type = 'integer', is_nullable = true},
    { name = 'attempts', type = 'integer', is_nullable = true},

    { name = 'case', type = 'string'},
    { name = 'event', type = 'string'},
    { name = 'extra', type = 'string'},
})


box.schema.sequence.create('lgCacheSequence',{min=1, start=1, if_not_exists = true})
lg:create_index('id', {sequence='lgCacheSequence', if_not_exists = true})

setBaseIndex('Logs')

lg:create_index('case', {
    unique = false,
    type = 'TREE',
    parts = { 'case' },
    if_not_exists = true,
})

lg:create_index('event', {
    unique = false,
    type = 'TREE',
    parts = { 'event' },
    if_not_exists = true,
})
