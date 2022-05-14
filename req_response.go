package golang_helpers

import (
	"net"
	"net/http"
	"strings"
)

func GetReferer(r *http.Request) string {
	return r.Header.Get("Referer")
}

func GetAuthToken(r *http.Request) string {
	token := ""
	tokenHeader := r.Header.Get("Authorization")
	if len(tokenHeader) > 0 {
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) == 2 {
			token = splitted[1]
		}
	}
	return token
}

func GetUserIP(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if len(ip) == 0 {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	if strings.Contains(ip, ",") {
		ipval := strings.Split(ip, ",")
		ip = ipval[1]
	}

	return ip
}
