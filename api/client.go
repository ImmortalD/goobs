package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andreykaipov/goobs/api/opcodes"
	uuid "github.com/nu7hatch/gouuid"
)

// Logger is a interface compatible with both the stdlib's logger and some
// third-party loggers.
type Logger interface {
	Printf(string, ...interface{})
}

type ResponsePair struct {
	*opcodes.RequestResponse
	ResponseType interface{}
}

// Client represents a requests client to the OBS websocket server. Its
// intention is to provide a means of communication between the top-level client
// and the category-level clients, so while its fields are exported, they should
// be of no interest to consumers of this library.
type Client struct {
	// The time we're willing to wait to receive a response from the server.
	ResponseTimeout time.Duration

	// Events broadcast by the server when actions happen within OBS.
	IncomingEvents chan interface{}

	// Raw JSON message responses from the websocket server.
	IncomingResponses chan *ResponsePair

	// Raw JSON message responses from the websocket server.
	Opcodes chan opcodes.Opcode

	Log Logger
}

// SendRequest abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
//
// To get the response for a sent request, we can just read the next response
// from our channel. This works fine in a single-threaded context, and the
// message IDs of both the sent request and response should match.
//
// In a concurrent context, this isn't necessarily true, but since
// gorilla/websocket doesn't handle concurrency anyways, who cares? We could
// technically add a mutex in between sending our request and reading from the
// channel, but ehh...
//
// Interestingly, it does seem thread-safe if I use a totally different
// connection, in that connection A won't get a response from OBS for a request
// from connection B. So, message IDs must be unique per client? More
// interestingly, events appear to be broadcast to every client, maybe because
// they have no associated message ID?
//
//
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#request-opcode-6
//
func (c *Client) SendRequest(name string, params interface{}) (interface{}, error) {
	c.Log.Printf("Sending %s Request", name)

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	id := uid.String()

	c.Opcodes <- &opcodes.Request{
		RequestType: name,
		RequestID:   id,
		RequestData: params,
	}

	var pair *ResponsePair
	select {
	case pair = <-c.IncomingResponses:
	case <-time.After(c.ResponseTimeout * time.Millisecond):
		return nil, fmt.Errorf("Timed out receiving response from server for request %q", name)
	}

	response := pair.RequestResponse
	responseType := pair.ResponseType

	// request ID on response should be a mirror of what the client sent, so
	// this is good to check. however, I'm pretty sure a mismatch like this
	// could only happen in a concurrent context... and I'm pretty sure
	// gorilla/websocket would panic before then, so... 🤷
	if response.RequestID != id {
		return nil, fmt.Errorf(
			"Sent request %s, with message ID %s, but next response in channel has message ID %s",
			name,
			id,
			response.RequestID,
		)
	}

	status := response.RequestStatus

	if status.Code != 100 {
		return nil, fmt.Errorf(
			"Got error from OBS executing request %q: %s",
			name,
			status.Comment,
		)
	}

	data := response.ResponseData

	if err := json.Unmarshal(data, responseType); err != nil {
		return nil, fmt.Errorf(
			"Couldn't unmarshal `%s` into an request response type of %q: %s",
			data,
			responseType,
			err,
		)
	}

	return responseType, nil
}