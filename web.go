package common

import "net/http"

const ContentTypeHeader = "Content-Type"
const CacheControlHeader = "Cache-Control"
const ContentTypeJson = "application/json"

type WebFunction func(response http.ResponseWriter, request *http.Request)
