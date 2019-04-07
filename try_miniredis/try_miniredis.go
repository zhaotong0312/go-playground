package try_miniredis

import (
	utilsRedis "bitbucket.paloaltonetworks.local/tzhao/compare-rtdns/utils/redis"
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

var testTLDList = []string{"com", "org", "tw", "ve", "in", "xyz", "gq", "com.tw", "org.ve", "gov.in", "tmall.com"}
var testWhitelistDbList = []string{"google.com", "youtube.com", "facebook.com", "baidu.com", "wikipedia.org", "qq.com", "taobao.com", "tmall.com", "yahoo.com", "amazon.com", "twitter.com", "jd.com", "login.tmall.com", "ltn.com.tw", "patria.org.ve", "indianrail.gov.in"}
var testDGADbList = []string{"tqqtsbhuws.ml", "vltljygrqr.club", "0yj0aj2tolls.gq", "gvchqvfopr.tk", "ojtksogefqmzlzdeisgkndebuthu.biz", "qfkssdluj.com", "1swghwuqbe.biz", "lqbsxrfgfoxx.cf", "hsqsnhdzx.com", "pnobzthtfamxibrgkbmrifhey.biz", "mxujwuvhxdkda.pw"}
var testMalDbList = []map[string]string{
	{"i": "125862478", "h": "oionih.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "107911799", "h": "vayevu.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "131576612", "h": "auiumg.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "23602961", "h": "ieaeoe.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "133638086", "h": "qiajbi.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "65058433", "h": "woncxik.eu", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "50858328", "h": "xiwoxa.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "107890434", "h": "yebunm.com", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "194360952", "h": "ypnklmbknkhp.info", "v:mal": "1", "ttl": "300", "v:c2": "1"},
	{"i": "136266920", "h": "bid48.com", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "127895210", "h": "ebptdliw.us", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "210043080", "h": "pm2pcr.667443.win", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "68613017", "h": "sdzebn.com", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "208339776", "h": "vzx0rr.qceg6.lhc48b5.win", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": " 119910296", "h": "wcehon.com", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "69870899", "h": "yzupssejhci.kopeoxbbret.com", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "136365070", "h": "mirloadplus.ru", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "194762589", "h": "pubg157.ddns.net", "v:mal": "1", "ttl": "300", "v:c2": "0"},
	{"i": "119460440", "h": "thojrmyyr.com", "v:mal": "1", "ttl": "300", "v:c2": "0"},
}

var redisDB = 1
var wlRedisDB = 2
var dgaWLRedisDB = 3

func NewMiniredis() (r *miniredis.Miniredis) {
	r, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	r.Select(redisDB)
	for _, item := range testMalDbList {
		host, in := item["h"]
		if !in {
			panic("map does not have \"h\" field.")
		}
		for k, v := range item {
			r.HSet(host, k, v)
		}
	}
	r.Select(wlRedisDB)
	for _, item := range testWhitelistDbList {
		r.HSet(item, "ut", "190328")
		r.HSet(item, "h", item)
	}
	r.Select(dgaWLRedisDB)
	for _, item := range testDGADbList {
		r.HSet(item, "h", item)
	}
	return r
}
func main() {
	redisMock := NewMiniredis()
	fmt.Println("redisMock created")
	redisCli := utilsRedis.NewRedis(nil, func(source interface{}) utilsRedis.Config {
		return utilsRedis.Config{
			Name: "test redis",
			Ip:   redisMock.Host(),
			Port: redisMock.Port(),
		}
	})

	redisConn := redisCli.Get()
	resp, _ := redisConn.Do("PING")
	fmt.Println(resp)

	redisConn.Send("SELECT", wlRedisDB)
	for _, item := range testWhitelistDbList {
		redisConn.Send("HGETALL", item)
	}
	redisConn.Flush()

	//for _, item := range testMalDbList {
	//	host, in := item["h"]
	//	if !in {
	//		panic("map does not have \"h\" field.")
	//	}
	//	redisConn.Send("HGETALL", host)
	//}

	for _, item := range testWhitelistDbList {
		sm, _ := redis.StringMap(redisConn.Receive())
		fmt.Println(item, sm)
	}
}
