package utils

import (
	"math/rand"
	"time"
)

// Init 初始化随机数生成器
func Init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomDelay 产生随机延迟
func RandomDelay(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}
