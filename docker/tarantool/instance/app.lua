box.cfg {
    listen = '*:3301';
    io_collect_interval = nil;
    readahead = 16320;
    memtx_memory = 1024 * 1024 * 1024; -- 1G
    memtx_min_tuple_size = 16;
    memtx_max_tuple_size = 128 * 1024 * 1024; -- 1G
    wal_mode = "none";
    wal_max_size = 256 * 1024 * 1024;
    checkpoint_interval = 60 * 60; -- one hour
    checkpoint_count = 6;
    force_recovery = true;
    log_level = 5;
    log_nonblock = false;
    too_long_threshold = 0.5;
}

box.schema.user.create('administrator', {password = 'SuperSecurePassword', if_not_exists = true})
box.schema.user.grant('administrator', 'alter,read,write,execute,create,drop', 'universe', nil, {if_not_exists=true})


function setBaseIndex (space)
    box.space[space]:create_index('id', {
        type = 'TREE',
        parts = {'id'},
        if_not_exists = true
    })

    box.space[space]:create_index('status', {
        type = 'TREE',
        unique = false,
        parts = {'status'},
        if_not_exists = true,
    })

    box.space[space]:create_index('ttl', {
        type = 'TREE',
        unique = false,
        parts = {'ttl'},
        if_not_exists = true,
   })

    box.space[space]:create_index('attempts', {
        type = 'TREE',
        unique = false,
        parts = {'attempts'},
        if_not_exists = true,
    })
end

function getCount(space)
    return box.space[space]:len()
end

function getByStatus(space, status)
    return box.space[space]:count(status, {iterator='GE'})
end

function getRowById(space)
    return box.space[space]:get{id}
end

function getResults(lim, offs)
    return box.space.DCCache.index.status:select({'r'}, {iterator='REQ', limit=lim, offset=offs})
end

require('DevStat')
require('Fifo')
require('Calls')
require('CallsQueue')
require('Logs')
