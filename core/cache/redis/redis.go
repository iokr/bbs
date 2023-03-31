package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	Config struct {
		Network  string // 网络类型, tcp or unix, 默认tcp
		Address  string
		Password string
		Prefix   string
		DB       int // redis数据库index, 默认0

		// 连接池容量及闲置连接数量
		PoolSize     int // 连接池最大socket连接数, 默认为4倍CPU数
		MinIdleCoons int // 在启动阶段创建指定数量的Idle连接, 并长期维持idle状态的连接数不少于指定数量

		// 超时
		DialTimeout  int // 连接建立超时时间, 默认5秒
		ReadTimeout  int // 读超时, 默认3秒, -1表示取消读超时
		WriteTimeout int // 写超时, 默认等于读超时
		PoolTimeout  int // 当所有连接都处在繁忙状态时, 客户端等待可用连接的最大等待时长, 默认为读超时+1秒

		// 闲置连接检查包括IdleTimeout, MaxConnAge
		IdleTimeOut int
	}

	Client struct {
		prefix string
		client *redis.Client
	}
)

func NewRedis(c *Config) *Client {
	redisClient := redis.NewClient(&redis.Options{
		Network:  c.Network,
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,

		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleCoons,

		DialTimeout:  time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
		PoolTimeout:  time.Duration(c.PoolTimeout) * time.Second,

		IdleTimeout: time.Duration(c.IdleTimeOut) * time.Second,
	})

	err := redisClient.Ping().Err()
	if err != nil {
		panic(err)
	}

	return &Client{
		prefix: c.Prefix,
		client: redisClient,
	}
}

func (r *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.prefix+key, value, expiration).Err()
}

func (r *Client) GetString(key string) (string, error) {
	return r.client.Get(r.prefix + key).Result()
}
