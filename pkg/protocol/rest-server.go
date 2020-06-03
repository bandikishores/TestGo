package protocol

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"bandi.com/TestGo/pkg/data"
	"bandi.com/TestGo/pkg/util"
	"github.com/gogo/gateway"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// AcceptHeader is the canonical header name for accept.
	AcceptHeader = "Accept"
	// AcceptEncodingHeader is the canonical header name for accept encoding.
	AcceptEncodingHeader = "Accept-Encoding"
	// ContentEncodingHeader is the canonical header name for content type.
	ContentEncodingHeader = "Content-Encoding"
	// ContentTypeHeader is the canonical header name for content type.
	ContentTypeHeader = "Content-Type"
	// JSONContentType is the JSON content type.
	JSONContentType = "application/json"
	// AltJSONContentType is the alternate JSON content type.
	AltJSONContentType = "application/x-json"
	// ProtoContentType is the protobuf content type.
	ProtoContentType = "application/x-protobuf"
	// AltProtoContentType is the alternate protobuf content type.
	AltProtoContentType = "application/x-google-protobuf"
	// PlaintextContentType is the plaintext content type.
	PlaintextContentType = "text/plain"
	// GzipEncoding is the gzip encoding.
	GzipEncoding = "gzip"
	// JSONPrettyContentType is the JSON content type + Pretty Print
	JSONPrettyContentType = "application/json+pretty"
)

// If our header starts with X-Custom, we let it through
func headerMatcher(header string) (string, bool) {
	if strings.HasPrefix(header, "X-Custom-") {
		return header, true
	}
	return runtime.DefaultHeaderMatcher(header)
}

// RunRestServer runs HTTP/REST gateway
func RunRestServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jsonpb := &gateway.JSONPb{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}
	jsonpbpretty := &gateway.JSONPb{
		Indent:       "  ",
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}
	errorJSONPb := util.NewErrorJSON(jsonpb)
	errorJSONPbPretty := util.NewErrorJSON(jsonpbpretty)
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcher),
		runtime.WithMarshalerOption(JSONContentType, errorJSONPb),
		runtime.WithMarshalerOption(AltJSONContentType, errorJSONPb),
		runtime.WithMarshalerOption(JSONPrettyContentType, errorJSONPbPretty),
		runtime.WithProtoErrorHandler(customHTTPError),
		runtime.WithStreamErrorHandler(handleStreamError),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := data.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	if err := data.RegisterBandiServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Printf("starting skygraph HTTP/REST gateway at port=%v...\n", httpPort)
	return srv.ListenAndServe()
}

// handleStreamError overrides default behavior for computing an error
// message for a server stream.
//
// It uses a default "502 Bad Gateway" HTTP code; only emits "safe"
// messages; and does not set gRPC code or details fields (so they will
// be omitted from the resulting JSON object that is sent to client).
func handleStreamError(ctx context.Context, err error) *runtime.StreamError {
	grpcCode := codes.Unknown
	msg := "unexpected error"
	var grpcDetails []*any.Any
	if s, ok := status.FromError(err); ok {
		grpcCode = s.Code()
		// default message, based on the name of the gRPC code
		// msg = code.String()
		msg = s.Message()
		grpcDetails = s.Proto().GetDetails()
		// see if error details include "safe" message to send
		// to external callers
		/*for _, msg = s.Details() {
			if safe, ok := msg.(*SafeMessage); ok {
				msg = safe.Text
				break
			}
		}*/
	}
	fmt.Println("", grpcDetails)
	httpCode := runtime.HTTPStatusFromCode(grpcCode)
	return &runtime.StreamError{
		GrpcCode:   int32(grpcCode),
		HttpCode:   int32(httpCode),
		HttpStatus: http.StatusText(httpCode),
		Message:    msg,
		Details:    grpcDetails,
	}
}

// CustomStreamError ...
type CustomStreamError struct {
	GrpcCode   int32      `protobuf:"varint,1,opt,name=grpc_code,json=grpcCode,proto3" json:"grpc_code,omitempty"`
	HTTPCode   int32      `protobuf:"varint,2,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Message    string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	HTTPStatus string     `protobuf:"bytes,4,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	Details    []*any.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

// Reset ...
func (m *CustomStreamError) Reset() { *m = CustomStreamError{} }

// String ...
func (m *CustomStreamError) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*CustomStreamError) ProtoMessage() {}

func errorChunk(err error) map[string]interface{} {
	return map[string]interface{}{"error": err}
}
func errorChunkStreamError(err *CustomStreamError) map[string]proto.Message {
	return map[string]proto.Message{"error": err}
}
func customHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	/*	st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *data.Error:
				marshaler
			}
		}
	*/
	/*if ctx.Value(da.OrgID) != nil {
		fmt.Println("Wow Present")
	}
	runtime.DefaultHTTPProtoErrorHandler(ctx, mux, marshaler, w, req, err)*/
	// return Internal when Marshal failed
	const fallback = `{"message": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	contentType := marshaler.ContentType()
	// Check marshaler on run time in order to keep backwards compatability
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	streamError := handleStreamError(ctx, err)
	customStreamError := &CustomStreamError{
		GrpcCode:   streamError.GrpcCode,
		HTTPCode:   streamError.HttpCode,
		HTTPStatus: streamError.HttpStatus,
		Message:    streamError.Message,
		Details:    streamError.Details,
	}
	buf, merr := marshaler.Marshal(errorChunkStreamError(customStreamError))
	if merr != nil {
		fmt.Printf("Failed to Marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			fmt.Printf("Failed to write response: %v", err)
		}
		return
	}
	st := runtime.HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		fmt.Printf("Failed to write response: %v", err)
	}

}
