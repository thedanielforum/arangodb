package arangodb

import (
	"time"
	"net/http"
	"os"
	"github.com/apex/log/handlers/text"
	"github.com/apex/log"
	"sync"
)

const(
	defaultTimeOut = time.Second * 10
)

func init() {
	// Set up APEX log handler
	log.SetHandler(text.New(os.Stderr))
}

// Config for the database session.
type Opts struct {
	Timeout           time.Duration
	KeepAlivePeriod   time.Duration
	// By default use JWT to authenticate.
	// TODO create basic auth function lol c:
	UseHttpBasicAuth  bool
	// Log all http requests to db.
	DebugMode         bool
}

type Connection struct {
	client *http.Client
	header http.Header

	mu     sync.Mutex
	// Connection options
	opts   *Opts
	// Host address
	host   string
	// Database
	db     string
	// Authentication token
	token  string
}

func NewConnection(host, username, password string, opts *Opts) (*Connection, error) {
	var err error
	c := new(Connection)
	c.opts    = opts
	c.host = buildHostAddress(host, false)
	c.header  = http.Header{}

	// Set default headers
	c.header.Set("Content-Type", "application/json")

	// Set custom timeout.
	// See https://goo.gl/NLk64L
	timeOut := defaultTimeOut
	if c.opts.Timeout > 0 {
		timeOut = c.opts.Timeout
	}

	// Connect to server
	c.client = &http.Client{
		Timeout: timeOut,
	}

	// Authenticate to the database
	err = c.authenticate(username, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

