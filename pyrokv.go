package pyrokvgo

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/obsidianreachltd/pyrokv-go/internal/encoder"
	"github.com/obsidianreachltd/pyrokv-go/internal/errors"
	"github.com/obsidianreachltd/pyrokv-go/internal/frame"
	"github.com/obsidianreachltd/pyrokv-go/internal/header"
	"github.com/obsidianreachltd/pyrokv-go/internal/kvstore"
	"github.com/obsidianreachltd/pyrokv-go/internal/payload"
)

type PyroKVClient struct {
	conn      net.Conn
	mtx       sync.Mutex
	requests  map[uint32]chan *frame.Frame
	nextReqID uint32
}

const timeoutDuration = 2 * time.Second

func NewPyroKVClient(password *string) (*PyroKVClient, error) {
	host, exists := os.LookupEnv("PYROKV_HOST")
	if !exists {
		host = "localhost"
	}
	port, exists := os.LookupEnv("PYROKV_PORT")
	if !exists {
		port = "8001"
	}
	// Create a new TCP connection
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	client := &PyroKVClient{
		conn:     conn,
		requests: make(map[uint32]chan *frame.Frame),
	}
	// Start listening for responses
	go client.listenResponses()

	// Authenticate if password is provided
	if password != nil {
		if err := client.Authenticate(*password); err != nil {
			client.Close()
			return nil, err
		}
	}

	return client, nil
}

func (c *PyroKVClient) Close() error {
	return (c.conn).Close()
}

func (c *PyroKVClient) incrementID() uint32 {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.nextReqID = uint32(time.Now().UnixNano())
	return c.nextReqID
}

func (c *PyroKVClient) listenResponses() {
	// TODO: use a proper frame reader that handles partial reads & coalesced frames
	buf := make([]byte, 64<<10)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			return
		}
		fr, err := frame.FrameFromBytes(buf[:n])
		if err != nil {
			continue
		}
		if fr.Header.PayloadLen != uint32(len(fr.Payload)) {
			log.Println("Frame payload length mismatch:", fr.Header.ReqID)
			continue
		}

		c.mtx.Lock()
		ch := c.requests[fr.Header.ReqID]
		c.mtx.Unlock()
		if ch != nil {
			select {
			case ch <- fr:
			default:
			}
		}
	}
}

func (c *PyroKVClient) checkHeaderIsResponseWithNoError(fr *frame.Frame) error {
	if fr.Header.FrameType != header.FrameType_Response {
		return errors.ErrorFromResponsePayload(errors.ErrBadRequest)
	}
	if fr.Header.Flags&header.Flag_Error != 0 {
		return errors.ErrorFromResponsePayload(fr.Payload[0])
	}
	return nil
}

func (c *PyroKVClient) Authenticate(password string) error {
	reqID := c.incrementID()

	pyld := payload.NewAuthRequestPayload(password)
	hd := &header.Header{
		Magic:      header.MAGIC,
		Version:    header.VERSION,
		Operation:  header.OpCode_Auth,
		FrameType:  header.FrameType_Request,
		Flags:      0,
		ReqID:      reqID,
		PayloadLen: uint32(len(pyld)),
	}
	fr := frame.NewFrame(hd, pyld)
	raw := fr.ToBytes()
	respCh := make(chan *frame.Frame, 1)
	c.mtx.Lock()
	c.requests[reqID] = respCh
	c.mtx.Unlock()

	if _, err := c.conn.Write(raw); err != nil {
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return err
	}

	// wait for response OR timeout; no default!
	timer := time.NewTimer(timeoutDuration)
	defer timer.Stop()

	select {
	case res := <-respCh:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()

		if err := c.checkHeaderIsResponseWithNoError(res); err != nil {
			return err
		}
		return nil

	case <-timer.C:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return errors.ErrorFromResponsePayload(errors.ErrClientTimeout)
	}
}

func (c *PyroKVClient) SetWithExpiry(key string, value any, expiry time.Time) error {
	reqID := c.incrementID()

	// Serialize value
	b, err := encoder.Marshal(value)
	if err != nil {
		return err
	}

	pyld := payload.NewSetRequestPayload(key, b, expiry.Unix())
	hd := &header.Header{
		Magic:      header.MAGIC,
		Version:    header.VERSION,
		Operation:  header.OpCode_Set,
		FrameType:  header.FrameType_Request,
		Flags:      0,
		ReqID:      reqID,
		PayloadLen: uint32(len(pyld)),
	}
	fr := frame.NewFrame(hd, pyld)
	raw := fr.ToBytes()

	// register response channel BEFORE write
	respCh := make(chan *frame.Frame, 1)
	c.mtx.Lock()
	c.requests[reqID] = respCh
	c.mtx.Unlock()

	if _, err := c.conn.Write(raw); err != nil {
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return err
	}

	// wait for response OR timeout; no default!
	timer := time.NewTimer(timeoutDuration)
	defer timer.Stop()

	select {
	case res := <-respCh:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()

		if err := c.checkHeaderIsResponseWithNoError(res); err != nil {
			return err
		}
		return nil

	case <-timer.C:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return errors.ErrorFromResponsePayload(errors.ErrClientTimeout)
	}
}

func (c *PyroKVClient) Set(key string, value any) error {
	return c.SetWithExpiry(key, value, time.Time{})
}

func (c *PyroKVClient) GetBytes(key string) ([]byte, error) {
	reqID := c.incrementID()
	pyld := payload.NewGetRequestPayload(key)
	// Build the Header
	hd := &header.Header{
		Magic:      header.MAGIC,
		Version:    header.VERSION,
		Operation:  header.OpCode_Get,
		FrameType:  header.FrameType_Request,
		Flags:      0,
		ReqID:      reqID,
		PayloadLen: uint32(len(pyld)),
	}
	fr := frame.NewFrame(hd, pyld)
	raw := fr.ToBytes()

	respCh := make(chan *frame.Frame, 1)
	c.mtx.Lock()
	c.requests[reqID] = respCh
	c.mtx.Unlock()

	if _, err := c.conn.Write(raw); err != nil {
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return nil, err
	}

	// wait for response OR timeout; no default!
	timer := time.NewTimer(timeoutDuration)
	defer timer.Stop()

	select {
	case res := <-respCh:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()

		if err := c.checkHeaderIsResponseWithNoError(res); err != nil {
			return nil, err
		}
		val, err := kvstore.KVStoreFromBytes(res.Payload)
		if err != nil {
			return nil, err
		}
		return val.Value, nil

	case <-timer.C:
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return nil, errors.ErrorFromResponsePayload(errors.ErrClientTimeout)
	}
}

func (c *PyroKVClient) Get(key string, value any) error {
	data, err := c.GetBytes(key)
	if err != nil {
		return err
	}
	if err := encoder.Unmarshal(data, value); err != nil {
		return err
	}
	return nil
}

func (c *PyroKVClient) Delete(key string) error {
	reqID := c.incrementID()
	pyld := payload.NewGetRequestPayload(key)
	// Build the Header
	hd := &header.Header{
		Magic:      header.MAGIC,
		Version:    header.VERSION,
		Operation:  header.OpCode_Del,
		FrameType:  header.FrameType_Request,
		Flags:      0,
		ReqID:      reqID,
		PayloadLen: uint32(len(pyld)),
	}
	fr := frame.NewFrame(hd, pyld)
	raw := fr.ToBytes()
	respCh := make(chan *frame.Frame, 1)
	c.mtx.Lock()
	c.requests[reqID] = respCh
	c.mtx.Unlock()
	if _, err := c.conn.Write(raw); err != nil {
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return err
	}
	select {
	case res := <-respCh:
		if err := c.checkHeaderIsResponseWithNoError(res); err != nil {
			c.mtx.Lock()
			delete(c.requests, reqID)
			c.mtx.Unlock()
			return err
		}
	case <-time.After(timeoutDuration):
		c.mtx.Lock()
		delete(c.requests, reqID)
		c.mtx.Unlock()
		return errors.ErrorFromResponsePayload(errors.ErrClientTimeout)
	}
	c.mtx.Lock()
	delete(c.requests, reqID)
	c.mtx.Unlock()
	return nil
}
