package view

import (
	"net/http"
	"strconv"
	"strings"
)

func newRequest(httpRequest *http.Request) *Request {
	return &Request{
		Request: httpRequest,
	}
}

type Request struct {
	*http.Request
	Params map[string]string
}

// AddProtocolAndHostToURL adds the protocol (http:// or https://)
// and request host (domain or IP) to an URL if not present.
func (request *Request) AddProtocolAndHostToURL(url string) string {
	if len(url) > 0 && url[0] == '/' {
		if request.TLS != nil {
			url = "https://" + request.Host + url
		} else {
			url = "http://" + request.Host + url
		}
	}
	return url
}

// URL returns the complete URL of the request including protocol and host.
func (request *Request) URLString() string {
	return request.AddProtocolAndHostToURL(request.RequestURI)
}

// // todo: all browsers
// func (request *Request) ParseUserAgent() (renderer string, version utils.VersionTuple, err error) {
// 	s := request.UserAgent()
// 	switch {
// 	case strings.Contains(s, "Gecko"):
// 		if i := strings.Index(s, "rv:"); i != -1 {
// 			i += len("rv:")
// 			if l := strings.IndexAny(s[i:], "); "); l != -1 {
// 				if version, err = utils.ParseVersionTuple(s[i : i+l]); err == nil {
// 					return "Gecko", version, nil
// 				}
// 			}
// 		}
// 	case strings.Contains(s, "MSIE "):
// 		if i := strings.Index(s, "MSIE "); i != -1 {
// 			i += len("MSIE ")
// 			if l := strings.IndexAny(s[i:], "); "); l != -1 {
// 				if version, err = utils.ParseVersionTuple(s[i : i+l]); err == nil {
// 					return "MSIE", version, nil
// 				}
// 			}
// 		}
// 	}
// 	return "", nil, nil
// }

func (request *Request) Port() uint16 {
	i := strings.LastIndex(request.Host, ":")
	if i == -1 {
		return 80
	}
	port, _ := strconv.ParseInt(request.Host[i+1:], 10, 16)
	return uint16(port)
}

func (request *Request) GetSecureCookie(name string) (string, bool) {
	panic("not implemented")
}
