package app

import (
	"encoding/json"
	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"mapserver/cnst"
	"net"
	"net/http"
	"regexp"
	"time"
)

func respondWithError(code int, s string, c *gin.Context) {
	resp := map[string]string{"error": s}
	c.JSON(code, resp)
	c.Abort()
}

// CheckClientID function
func CheckClientID() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Request.Header.Get(cnst.XClientID)
		if clientID == cnst.EMPTY {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}
		match, _ := regexp.MatchString("^[^:]*:{1}[^:]*$", clientID)
		if match == false {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}
	}
}

func routeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.DefaultQuery("origin", cnst.EMPTY)
		destination := c.DefaultQuery("destination", cnst.EMPTY)
		if origin == cnst.EMPTY || destination == cnst.EMPTY {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}

		matchOrigin, _ := regexp.MatchString("^[^,][0-9]*.{1}?[.][0-9]*$", origin)
		if matchOrigin == false {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}

		matchDestination, _ := regexp.MatchString("^[^,][0-9]*.{1}?[.][0-9]*$", destination)
		if matchDestination == false {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}
	}
}

func latLanMiddleware(q1, q2 string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := c.DefaultQuery(q1, cnst.EMPTY)
		end := c.DefaultQuery(q2, cnst.EMPTY)
		if start == cnst.EMPTY || end == cnst.EMPTY {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}

		matchStart, _ := regexp.MatchString("^[0-9].{1}?[.][0-9]*,[0-9].{1}?[.][0-9]*$", start)
		if matchStart == false {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}

		matchEnd, _ := regexp.MatchString("^[0-9].{1}?[.][0-9]*,[0-9].{1}?[.][0-9]*$", end)
		if matchEnd == false {
			respondWithError(http.StatusBadRequest, cnst.BadRequestMessage, c)
			return
		}
	}
}

// TokenAuthMiddleware function
func TokenAuthMiddleware(config *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// read from header
		token := c.Request.Header.Get("api-token")

		if token == cnst.EMPTY {
			respondWithError(http.StatusUnauthorized, cnst.APITokenRequired, c)
			return
		}
		if token != config.APIToken {
			respondWithError(http.StatusUnauthorized, cnst.APITokenRequired, c)
			return
		}
		c.Next()
	}
}

// RequestIDMiddleware function
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := uuid.NewV4()
		c.Writer.Header().Set(cnst.XRequestID, u.String())
		c.Next()
	}
}

// RequestLogger function middleware
// todo

// Throttling middleware
func Throttling() gin.HandlerFunc {
	internalErrorResp := MessageError{Body: cnst.MessageTooManyRequests}
	internalOutError, err := json.Marshal(internalErrorResp)
	if err != nil {
		log.Printf("error %s", err.Error())
	}
	return func(c *gin.Context) {
		throttle.Policy(&throttle.Quota{
			Limit:  cnst.TLimit,
			Within: time.Minute,
		}, &throttle.Options{
			StatusCode: http.StatusTooManyRequests,
			Message:    string(internalOutError),

			// this function represents the throttling algorithm
			IdentificationFunction: func(req *http.Request) string {
				if requestId := req.Header.Get(cnst.RequestID); requestId != cnst.EMPTY {
					return requestId
				}
				ip, _, err := net.SplitHostPort(req.RemoteAddr)
				if err != nil {
					panic(err.Error())
				}
				return ip
			},
			KeyPrefix: cnst.PrefixTooManyRequests,
			Disabled:  false,
		})
		c.Next()
	}
}
