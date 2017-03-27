package arangodb

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"github.com/thedanielforum/arangodb/types"
	"github.com/apex/log"
)

// post sends a post request to the api.
// endpoint should not start with a slash.
func (c *Connection) post(endpoint string, body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)

	// Prepare request
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.host, endpoint),
		buf,
	)
	req.Header = c.header
	if err != nil {
		return nil, err
	}

	// Execute
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// UnMarshal response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check response code for error messages.
	if resp.StatusCode > 203 {
		err = types.NewDbError(respBody).ToError()
		log.WithError(err).Warn("error response from db api")
		return nil, err
	}

	return respBody, nil
}

func (c *Connection) get(endpoint string) ([]byte, error) {
	// Prepare request
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", c.host, endpoint),
		nil,
	)
	req.Header = c.header
	if err != nil {
		return nil, err
	}

	// Execute
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.config.DebugMode {
		debugHttpReqResp(req, resp)
	}

	// UnMarshal response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check response code for error messages.
	if resp.StatusCode > 203 {
		err = types.NewDbError(respBody).ToError()
		log.WithError(err).Warn("error response from db api")
		return nil, err
	}

	return respBody, nil
}
