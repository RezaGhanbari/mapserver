package cnst

const (
	// Google constants
	WazeCacheTime = 60
	WazeMapURL = "https://www.waze.com/"
	WazeMapRouteURL = "row-RoutingManager/routingRequest?at=0&clientVersion=4.0.0&from=x%%3A"+ "%v" +"%%20y%%3A" + "%v" +
		"&nPaths=3&options=AVOID_TRAILS%%3At%%2CALLOW_UTURNS%%3At&returnGeometries=true"+
		"&returnInstructions=true&returnJSON=true&timeout=60000&to=x%%3A"+ "%v" +"%%20y%%3A" + "%v"



	// Common constants
	RequestID              = "X-Request-Id"
	MessageTooManyRequests = "Too many requests, try later."
	PrefixTooManyRequests  = "T-GEO"
	Language               = "fa"
	Google                 = "GOOGLE"
	ACCEPT                 = "accept"
	CONTENT                = "Content-Type"
	XClientID              = "X-Client-Id"
	AcceptHeader           = "application/json"
	ContentType            = "application/json; charset=utf-8"
	BadRequestMessage      = "error, credentials are not provided"
	PingMessage            = "ok"
	RedisError             = "error, could not connect to redis"
	RedisOk                = "Successful!"
	ErrorPrefix            = "ERROR_"
	GoogleMapMode          = "driving"
	Space                  = " "
	SpaceInURL             = "%20"
	SpaceComma             = "، "
	LatinComma             = ","
	OutOfServiceArea       = "محدوده ی خارج از سرویس"
	COLON                  = ":"
	GET                    = "GET"
	SET                    = "SET"
	POST                   = "POST"
	DEL                   = "DEL"
	INCR                   = "INCR"
	CacheTime              = 10
	TLimit                 = 1000000
	XRequestID             = "X-Request-Id"
	XGeoID                 = "X-GEO-Id"
	InvalidAPIToken        = "error, invalid API token"
	APITokenRequired       = "API token required"
	elasticIndexName       = "geo_logger"
	elasticTypeName        = "request_log"
	EMPTY                  = ""
	Origin = "Origin"
	Referer = "Referer"


)

