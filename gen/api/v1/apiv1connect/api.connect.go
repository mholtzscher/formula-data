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
	// FormulaDataServiceGetSeasonByIdProcedure is the fully-qualified name of the FormulaDataService's
	// GetSeasonById RPC.
	FormulaDataServiceGetSeasonByIdProcedure = "/api.v1.FormulaDataService/GetSeasonById"
	// FormulaDataServiceGetAllSeasonsProcedure is the fully-qualified name of the FormulaDataService's
	// GetAllSeasons RPC.
	FormulaDataServiceGetAllSeasonsProcedure = "/api.v1.FormulaDataService/GetAllSeasons"
	// FormulaDataServiceCreateDriverProcedure is the fully-qualified name of the FormulaDataService's
	// CreateDriver RPC.
	FormulaDataServiceCreateDriverProcedure = "/api.v1.FormulaDataService/CreateDriver"
	// FormulaDataServiceGetDriverByIdProcedure is the fully-qualified name of the FormulaDataService's
	// GetDriverById RPC.
	FormulaDataServiceGetDriverByIdProcedure = "/api.v1.FormulaDataService/GetDriverById"
	// FormulaDataServiceCreateTeamProcedure is the fully-qualified name of the FormulaDataService's
	// CreateTeam RPC.
	FormulaDataServiceCreateTeamProcedure = "/api.v1.FormulaDataService/CreateTeam"
	// FormulaDataServiceCreateRaceProcedure is the fully-qualified name of the FormulaDataService's
	// CreateRace RPC.
	FormulaDataServiceCreateRaceProcedure = "/api.v1.FormulaDataService/CreateRace"
	// FormulaDataServiceGetRaceByIdProcedure is the fully-qualified name of the FormulaDataService's
	// GetRaceById RPC.
	FormulaDataServiceGetRaceByIdProcedure = "/api.v1.FormulaDataService/GetRaceById"
	// FormulaDataServiceCreateResultProcedure is the fully-qualified name of the FormulaDataService's
	// CreateResult RPC.
	FormulaDataServiceCreateResultProcedure = "/api.v1.FormulaDataService/CreateResult"
	// FormulaDataServiceGetResultByIdProcedure is the fully-qualified name of the FormulaDataService's
	// GetResultById RPC.
	FormulaDataServiceGetResultByIdProcedure = "/api.v1.FormulaDataService/GetResultById"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	formulaDataServiceServiceDescriptor             = v1.File_api_v1_api_proto.Services().ByName("FormulaDataService")
	formulaDataServiceCreateSeasonMethodDescriptor  = formulaDataServiceServiceDescriptor.Methods().ByName("CreateSeason")
	formulaDataServiceGetSeasonByIdMethodDescriptor = formulaDataServiceServiceDescriptor.Methods().ByName("GetSeasonById")
	formulaDataServiceGetAllSeasonsMethodDescriptor = formulaDataServiceServiceDescriptor.Methods().ByName("GetAllSeasons")
	formulaDataServiceCreateDriverMethodDescriptor  = formulaDataServiceServiceDescriptor.Methods().ByName("CreateDriver")
	formulaDataServiceGetDriverByIdMethodDescriptor = formulaDataServiceServiceDescriptor.Methods().ByName("GetDriverById")
	formulaDataServiceCreateTeamMethodDescriptor    = formulaDataServiceServiceDescriptor.Methods().ByName("CreateTeam")
	formulaDataServiceCreateRaceMethodDescriptor    = formulaDataServiceServiceDescriptor.Methods().ByName("CreateRace")
	formulaDataServiceGetRaceByIdMethodDescriptor   = formulaDataServiceServiceDescriptor.Methods().ByName("GetRaceById")
	formulaDataServiceCreateResultMethodDescriptor  = formulaDataServiceServiceDescriptor.Methods().ByName("CreateResult")
	formulaDataServiceGetResultByIdMethodDescriptor = formulaDataServiceServiceDescriptor.Methods().ByName("GetResultById")
)

// FormulaDataServiceClient is a client for the api.v1.FormulaDataService service.
type FormulaDataServiceClient interface {
	CreateSeason(context.Context, *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error)
	GetSeasonById(context.Context, *connect.Request[v1.GetSeasonByIdRequest]) (*connect.Response[v1.GetSeasonByIdResponse], error)
	GetAllSeasons(context.Context, *connect.Request[v1.GetAllSeasonsRequest]) (*connect.Response[v1.GetAllSeasonsResponse], error)
	CreateDriver(context.Context, *connect.Request[v1.CreateDriverRequest]) (*connect.Response[v1.CreateDriverResponse], error)
	GetDriverById(context.Context, *connect.Request[v1.GetDriverByIdRequest]) (*connect.Response[v1.GetDriverByIdResponse], error)
	CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error)
	CreateRace(context.Context, *connect.Request[v1.CreateRaceRequest]) (*connect.Response[v1.CreateRaceResponse], error)
	GetRaceById(context.Context, *connect.Request[v1.GetRaceByIdRequest]) (*connect.Response[v1.GetRaceByIdResponse], error)
	CreateResult(context.Context, *connect.Request[v1.CreateResultRequest]) (*connect.Response[v1.CreateResultResponse], error)
	GetResultById(context.Context, *connect.Request[v1.GetResultByIdRequest]) (*connect.Response[v1.GetResultByIdResponse], error)
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
		getSeasonById: connect.NewClient[v1.GetSeasonByIdRequest, v1.GetSeasonByIdResponse](
			httpClient,
			baseURL+FormulaDataServiceGetSeasonByIdProcedure,
			connect.WithSchema(formulaDataServiceGetSeasonByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAllSeasons: connect.NewClient[v1.GetAllSeasonsRequest, v1.GetAllSeasonsResponse](
			httpClient,
			baseURL+FormulaDataServiceGetAllSeasonsProcedure,
			connect.WithSchema(formulaDataServiceGetAllSeasonsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createDriver: connect.NewClient[v1.CreateDriverRequest, v1.CreateDriverResponse](
			httpClient,
			baseURL+FormulaDataServiceCreateDriverProcedure,
			connect.WithSchema(formulaDataServiceCreateDriverMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDriverById: connect.NewClient[v1.GetDriverByIdRequest, v1.GetDriverByIdResponse](
			httpClient,
			baseURL+FormulaDataServiceGetDriverByIdProcedure,
			connect.WithSchema(formulaDataServiceGetDriverByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createTeam: connect.NewClient[v1.CreateTeamRequest, v1.CreateTeamResponse](
			httpClient,
			baseURL+FormulaDataServiceCreateTeamProcedure,
			connect.WithSchema(formulaDataServiceCreateTeamMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createRace: connect.NewClient[v1.CreateRaceRequest, v1.CreateRaceResponse](
			httpClient,
			baseURL+FormulaDataServiceCreateRaceProcedure,
			connect.WithSchema(formulaDataServiceCreateRaceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRaceById: connect.NewClient[v1.GetRaceByIdRequest, v1.GetRaceByIdResponse](
			httpClient,
			baseURL+FormulaDataServiceGetRaceByIdProcedure,
			connect.WithSchema(formulaDataServiceGetRaceByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createResult: connect.NewClient[v1.CreateResultRequest, v1.CreateResultResponse](
			httpClient,
			baseURL+FormulaDataServiceCreateResultProcedure,
			connect.WithSchema(formulaDataServiceCreateResultMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getResultById: connect.NewClient[v1.GetResultByIdRequest, v1.GetResultByIdResponse](
			httpClient,
			baseURL+FormulaDataServiceGetResultByIdProcedure,
			connect.WithSchema(formulaDataServiceGetResultByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// formulaDataServiceClient implements FormulaDataServiceClient.
type formulaDataServiceClient struct {
	createSeason  *connect.Client[v1.CreateSeasonRequest, v1.CreateSeasonResponse]
	getSeasonById *connect.Client[v1.GetSeasonByIdRequest, v1.GetSeasonByIdResponse]
	getAllSeasons *connect.Client[v1.GetAllSeasonsRequest, v1.GetAllSeasonsResponse]
	createDriver  *connect.Client[v1.CreateDriverRequest, v1.CreateDriverResponse]
	getDriverById *connect.Client[v1.GetDriverByIdRequest, v1.GetDriverByIdResponse]
	createTeam    *connect.Client[v1.CreateTeamRequest, v1.CreateTeamResponse]
	createRace    *connect.Client[v1.CreateRaceRequest, v1.CreateRaceResponse]
	getRaceById   *connect.Client[v1.GetRaceByIdRequest, v1.GetRaceByIdResponse]
	createResult  *connect.Client[v1.CreateResultRequest, v1.CreateResultResponse]
	getResultById *connect.Client[v1.GetResultByIdRequest, v1.GetResultByIdResponse]
}

// CreateSeason calls api.v1.FormulaDataService.CreateSeason.
func (c *formulaDataServiceClient) CreateSeason(ctx context.Context, req *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error) {
	return c.createSeason.CallUnary(ctx, req)
}

// GetSeasonById calls api.v1.FormulaDataService.GetSeasonById.
func (c *formulaDataServiceClient) GetSeasonById(ctx context.Context, req *connect.Request[v1.GetSeasonByIdRequest]) (*connect.Response[v1.GetSeasonByIdResponse], error) {
	return c.getSeasonById.CallUnary(ctx, req)
}

// GetAllSeasons calls api.v1.FormulaDataService.GetAllSeasons.
func (c *formulaDataServiceClient) GetAllSeasons(ctx context.Context, req *connect.Request[v1.GetAllSeasonsRequest]) (*connect.Response[v1.GetAllSeasonsResponse], error) {
	return c.getAllSeasons.CallUnary(ctx, req)
}

// CreateDriver calls api.v1.FormulaDataService.CreateDriver.
func (c *formulaDataServiceClient) CreateDriver(ctx context.Context, req *connect.Request[v1.CreateDriverRequest]) (*connect.Response[v1.CreateDriverResponse], error) {
	return c.createDriver.CallUnary(ctx, req)
}

// GetDriverById calls api.v1.FormulaDataService.GetDriverById.
func (c *formulaDataServiceClient) GetDriverById(ctx context.Context, req *connect.Request[v1.GetDriverByIdRequest]) (*connect.Response[v1.GetDriverByIdResponse], error) {
	return c.getDriverById.CallUnary(ctx, req)
}

// CreateTeam calls api.v1.FormulaDataService.CreateTeam.
func (c *formulaDataServiceClient) CreateTeam(ctx context.Context, req *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error) {
	return c.createTeam.CallUnary(ctx, req)
}

// CreateRace calls api.v1.FormulaDataService.CreateRace.
func (c *formulaDataServiceClient) CreateRace(ctx context.Context, req *connect.Request[v1.CreateRaceRequest]) (*connect.Response[v1.CreateRaceResponse], error) {
	return c.createRace.CallUnary(ctx, req)
}

// GetRaceById calls api.v1.FormulaDataService.GetRaceById.
func (c *formulaDataServiceClient) GetRaceById(ctx context.Context, req *connect.Request[v1.GetRaceByIdRequest]) (*connect.Response[v1.GetRaceByIdResponse], error) {
	return c.getRaceById.CallUnary(ctx, req)
}

// CreateResult calls api.v1.FormulaDataService.CreateResult.
func (c *formulaDataServiceClient) CreateResult(ctx context.Context, req *connect.Request[v1.CreateResultRequest]) (*connect.Response[v1.CreateResultResponse], error) {
	return c.createResult.CallUnary(ctx, req)
}

// GetResultById calls api.v1.FormulaDataService.GetResultById.
func (c *formulaDataServiceClient) GetResultById(ctx context.Context, req *connect.Request[v1.GetResultByIdRequest]) (*connect.Response[v1.GetResultByIdResponse], error) {
	return c.getResultById.CallUnary(ctx, req)
}

// FormulaDataServiceHandler is an implementation of the api.v1.FormulaDataService service.
type FormulaDataServiceHandler interface {
	CreateSeason(context.Context, *connect.Request[v1.CreateSeasonRequest]) (*connect.Response[v1.CreateSeasonResponse], error)
	GetSeasonById(context.Context, *connect.Request[v1.GetSeasonByIdRequest]) (*connect.Response[v1.GetSeasonByIdResponse], error)
	GetAllSeasons(context.Context, *connect.Request[v1.GetAllSeasonsRequest]) (*connect.Response[v1.GetAllSeasonsResponse], error)
	CreateDriver(context.Context, *connect.Request[v1.CreateDriverRequest]) (*connect.Response[v1.CreateDriverResponse], error)
	GetDriverById(context.Context, *connect.Request[v1.GetDriverByIdRequest]) (*connect.Response[v1.GetDriverByIdResponse], error)
	CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error)
	CreateRace(context.Context, *connect.Request[v1.CreateRaceRequest]) (*connect.Response[v1.CreateRaceResponse], error)
	GetRaceById(context.Context, *connect.Request[v1.GetRaceByIdRequest]) (*connect.Response[v1.GetRaceByIdResponse], error)
	CreateResult(context.Context, *connect.Request[v1.CreateResultRequest]) (*connect.Response[v1.CreateResultResponse], error)
	GetResultById(context.Context, *connect.Request[v1.GetResultByIdRequest]) (*connect.Response[v1.GetResultByIdResponse], error)
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
	formulaDataServiceGetSeasonByIdHandler := connect.NewUnaryHandler(
		FormulaDataServiceGetSeasonByIdProcedure,
		svc.GetSeasonById,
		connect.WithSchema(formulaDataServiceGetSeasonByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceGetAllSeasonsHandler := connect.NewUnaryHandler(
		FormulaDataServiceGetAllSeasonsProcedure,
		svc.GetAllSeasons,
		connect.WithSchema(formulaDataServiceGetAllSeasonsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceCreateDriverHandler := connect.NewUnaryHandler(
		FormulaDataServiceCreateDriverProcedure,
		svc.CreateDriver,
		connect.WithSchema(formulaDataServiceCreateDriverMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceGetDriverByIdHandler := connect.NewUnaryHandler(
		FormulaDataServiceGetDriverByIdProcedure,
		svc.GetDriverById,
		connect.WithSchema(formulaDataServiceGetDriverByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceCreateTeamHandler := connect.NewUnaryHandler(
		FormulaDataServiceCreateTeamProcedure,
		svc.CreateTeam,
		connect.WithSchema(formulaDataServiceCreateTeamMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceCreateRaceHandler := connect.NewUnaryHandler(
		FormulaDataServiceCreateRaceProcedure,
		svc.CreateRace,
		connect.WithSchema(formulaDataServiceCreateRaceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceGetRaceByIdHandler := connect.NewUnaryHandler(
		FormulaDataServiceGetRaceByIdProcedure,
		svc.GetRaceById,
		connect.WithSchema(formulaDataServiceGetRaceByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceCreateResultHandler := connect.NewUnaryHandler(
		FormulaDataServiceCreateResultProcedure,
		svc.CreateResult,
		connect.WithSchema(formulaDataServiceCreateResultMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	formulaDataServiceGetResultByIdHandler := connect.NewUnaryHandler(
		FormulaDataServiceGetResultByIdProcedure,
		svc.GetResultById,
		connect.WithSchema(formulaDataServiceGetResultByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.FormulaDataService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FormulaDataServiceCreateSeasonProcedure:
			formulaDataServiceCreateSeasonHandler.ServeHTTP(w, r)
		case FormulaDataServiceGetSeasonByIdProcedure:
			formulaDataServiceGetSeasonByIdHandler.ServeHTTP(w, r)
		case FormulaDataServiceGetAllSeasonsProcedure:
			formulaDataServiceGetAllSeasonsHandler.ServeHTTP(w, r)
		case FormulaDataServiceCreateDriverProcedure:
			formulaDataServiceCreateDriverHandler.ServeHTTP(w, r)
		case FormulaDataServiceGetDriverByIdProcedure:
			formulaDataServiceGetDriverByIdHandler.ServeHTTP(w, r)
		case FormulaDataServiceCreateTeamProcedure:
			formulaDataServiceCreateTeamHandler.ServeHTTP(w, r)
		case FormulaDataServiceCreateRaceProcedure:
			formulaDataServiceCreateRaceHandler.ServeHTTP(w, r)
		case FormulaDataServiceGetRaceByIdProcedure:
			formulaDataServiceGetRaceByIdHandler.ServeHTTP(w, r)
		case FormulaDataServiceCreateResultProcedure:
			formulaDataServiceCreateResultHandler.ServeHTTP(w, r)
		case FormulaDataServiceGetResultByIdProcedure:
			formulaDataServiceGetResultByIdHandler.ServeHTTP(w, r)
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

func (UnimplementedFormulaDataServiceHandler) GetSeasonById(context.Context, *connect.Request[v1.GetSeasonByIdRequest]) (*connect.Response[v1.GetSeasonByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.GetSeasonById is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) GetAllSeasons(context.Context, *connect.Request[v1.GetAllSeasonsRequest]) (*connect.Response[v1.GetAllSeasonsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.GetAllSeasons is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) CreateDriver(context.Context, *connect.Request[v1.CreateDriverRequest]) (*connect.Response[v1.CreateDriverResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.CreateDriver is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) GetDriverById(context.Context, *connect.Request[v1.GetDriverByIdRequest]) (*connect.Response[v1.GetDriverByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.GetDriverById is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) CreateTeam(context.Context, *connect.Request[v1.CreateTeamRequest]) (*connect.Response[v1.CreateTeamResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.CreateTeam is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) CreateRace(context.Context, *connect.Request[v1.CreateRaceRequest]) (*connect.Response[v1.CreateRaceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.CreateRace is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) GetRaceById(context.Context, *connect.Request[v1.GetRaceByIdRequest]) (*connect.Response[v1.GetRaceByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.GetRaceById is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) CreateResult(context.Context, *connect.Request[v1.CreateResultRequest]) (*connect.Response[v1.CreateResultResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.CreateResult is not implemented"))
}

func (UnimplementedFormulaDataServiceHandler) GetResultById(context.Context, *connect.Request[v1.GetResultByIdRequest]) (*connect.Response[v1.GetResultByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.FormulaDataService.GetResultById is not implemented"))
}
