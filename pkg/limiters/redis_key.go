package limiters

import (
	"fmt"
)

func redisKey(prefix, key string) string {
	return fmt.Sprintf("%s/%s", prefix, key)
}
