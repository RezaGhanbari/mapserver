package app

import (
	"encoding/json"
	"fmt"
	"github.com/corpix/uarand"
	"github.com/gin-gonic/gin"
	"log"
	"mapserver/cnst"
	"net/http"
	"strings"
)

func googleRoute(c *gin.Context) {
	origin := c.Query("origin")
	destination := c.Query("destination")

	splitOrigin := strings.Split(origin, cnst.LatinComma)
	latOrigin, lngOrigin := splitOrigin[0], splitOrigin[1]

	splitDestination := strings.Split(destination, cnst.LatinComma)
	latDest, lngDest := splitDestination[0], splitDestination[1]
	log.Println(latOrigin, lngOrigin, latDest, lngDest)

	url := cnst.WazeMapURL + fmt.Sprintf(cnst.WazeMapRouteURL, lngOrigin, latOrigin, lngDest, latDest)
	//url := GoogleMapURL + fmt.Sprintf(GoogleMapRouteURL,
	//	latOrigin, lngOrigin, latDest, lngDest, GoogleMapMode, googleKey)
	//
	req, err := http.NewRequest(cnst.GET, url, nil)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err.Error(), c)
		return
	}
	req.Header.Add("dnt", "1")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "en-US,en;q=0.9,fa;q=0.8,la;q=0.7,ar;q=0.6,fr;q=0.5")
	req.Header.Add("user-agent", uarand.GetRandom())
	req.Header.Add("accept", "*/*")
	req.Header.Add("referer", "https://www.waze.com/livemap")
	req.Header.Add("authority", "www.waze.com")


	res, err := http.DefaultClient.Do(req)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err.Error(), c)
		return
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	wazeRoute := new(BaseWaze)
	err = json.NewDecoder(res.Body).Decode(&wazeRoute)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err.Error(), c)
		return
	}
	c.Header(cnst.CONTENT, cnst.ContentType)
	c.JSON(res.StatusCode, wazeRoute)
}
