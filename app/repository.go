package app

import (
	"bytes"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type App struct {
	redisPool *redis.Pool
	router    *gin.Engine
	Config    *Config
	Store     *persistence.RedisStore
}

// MessageError struct
type MessageError struct {
	Body string `json:"body"`
}


type BodyLogWriter struct {
	ResponseWriter gin.ResponseWriter
	Body *bytes.Buffer
}

type Config struct {
	MapName         string `envconfig:"MAP_NAME"`
	ReversMapName   string `envconfig:"REVERSE_MAP_NAME"`
	SearchMapName   string `envconfig:"SEARCH_MAP_NAME"`
	RouteMapName    string `envconfig:"ROUTE_MAP_NAME"`
	DistanceMapName string `envconfig:"DISTANCE_MAP_NAME"`
	DurationMapName string `envconfig:"DURATION_MAP_NAME"`
	Server          string `envconfig:"SERVER"`
	Port            string `envconfig:"PORT"`
	APIToken        string `envconfig:"API_TOKEN"`
	DB              string `envconfig:"DB"`
	RedisHost       string `envconfig:"REDIS_HOST"`
	RedisPort       string `envconfig:"REDIS_PORT"`
}

type Path struct {
	SegmentId int64   `json:"segmentId"`
	NodeId    int64   `json:"nodeId"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction bool    `json:"direction"`
}

type Instruction struct {
	Opcode          string `json:"opcode"`
	Arg             int    `json:"arg"`
	InstructionText string `json:"instructionText"`
	LaneGuidance    string `json:"laneGuidance"`
	Name            string `json:"name"`
	Tts             string `json:"tts"`
}

type Result struct {
	Path                     Path        `json:"path"`
	Street                   int         `json:"street"`
	AltStreets               string      `json:"altStreets"`
	Distance                 int         `json:"distance"`
	Length                   int         `json:"length"`
	CrossTime                int         `json:"crossTime"`
	CrossTimeWithoutRealTime int         `json:"crossTimeWithoutRealTime"`
	Tiles                    string      `json:"tiles"`
	ClientIds                string      `json:"clientIds"`
	KnownDirection           bool        `json:"knownDirection"`
	Penalty                  int         `json:"penalty"`
	RoadType                 int         `json:"roadType"`
	IsToll                   bool        `json:"isToll"`
	NaiveRoute               string      `json:"naiveRoute"`
	DetourSavings            int         `json:"detourSavings"`
	DetourSavingsNoRT        int         `json:"detourSavingsNoRT"`
	UseHovLane               bool        `json:"useHovLane"`
	Attributes               int         `json:"attributes"`
	Lane                     string      `json:"lane"`
	LaneType                 string      `json:"laneType"`
	Areas                    []string    `json:"areas"`
	RequiredPermits          []string    `json:"requiredPermits"`
	DetourRoute              string      `json:"detourRoute"`
	NaiveRouteFullResult     string      `json:"naiveRouteFullResult"`
	DetourRouteFullResult    string      `json:"detourRouteFullResult"`
	MergeOffset              int         `json:"mergeOffset"`
	AvoidStatus              string      `json:"avoidStatus"`
	LaneDefinition           string      `json:"laneDefinition"`
	AdditionalInstruction    string      `json:"additionalInstruction"`
	Instruction              Instruction `json:"instruction"`
}

type Response struct {
	Results                 []Result `json:"results"`
	StreetNames             []string `json:"streetNames"`
	TileIds                 []string `json:"tileIds"`
	TileUpdateTimes         []string `json:"tileUpdateTimes"`
	Geom                    string   `json:"geom"`
	FromFraction            float64  `json:"fromFraction"`
	ToFraction              float64  `json:"toFraction"`
	SameFromSegment         bool     `json:"sameFromSegment"`
	SameToSegment           bool     `json:"sameToSegment"`
	AstarPoints             string   `json:"astarPoints"`
	WayPointIndexes         string   `json:"wayPointIndexes"`
	WayPointFractions       string   `json:"wayPointFractions"`
	TollMeters              int      `json:"tollMeters"`
	PreferedRouteId         int      `json:"preferedRouteId"`
	IsInvalid               bool     `json:"isInvalid"`
	IsBlocked               bool     `json:"isBlocked"`
	ServerUniqueId          string   `json:"serverUniqueId"`
	DisplayRoute            bool     `json:"displayRoute"`
	AstarVisited            int      `json:"astarVisited"`
	AstarResult             string   `json:"astarResult"`
	AstarData               string   `json:"astarData"`
	IsRestricted            bool     `json:"isRestricted"`
	AvoidStatus             string   `json:"avoidStatus"`
	DueToOverride           string   `json:"dueToOverride"`
	PassesThroughDangerArea bool     `json:"passesThroughDangerArea"`
	DistanceFromSource      int      `json:"distanceFromSource"`
	DistanceFromTarget      int      `json:"distanceFromTarget"`
	MinPassengers           int      `json:"minPassengers"`
	HovIndex                int      `json:"hovIndex"`
	TimeZone                string   `json:"timeZone"`
	RouteType               []string `json:"routeType"`
	RouteAttr               []string `json:"routeAttr"`
	AstarCost               int      `json:"astarCost"`
	ReorderChoice           string   `json:"reorderChoice"`
	TotalRouteTime          int      `json:"totalRouteTime"`
	LaneTypes               []string `json:"laneTypes"`
	PreferredStoppingPoints string   `json:"preferredStoppingPoints"`
	Areas                   []string `json:"areas"`
	RequiredPermits         []string `json:"requiredPermits"`
	EtaHistograms           []string `json:"etaHistograms"`
	SegGeoms                string   `json:"segGeoms"`
	RouteName               string   `json:"routeName"`
	RouteNameStreetIds      []int    `json:"routeNameStreetIds"`
	Open                    bool     `json:"open"`
}

type Coord struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z string  `json:"z"`
}

type Alternative struct {
	Response Response `json:"response"`
	Coords   []Coord  `json:"coords"`
	SegCords string   `json:"segCoords"`
}

type BaseWaze struct {
	Alternatives []Alternative `json:"alternatives"`
}

type RouteRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
