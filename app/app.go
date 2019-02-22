package app

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
	"mapserver/cnst"
	"mapserver/redis"
	"time"
)

func (a *App) Initialize() {
	// TODO get config from config-manager
	var config Config
	err := envconfig.Process(cnst.EMPTY, &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	gin.SetMode(gin.DebugMode)
	r := gin.New()

	redisPool := redis.Init(config.RedisHost, config.RedisPort)
	store := persistence.NewRedisCacheWithPool(redisPool, time.Minute)

	a.Config = &config
	a.router = r
	a.redisPool = redisPool
	a.Store = store
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	r := a.router
	r.Use(gin.Recovery())
	ApiV1 := r.Group("/v1")

	internal := ApiV1.Group("/i")
	internal.Use(RequestIDMiddleware())
	internal.GET("/healthcheck", HealthCheck)


	authorized := ApiV1.Group("/a")
	authorized.Use(RequestIDMiddleware())
	authorized.GET("/route", latLanMiddleware("origin", "destination"),
		cache.CachePage(a.Store, cnst.WazeCacheTime*time.Second, googleRoute))

}

func (a *App) Run(addr string) {
	log.Fatal(a.router.Run(cnst.COLON + addr))
}
