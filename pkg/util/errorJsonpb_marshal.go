package util

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/genproto/googleapis/rpc/status"
)

var _ gwruntime.Marshaler = (*ErrorJSONPb)(nil)

// ErrorJSONPb is a gwruntime.Marshaler that uses github.com/gogo/protobuf/jsonpb.
type ErrorJSONPb struct {
	marshaler runtime.Marshaler
	jsonpb.Marshaler
}

// NewErrorJSON - Creates a new instance of Error JSON Marshaler
func NewErrorJSON(marshaler runtime.Marshaler) *ErrorJSONPb {
	return &ErrorJSONPb{marshaler: marshaler}
}

// ContentType implements gwruntime.Marshaler.
func (j *ErrorJSONPb) ContentType() string {
	// NB: This is the same as httputil.JSONContentType which we can't use due to
	// an import cycle.
	const JSONContentType = "application/json"
	return JSONContentType
}

func errorChunk(err *status.Status) map[string]proto.Message {
	return map[string]proto.Message{"error": (*status.Status)(err)}
}

// Marshal implements gwruntime.Marshaler.
func (j *ErrorJSONPb) Marshal(v interface{}) ([]byte, error) {
	status, ok := v.(*status.Status)
	if ok {
		fmt.Println("Marshal Type ", status, reflect.TypeOf(v))
		v = errorChunk(status)
	} else {
		streamProtoMessage, ok := v.(map[string]proto.Message)
		if ok {
			fmt.Println("Stream Marshal Type ", streamProtoMessage, reflect.TypeOf(v))
		}
	}

	return j.marshaler.Marshal(v)
}

// Unmarshal implements gwruntime.Marshaler.
func (j *ErrorJSONPb) Unmarshal(data []byte, v interface{}) error {
	return j.marshaler.Unmarshal(data, v)
}

// NewDecoder implements gwruntime.Marshaler.
func (j *ErrorJSONPb) NewDecoder(r io.Reader) gwruntime.Decoder {
	return j.marshaler.NewDecoder(r)
}

// NewEncoder implements gwruntime.Marshaler.
func (j *ErrorJSONPb) NewEncoder(w io.Writer) gwruntime.Encoder {
	return j.marshaler.NewEncoder(w)
}

var _ gwruntime.Delimited = (*ErrorJSONPb)(nil)

// Delimiter implements gwruntime.Delimited.
func (j *ErrorJSONPb) Delimiter() []byte {
	return []byte("\n")
}
