local redis        = require "resty/redis"

local log          = ngx.log
local ERR          = ngx.ERR
local setmetatable = setmetatable

local _M           = {

}

local mt           = { __index = _M }

local function errlog(...)
  log(ERR, "Redis: ", ...)
end

function _M.exec(self, func)

  local red = redis:new()
  red:set_timeout(self.timeout)

  local ok, err = red:connect(self.host, self.port)
  if not ok then
    errlog("Cannot connect, host: " .. self.host .. ", port: " .. self.port)
    return nil, err
  end

  if self.password ~= '' then
    local count
    count, err = red:get_reused_times()
    if 0 == count then
      ok, err = red:auth(self.password)
      if not ok then
        ngx.say("failed to auth: ", err)
        return
      end
    elseif err then
      ngx.say("failed to get reused times: ", err)
      return
    end
  end
  red:select(self.database)

  local res, err = func(red)
  if res then
    local ok, err = red:set_keepalive(self.max_idle_time, self.pool_size)
    if not ok then
      red:close()
    end
  end
  return res, err
end

function _M.new(opts)
  local config = opts or {}


  local self   = {
    host          = config.host or "172.20.0.4",
    password      = config.password or '',
    port          = config.port or 6379,
    timeout       = config.timeout or 5000,
    database      = config.database or 0,
    max_idle_time = config.max_idle_time or 60000,
    pool_size     = config.pool_size or 100
  }

  return setmetatable(self, mt)
end

return _M
