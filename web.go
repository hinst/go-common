package common

import (
	"net/http"
	"strconv"
	"time"
)

const ContentTypeHeader = "Content-Type"
const CacheControlHeader = "Cache-Control"
const ContentTypeJson = "application/json"

type WebFunction func(response http.ResponseWriter, request *http.Request)

func SetCacheAge(response http.ResponseWriter, duration time.Duration) {
	response.Header().Set(CacheControlHeader, "max-age="+strconv.Itoa(int(duration.Seconds())))
}
