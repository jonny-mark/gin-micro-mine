package redis_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	redisMine "gin-micro-mine/pkg/redis"
	"context"
)

var _ = Describe("Redis", func() {
	Describe("redis init", func() {
		BeforeEach(func() {
			redisMine.InitTestRedis()
		})

		AfterEach(func() {
			Expect(redisMine.RedisClient.Close()).NotTo(HaveOccurred())
		})

		It("开始测试", func() {
			ctx := context.Background()
			cccVal := redisMine.RedisClient.Get(ctx,"ccc").Val()
			Expect(cccVal).To(Equal(""))
		})
	})
})
