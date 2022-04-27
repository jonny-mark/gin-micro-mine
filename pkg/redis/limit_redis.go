/**
 * @author jiangshangfang
 * @date 2022/2/12 11:03 PM
 **/
package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	// lua脚本实现令牌桶算法限流
	ScriptTokenLimit = `
local rateLimit = redis.pcall('HMGET',KEYS[1],'lastTime','tokens')
local lastTime = rateLimit[1]
local tokens = tonumber(rateLimit[2])
local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
if tokens == nil then
  tokens = capacity
else
  local deltaTokens = math.floor((now-lastTime)*rate)
  tokens = tokens+deltaTokens
  if tokens>capacity then
    tokens = capacity
  end
end
local result = false
lastTime = now
if(tokens>0) then
  result = true
  tokens = tokens-1
end
redis.call('HMSET',KEYS[1],'lastTime',lastTime,'tokens',tokens)
return result
`
)

//得到cache的名字
func getIdCacheName(idType string) string {
	return "LuaTokenBucket_" + idType
}

// LuaTokenBucket 通过lua脚本实现令牌桶算法限流
func LuaTokenBucket(key string, capacity, rate int64) (bool, error) {
	luaScript := redis.NewScript(ScriptTokenLimit)
	var ctx = context.Background()
	now := time.Now().Unix()

	_, err := luaScript.Run(ctx, RedisClient, []string{key}, capacity, rate, now).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
