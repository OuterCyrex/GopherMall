package initialize

import (
	"GopherMall/goods_srv/global"
	"GopherMall/goods_srv/model"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"strconv"
)

func InitElastic() {
	logger := log.New(os.Stdout, "[elastic]", log.LstdFlags)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("https://%s:%d",
			global.ServerConfig.Elastic.Host,
			global.ServerConfig.Elastic.Port,
		)),
		elastic.SetSniff(false),
		elastic.SetHttpClient(&http.Client{Transport: tr}),
		elastic.SetBasicAuth(
			global.ServerConfig.Elastic.UserName,
			global.ServerConfig.Elastic.Password,
		),
		elastic.SetInfoLog(logger),
	)

	if err != nil {
		zap.S().Panicf("Init ElasticSearch Failed: %v", err)
		return
	}

	ctx := context.Background()

	exist, err := client.IndexExists(model.EsGoods{}.GetIndexName()).Do(ctx)
	if err != nil {
		zap.S().Panicf("Init ElasticSearch Failed: %v", err)
		return
	}

	if !exist {
		_, err = client.CreateIndex(model.EsGoods{}.GetIndexName()).
			BodyString(model.EsGoods{}.GetMapping()).Do(ctx)
		if err != nil {
			zap.S().Panicf("Init ElasticSearch Failed: %v", err)
			return
		}
	}

	global.Elastic = client
}

func SyncElasticWithDB() {
	var goods []model.Goods

	global.DB.Find(&goods)

	for _, g := range goods {
		eg := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarkPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err := global.Elastic.Index().Index(eg.GetIndexName()).BodyJson(eg).Id(strconv.Itoa(int(eg.ID))).Do(context.Background())
		if err != nil {
			zap.S().Errorf("Transfer Failed: %v", err)
		}
	}
}
