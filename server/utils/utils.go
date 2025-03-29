package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomDuration(min, max int) time.Duration {
	if min > max || min < 1 {
		panic("Invalid range of time")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomSec := r.Intn(max-min+1) + min

	return time.Duration(randomSec) * time.Second
}
