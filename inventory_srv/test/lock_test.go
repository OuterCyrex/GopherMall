package test

import (
	proto "GopherMall/inventory_srv/proto/.InventoryProto"
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"sync"
	"testing"
)

func TestNegativeLock(t *testing.T) {
	var count = 0

	c, err := grpc.NewClient("127.0.0.1:64952", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
	client := proto.NewInventoryClient(c)

	ctx := context.Background()

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(w *sync.WaitGroup) {
			_, err := client.Sell(ctx, &proto.SellInfo{GoodsInfo: []*proto.GoodsInvInfo{
				{
					GoodsId: int32(426),
					Num:     int32(1),
				}, {
					GoodsId: int32(424),
					Num:     int32(1),
				},
			}})
			if err != nil {
				count += 1
			}
			w.Done()
		}(&wg)
	}

	wg.Wait()

	t.Logf("出错次数：%v", count)
}

func TestOptimisticLock(t *testing.T) {
	c, err := grpc.NewClient("127.0.0.1:55753", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
	client := proto.NewInventoryClient(c)

	ctx := context.Background()

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(w *sync.WaitGroup) {
			_, _ = client.Reback(ctx, &proto.SellInfo{GoodsInfo: []*proto.GoodsInvInfo{
				{
					GoodsId: int32(426),
					Num:     int32(1),
				},
			}})
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("127.0.0.1:6379"),
		DB:   0,
	})

	if rdb.Ping(context.Background()).Err() != nil {
		zap.S().Panicw("redis init failed", "err", "redis init failed")
	}

	rs := redsync.New(goredis.NewPool(rdb))

	mutex := rs.NewMutex("test")
	if err := mutex.Lock(); err != nil {
		t.Error(err)
	}

	if ok, err := mutex.Unlock(); !ok || err != nil {
		t.Error(err)
	}
}
