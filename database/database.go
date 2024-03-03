package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NY-Daystar/populate-redis/configuration"
	"github.com/NY-Daystar/populate-redis/utils"
	"go.uber.org/zap"

	"gopkg.in/redis.v3"
)

// NewRedisClient Create redis client connection with host (endpoint and port)
func NewRedisClient(cfg configuration.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	return client
}

// AddUser set key to redis client with specific data
func AddUser(client *redis.Client, logger *zap.Logger, verbose bool) {
	user := utils.GenerateName()

	if verbose {
		logger.Sugar().Debugf("Generating user `%v`", user)
	}

	parts := strings.Split(user, "-")

	firstname := parts[0]
	surname := parts[1]
	datetime := time.Now()

	key := fmt.Sprintf("users:%s", user)

	client.HSet(key, "firstname", firstname)
	client.HSet(key, "surname", surname)
	client.HSet(key, "createdAt", datetime.Format("2006-01-02 15:04:05"))
	client.HSet(key, "updateAt", datetime.Format("2006-01-02 15:04:05"))

	key = fmt.Sprintf("address:%s", user)

	client.HSet(key, "city", "New-York")
	client.HSet(key, "country", "United-States")
	client.HSet(key, "street", "Broadway")
	client.HSet(key, "number", "1")

	client.ZAdd("users", []redis.Z{{Score: float64(datetime.Unix()), Member: user}}...)
}

// FlushDatabase : Delete all data in the redis database
func FlushDatabase(client *redis.Client) {
	client.FlushAll()
}
