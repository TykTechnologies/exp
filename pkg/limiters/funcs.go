package limiters

import (
	"fmt"
	"reflect"
)

func checkResponseFromRedis(response interface{}, expected interface{}) error {
	if s, sok := response.(string); sok && s == "RACE_CONDITION" {
		return ErrRaceCondition
	}
	if !reflect.DeepEqual(response, expected) {
		return fmt.Errorf("got %+v from redis, expected %+v", response, expected)
	}
	return nil
}
