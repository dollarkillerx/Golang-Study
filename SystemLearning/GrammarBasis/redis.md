Redis
===
### 数据结构
- string (max 512M)
- hash 哈希表
- list 列表
- set 集合 (无序,不重复)
- zset 有序集合


#### string
``` 
set name dollarkiller
get name

setex key 过期时间s value
ttl key //查询过期时间 -1 永不过期
persist key //取消过期时间
expire key time//重设过期时间


mset key1 value1 key2 value2 //一次性设置多个值

mget key1 key2 //一次性获得多个值

incr/decr age //加or减 1

incrby/decrby age 5 //加or减 5 

append key value //追加

strlen key //获得value长度

setnx key value //当不存在的时候才设置 重要 (返回1 or 0)

del key [..key] //删除
```
- `keys *` 所有的
- `keys a*` 以a开头
- `keys *a` 以a结尾
- `exists key` 是否存在
- `type key` 查询key值的类型
####  hash
``` 
hset obj key value
hget obj key

hsetnx obj key val 不会重复添加

hgetall daddy 获取daddy所有
hkeys daddy 获取daddy所有的属性 

hexists obj key 

hdel obj key
```

#### list 队列有顺序可以重复
``` 
lpush key value left追加 
rpush key value right 追加

rpop key 
lpop key

llen key

lindex key 下标 //获取指定下标的数

lrange key 0 3 //从0到3
lrange key 0 -1 //从头到尾

blpop key timeOut //阻塞一直读
```
#### set集合不重复
``` 
sadd key value [...value]
scard key 查询集合元素个数

smembers key 查询set所有元素

sinter key1 key2 交集合
sdiff key1 key2 差集合
sunion key1 key2 并集合

sismember key value value是否是key里面的元素
```

#### zset 有序集合
``` 
zadd set 权重 value [...权重 value]
zrange set 0 -1
zcard set 个数
zcount set 6 10 权重6-10之间元素个数
zscore set key 返回key在set中的权重
```

#### 小结
``` 
select 数据库编号 选择数据库
flushdb 清空当前数据库
flushall 清空所有数据库
```
