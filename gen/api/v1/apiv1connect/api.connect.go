// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/api.proto

package apiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/mholtzscher/formula-data/gen/api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FormulaDataServiceName is the fully-qualified name of the FormulaDataService service.
	FormulaDataServiceName = "api.v1.FormulaDataService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FormulaDataServiceCreateSeasonProcedure is the fully-qualified name of the FormulaDataService's
	// CreateSeason RPC.
	FormulaDataServiceCreateSeasonProcedure = "/api.v1.FormulaDataService/CreateSeason"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	formulaDataServiceServiceDescriptor            = v1.File_api_v1_api_proto.Services().ByName("FormulaDataService")
	formulaDataServiceCreateSeasonMethodDescriptor = formulaDataServiceServiceDescriptor.Methods().ByName("CreateSeason")
)

// FormulaDataServiceClient is a client for the api.v1.FormulaDataService service.
type FormulaDataServiceClient interface {
	CreateSeason(context.Context, *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error)
}

// NewFormulaDataServiceClient constructs a client for the api.v1.FormulaDataService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFormulaDataServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FormulaDataServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &formulaDataServiceClient{
		createSeason: connect.NewClient[v1.CreateSeasonRequest, v1.CreateSeasonResponse](
			httpClient,
			baseURL+FormulaDataServiceCreateSeasonProcedure,
			connect.WithSchema(formulaDataServiceCreateSeasonMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// formulaDataServiceClient implements FormulaDataServiceClient.
type formulaDataServiceClient struct {
	createSeason *connect.Client[v1.CreateSeasonRequest, v1.CreateSeasonResponse]
}

// CreateSeason calls api.v1.FormulaDataService.CreateSeason.
func (c *formulaDataServiceClient) CreateSeason(ctx context.Context, req *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error) {
	return c.createSeason.CallUnary(ctx, req)
}

// FormulaDataServiceHandler is an implementation of the api.v1.FormulaDataService service.
type FormulaDataServiceHandler interface {
	CreateSeason(context.Context, *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error)
}

// NewFormulaDataServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFormulaDataServiceHandler(svc FormulaDataServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	formulaDataServiceCreateSeasonHandler := connect.NewUnaryHandler(
		FormulaDataServiceCreateSeasonProcedure,
		svc.CreateSeason,
		connect.WithSchema(formulaDataServiceCreateSeasonMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.FormulaDataService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FormulaDataServiceCreateSeasonProcedure:
			formulaDataServiceCreateSeasonHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFormulaDataServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFormulaDataServiceHandler struct{}

func (UnimplementedFormulaDataServiceHandler) CreateSeason(context.Context, *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.CreateSeason is not implemented"))
}