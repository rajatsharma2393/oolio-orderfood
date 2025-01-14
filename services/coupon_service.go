package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type CouponService struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCouponService() *CouponService {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})
	return &CouponService{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (cs *CouponService) IsValidCoupon(coupon string) bool {
	if len(coupon) < 8 || len(coupon) > 10 {
		return false
	}
	redisKey := coupon
	fileNames, err := cs.rdb.SMembers(cs.ctx, redisKey).Result()
	if err != nil {
		log.Printf("Error checking coupon %s: %v\n", coupon, err)
		return false
	}

	if len(fileNames) >= 2 {
		return true
	}

	return false
}
