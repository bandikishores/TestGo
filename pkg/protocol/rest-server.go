package protocol

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"bandi.com/main/pkg/data"
	"bandi.com/main/pkg/util"
	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
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
	protopb := new(util.ProtoPb)
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcher),
		runtime.WithMarshalerOption(JSONContentType, jsonpb),
		runtime.WithMarshalerOption(AltJSONContentType, jsonpb),
		runtime.WithMarshalerOption(JSONPrettyContentType, jsonpbpretty),
		runtime.WithMarshalerOption(ProtoContentType, protopb),
		runtime.WithMarshalerOption(AltProtoContentType, protopb),
	//	runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := data.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
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
