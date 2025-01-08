package main

import (
	proto "GopherMall/inventory_srv/proto/.InventoryProto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"testing"
)

func TestNegativeLock(t *testing.T) {
	c, err := grpc.NewClient("127.0.0.1:50502", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
	client := proto.NewInventoryClient(c)

	ctx := context.Background()

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(w *sync.WaitGroup) {
			_, _ = client.Sell(ctx, &proto.SellInfo{GoodsInfo: []*proto.GoodsInvInfo{
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
