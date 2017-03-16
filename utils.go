package arango

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"github.com/apex/log"
)

func buildHostAddress(host string, tls bool) string {
	if tls {
		return fmt.Sprintf("https://%s", host)
	}
	return fmt.Sprintf("http://%s", host)
}

func debugHttpReqResp(req *http.Request, resp *http.Response) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.WithError(err).Fatal("error while loging request")
	}
	log.Infof("--- REQUEST START ---\n%s\n--- REQUEST END ---", reqDump)
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.WithError(err).Fatal("error while loging response")
	}
	log.Infof("--- RESPONSE START ---\n%s\n--- RESPONSE END ---", respDump)
}
