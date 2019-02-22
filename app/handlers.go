package app

import (
	"github.com/gin-gonic/gin"
	"mapserver/cnst"
	"math"
	"net/http"
	"regexp"
	"strconv"
)

//TODO LogReporter function

// check latitude and longitude, with regex
func CheckLatitudeLongitude(item string) (bool, error) {
	itudeMatch, err := regexp.MatchString("^[^.][0-9]*.{1}[^.][0-9]*$", item)
	if err != nil {
		return false, err
	}
	if itudeMatch == true {
		return true, nil
	}
	return false, nil
}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + cnst.Space + singular + cnst.Space
	} else {
		result = strconv.Itoa(count) + cnst.Space + singular + "s "
	}
	return
}

func SecondsToHuman(input int) (result string) {
	years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
	seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
	months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
	seconds = input % (60 * 60 * 24 * 7 * 30)
	weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
	seconds = input % (60 * 60 * 24 * 7)
	days := math.Floor(float64(seconds) / 60 / 60 / 24)
	seconds = input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	if years > 0 {
		result = plural(int(years), "year") + plural(int(months), "month") +
			plural(int(weeks), "week") + plural(int(days), "day") +
			plural(int(hours), "hour") + plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else if months > 0 {
		result = plural(int(months), "month") +
			plural(int(weeks), "week") +
			plural(int(days), "day") +
			plural(int(hours), "hour") +
			plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else if weeks > 0 {
		result = plural(int(weeks), "week") +
			plural(int(days), "day") +
			plural(int(hours), "hour") +
			plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else if days > 0 {
		result = plural(int(days), "day") +
			plural(int(hours), "hour") +
			plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else if hours > 0 {
		result = plural(int(hours), "hour") +
			plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else if minutes > 0 {
		result = plural(int(minutes), "minute") +
			plural(int(seconds), "second")
	} else {
		result = plural(int(seconds), "second")
	}
	return
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": cnst.PingMessage,
	})
}


