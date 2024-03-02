// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: simoompb/v1/simoom.proto

package simoompbconnect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/minguu42/simoom/lib/go/simoompb/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_7_0

const (
	// SimoomServiceName is the fully-qualified name of the SimoomService service.
	SimoomServiceName = "simoompb.v1.SimoomService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SimoomServiceCheckHealthProcedure is the fully-qualified name of the SimoomService's CheckHealth
	// RPC.
	SimoomServiceCheckHealthProcedure = "/simoompb.v1.SimoomService/CheckHealth"
	// SimoomServiceSignUpProcedure is the fully-qualified name of the SimoomService's SignUp RPC.
	SimoomServiceSignUpProcedure = "/simoompb.v1.SimoomService/SignUp"
	// SimoomServiceSignInProcedure is the fully-qualified name of the SimoomService's SignIn RPC.
	SimoomServiceSignInProcedure = "/simoompb.v1.SimoomService/SignIn"
	// SimoomServiceRefreshTokenProcedure is the fully-qualified name of the SimoomService's
	// RefreshToken RPC.
	SimoomServiceRefreshTokenProcedure = "/simoompb.v1.SimoomService/RefreshToken"
	// SimoomServiceCreateProjectProcedure is the fully-qualified name of the SimoomService's
	// CreateProject RPC.
	SimoomServiceCreateProjectProcedure = "/simoompb.v1.SimoomService/CreateProject"
	// SimoomServiceListProjectsProcedure is the fully-qualified name of the SimoomService's
	// ListProjects RPC.
	SimoomServiceListProjectsProcedure = "/simoompb.v1.SimoomService/ListProjects"
	// SimoomServiceUpdateProjectProcedure is the fully-qualified name of the SimoomService's
	// UpdateProject RPC.
	SimoomServiceUpdateProjectProcedure = "/simoompb.v1.SimoomService/UpdateProject"
	// SimoomServiceDeleteProjectProcedure is the fully-qualified name of the SimoomService's
	// DeleteProject RPC.
	SimoomServiceDeleteProjectProcedure = "/simoompb.v1.SimoomService/DeleteProject"
	// SimoomServiceCreateTaskProcedure is the fully-qualified name of the SimoomService's CreateTask
	// RPC.
	SimoomServiceCreateTaskProcedure = "/simoompb.v1.SimoomService/CreateTask"
	// SimoomServiceListTasksProcedure is the fully-qualified name of the SimoomService's ListTasks RPC.
	SimoomServiceListTasksProcedure = "/simoompb.v1.SimoomService/ListTasks"
	// SimoomServiceUpdateTaskProcedure is the fully-qualified name of the SimoomService's UpdateTask
	// RPC.
	SimoomServiceUpdateTaskProcedure = "/simoompb.v1.SimoomService/UpdateTask"
	// SimoomServiceDeleteTaskProcedure is the fully-qualified name of the SimoomService's DeleteTask
	// RPC.
	SimoomServiceDeleteTaskProcedure = "/simoompb.v1.SimoomService/DeleteTask"
	// SimoomServiceCreateStepProcedure is the fully-qualified name of the SimoomService's CreateStep
	// RPC.
	SimoomServiceCreateStepProcedure = "/simoompb.v1.SimoomService/CreateStep"
	// SimoomServiceUpdateStepProcedure is the fully-qualified name of the SimoomService's UpdateStep
	// RPC.
	SimoomServiceUpdateStepProcedure = "/simoompb.v1.SimoomService/UpdateStep"
	// SimoomServiceDeleteStepProcedure is the fully-qualified name of the SimoomService's DeleteStep
	// RPC.
	SimoomServiceDeleteStepProcedure = "/simoompb.v1.SimoomService/DeleteStep"
	// SimoomServiceCreateTagProcedure is the fully-qualified name of the SimoomService's CreateTag RPC.
	SimoomServiceCreateTagProcedure = "/simoompb.v1.SimoomService/CreateTag"
	// SimoomServiceListTagsProcedure is the fully-qualified name of the SimoomService's ListTags RPC.
	SimoomServiceListTagsProcedure = "/simoompb.v1.SimoomService/ListTags"
	// SimoomServiceUpdateTagProcedure is the fully-qualified name of the SimoomService's UpdateTag RPC.
	SimoomServiceUpdateTagProcedure = "/simoompb.v1.SimoomService/UpdateTag"
	// SimoomServiceDeleteTagProcedure is the fully-qualified name of the SimoomService's DeleteTag RPC.
	SimoomServiceDeleteTagProcedure = "/simoompb.v1.SimoomService/DeleteTag"
)

// SimoomServiceClient is a client for the simoompb.v1.SimoomService service.
type SimoomServiceClient interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
	SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[v1.SignUpResponse], error)
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error)
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.Project], error)
	ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.Projects], error)
	UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.Project], error)
	DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error)
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.Task], error)
	ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.Tasks], error)
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.Task], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error)
	CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.Step], error)
	UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.Step], error)
	DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error)
	CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.Tag], error)
	ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.Tags], error)
	UpdateTag(context.Context, *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.Tag], error)
	DeleteTag(context.Context, *connect.Request[v1.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewSimoomServiceClient constructs a client for the simoompb.v1.SimoomService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSimoomServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SimoomServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &simoomServiceClient{
		checkHealth: connect.NewClient[v1.CheckHealthRequest, v1.CheckHealthResponse](
			httpClient,
			baseURL+SimoomServiceCheckHealthProcedure,
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		signUp: connect.NewClient[v1.SignUpRequest, v1.SignUpResponse](
			httpClient,
			baseURL+SimoomServiceSignUpProcedure,
			opts...,
		),
		signIn: connect.NewClient[v1.SignInRequest, v1.SignInResponse](
			httpClient,
			baseURL+SimoomServiceSignInProcedure,
			opts...,
		),
		refreshToken: connect.NewClient[v1.RefreshTokenRequest, v1.RefreshTokenResponse](
			httpClient,
			baseURL+SimoomServiceRefreshTokenProcedure,
			opts...,
		),
		createProject: connect.NewClient[v1.CreateProjectRequest, v1.Project](
			httpClient,
			baseURL+SimoomServiceCreateProjectProcedure,
			opts...,
		),
		listProjects: connect.NewClient[v1.ListProjectsRequest, v1.Projects](
			httpClient,
			baseURL+SimoomServiceListProjectsProcedure,
			opts...,
		),
		updateProject: connect.NewClient[v1.UpdateProjectRequest, v1.Project](
			httpClient,
			baseURL+SimoomServiceUpdateProjectProcedure,
			opts...,
		),
		deleteProject: connect.NewClient[v1.DeleteProjectRequest, emptypb.Empty](
			httpClient,
			baseURL+SimoomServiceDeleteProjectProcedure,
			opts...,
		),
		createTask: connect.NewClient[v1.CreateTaskRequest, v1.Task](
			httpClient,
			baseURL+SimoomServiceCreateTaskProcedure,
			opts...,
		),
		listTasks: connect.NewClient[v1.ListTasksRequest, v1.Tasks](
			httpClient,
			baseURL+SimoomServiceListTasksProcedure,
			opts...,
		),
		updateTask: connect.NewClient[v1.UpdateTaskRequest, v1.Task](
			httpClient,
			baseURL+SimoomServiceUpdateTaskProcedure,
			opts...,
		),
		deleteTask: connect.NewClient[v1.DeleteTaskRequest, emptypb.Empty](
			httpClient,
			baseURL+SimoomServiceDeleteTaskProcedure,
			opts...,
		),
		createStep: connect.NewClient[v1.CreateStepRequest, v1.Step](
			httpClient,
			baseURL+SimoomServiceCreateStepProcedure,
			opts...,
		),
		updateStep: connect.NewClient[v1.UpdateStepRequest, v1.Step](
			httpClient,
			baseURL+SimoomServiceUpdateStepProcedure,
			opts...,
		),
		deleteStep: connect.NewClient[v1.DeleteStepRequest, emptypb.Empty](
			httpClient,
			baseURL+SimoomServiceDeleteStepProcedure,
			opts...,
		),
		createTag: connect.NewClient[v1.CreateTagRequest, v1.Tag](
			httpClient,
			baseURL+SimoomServiceCreateTagProcedure,
			opts...,
		),
		listTags: connect.NewClient[v1.ListTagsRequest, v1.Tags](
			httpClient,
			baseURL+SimoomServiceListTagsProcedure,
			opts...,
		),
		updateTag: connect.NewClient[v1.UpdateTagRequest, v1.Tag](
			httpClient,
			baseURL+SimoomServiceUpdateTagProcedure,
			opts...,
		),
		deleteTag: connect.NewClient[v1.DeleteTagRequest, emptypb.Empty](
			httpClient,
			baseURL+SimoomServiceDeleteTagProcedure,
			opts...,
		),
	}
}

// simoomServiceClient implements SimoomServiceClient.
type simoomServiceClient struct {
	checkHealth   *connect.Client[v1.CheckHealthRequest, v1.CheckHealthResponse]
	signUp        *connect.Client[v1.SignUpRequest, v1.SignUpResponse]
	signIn        *connect.Client[v1.SignInRequest, v1.SignInResponse]
	refreshToken  *connect.Client[v1.RefreshTokenRequest, v1.RefreshTokenResponse]
	createProject *connect.Client[v1.CreateProjectRequest, v1.Project]
	listProjects  *connect.Client[v1.ListProjectsRequest, v1.Projects]
	updateProject *connect.Client[v1.UpdateProjectRequest, v1.Project]
	deleteProject *connect.Client[v1.DeleteProjectRequest, emptypb.Empty]
	createTask    *connect.Client[v1.CreateTaskRequest, v1.Task]
	listTasks     *connect.Client[v1.ListTasksRequest, v1.Tasks]
	updateTask    *connect.Client[v1.UpdateTaskRequest, v1.Task]
	deleteTask    *connect.Client[v1.DeleteTaskRequest, emptypb.Empty]
	createStep    *connect.Client[v1.CreateStepRequest, v1.Step]
	updateStep    *connect.Client[v1.UpdateStepRequest, v1.Step]
	deleteStep    *connect.Client[v1.DeleteStepRequest, emptypb.Empty]
	createTag     *connect.Client[v1.CreateTagRequest, v1.Tag]
	listTags      *connect.Client[v1.ListTagsRequest, v1.Tags]
	updateTag     *connect.Client[v1.UpdateTagRequest, v1.Tag]
	deleteTag     *connect.Client[v1.DeleteTagRequest, emptypb.Empty]
}

// CheckHealth calls simoompb.v1.SimoomService.CheckHealth.
func (c *simoomServiceClient) CheckHealth(ctx context.Context, req *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return c.checkHealth.CallUnary(ctx, req)
}

// SignUp calls simoompb.v1.SimoomService.SignUp.
func (c *simoomServiceClient) SignUp(ctx context.Context, req *connect.Request[v1.SignUpRequest]) (*connect.Response[v1.SignUpResponse], error) {
	return c.signUp.CallUnary(ctx, req)
}

// SignIn calls simoompb.v1.SimoomService.SignIn.
func (c *simoomServiceClient) SignIn(ctx context.Context, req *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// RefreshToken calls simoompb.v1.SimoomService.RefreshToken.
func (c *simoomServiceClient) RefreshToken(ctx context.Context, req *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return c.refreshToken.CallUnary(ctx, req)
}

// CreateProject calls simoompb.v1.SimoomService.CreateProject.
func (c *simoomServiceClient) CreateProject(ctx context.Context, req *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.Project], error) {
	return c.createProject.CallUnary(ctx, req)
}

// ListProjects calls simoompb.v1.SimoomService.ListProjects.
func (c *simoomServiceClient) ListProjects(ctx context.Context, req *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.Projects], error) {
	return c.listProjects.CallUnary(ctx, req)
}

// UpdateProject calls simoompb.v1.SimoomService.UpdateProject.
func (c *simoomServiceClient) UpdateProject(ctx context.Context, req *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.Project], error) {
	return c.updateProject.CallUnary(ctx, req)
}

// DeleteProject calls simoompb.v1.SimoomService.DeleteProject.
func (c *simoomServiceClient) DeleteProject(ctx context.Context, req *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteProject.CallUnary(ctx, req)
}

// CreateTask calls simoompb.v1.SimoomService.CreateTask.
func (c *simoomServiceClient) CreateTask(ctx context.Context, req *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.Task], error) {
	return c.createTask.CallUnary(ctx, req)
}

// ListTasks calls simoompb.v1.SimoomService.ListTasks.
func (c *simoomServiceClient) ListTasks(ctx context.Context, req *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.Tasks], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// UpdateTask calls simoompb.v1.SimoomService.UpdateTask.
func (c *simoomServiceClient) UpdateTask(ctx context.Context, req *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.Task], error) {
	return c.updateTask.CallUnary(ctx, req)
}

// DeleteTask calls simoompb.v1.SimoomService.DeleteTask.
func (c *simoomServiceClient) DeleteTask(ctx context.Context, req *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// CreateStep calls simoompb.v1.SimoomService.CreateStep.
func (c *simoomServiceClient) CreateStep(ctx context.Context, req *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.Step], error) {
	return c.createStep.CallUnary(ctx, req)
}

// UpdateStep calls simoompb.v1.SimoomService.UpdateStep.
func (c *simoomServiceClient) UpdateStep(ctx context.Context, req *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.Step], error) {
	return c.updateStep.CallUnary(ctx, req)
}

// DeleteStep calls simoompb.v1.SimoomService.DeleteStep.
func (c *simoomServiceClient) DeleteStep(ctx context.Context, req *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteStep.CallUnary(ctx, req)
}

// CreateTag calls simoompb.v1.SimoomService.CreateTag.
func (c *simoomServiceClient) CreateTag(ctx context.Context, req *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.Tag], error) {
	return c.createTag.CallUnary(ctx, req)
}

// ListTags calls simoompb.v1.SimoomService.ListTags.
func (c *simoomServiceClient) ListTags(ctx context.Context, req *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.Tags], error) {
	return c.listTags.CallUnary(ctx, req)
}

// UpdateTag calls simoompb.v1.SimoomService.UpdateTag.
func (c *simoomServiceClient) UpdateTag(ctx context.Context, req *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.Tag], error) {
	return c.updateTag.CallUnary(ctx, req)
}

// DeleteTag calls simoompb.v1.SimoomService.DeleteTag.
func (c *simoomServiceClient) DeleteTag(ctx context.Context, req *connect.Request[v1.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteTag.CallUnary(ctx, req)
}

// SimoomServiceHandler is an implementation of the simoompb.v1.SimoomService service.
type SimoomServiceHandler interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
	SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[v1.SignUpResponse], error)
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error)
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.Project], error)
	ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.Projects], error)
	UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.Project], error)
	DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error)
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.Task], error)
	ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.Tasks], error)
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.Task], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error)
	CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.Step], error)
	UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.Step], error)
	DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error)
	CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.Tag], error)
	ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.Tags], error)
	UpdateTag(context.Context, *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.Tag], error)
	DeleteTag(context.Context, *connect.Request[v1.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewSimoomServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSimoomServiceHandler(svc SimoomServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	simoomServiceCheckHealthHandler := connect.NewUnaryHandler(
		SimoomServiceCheckHealthProcedure,
		svc.CheckHealth,
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	simoomServiceSignUpHandler := connect.NewUnaryHandler(
		SimoomServiceSignUpProcedure,
		svc.SignUp,
		opts...,
	)
	simoomServiceSignInHandler := connect.NewUnaryHandler(
		SimoomServiceSignInProcedure,
		svc.SignIn,
		opts...,
	)
	simoomServiceRefreshTokenHandler := connect.NewUnaryHandler(
		SimoomServiceRefreshTokenProcedure,
		svc.RefreshToken,
		opts...,
	)
	simoomServiceCreateProjectHandler := connect.NewUnaryHandler(
		SimoomServiceCreateProjectProcedure,
		svc.CreateProject,
		opts...,
	)
	simoomServiceListProjectsHandler := connect.NewUnaryHandler(
		SimoomServiceListProjectsProcedure,
		svc.ListProjects,
		opts...,
	)
	simoomServiceUpdateProjectHandler := connect.NewUnaryHandler(
		SimoomServiceUpdateProjectProcedure,
		svc.UpdateProject,
		opts...,
	)
	simoomServiceDeleteProjectHandler := connect.NewUnaryHandler(
		SimoomServiceDeleteProjectProcedure,
		svc.DeleteProject,
		opts...,
	)
	simoomServiceCreateTaskHandler := connect.NewUnaryHandler(
		SimoomServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	)
	simoomServiceListTasksHandler := connect.NewUnaryHandler(
		SimoomServiceListTasksProcedure,
		svc.ListTasks,
		opts...,
	)
	simoomServiceUpdateTaskHandler := connect.NewUnaryHandler(
		SimoomServiceUpdateTaskProcedure,
		svc.UpdateTask,
		opts...,
	)
	simoomServiceDeleteTaskHandler := connect.NewUnaryHandler(
		SimoomServiceDeleteTaskProcedure,
		svc.DeleteTask,
		opts...,
	)
	simoomServiceCreateStepHandler := connect.NewUnaryHandler(
		SimoomServiceCreateStepProcedure,
		svc.CreateStep,
		opts...,
	)
	simoomServiceUpdateStepHandler := connect.NewUnaryHandler(
		SimoomServiceUpdateStepProcedure,
		svc.UpdateStep,
		opts...,
	)
	simoomServiceDeleteStepHandler := connect.NewUnaryHandler(
		SimoomServiceDeleteStepProcedure,
		svc.DeleteStep,
		opts...,
	)
	simoomServiceCreateTagHandler := connect.NewUnaryHandler(
		SimoomServiceCreateTagProcedure,
		svc.CreateTag,
		opts...,
	)
	simoomServiceListTagsHandler := connect.NewUnaryHandler(
		SimoomServiceListTagsProcedure,
		svc.ListTags,
		opts...,
	)
	simoomServiceUpdateTagHandler := connect.NewUnaryHandler(
		SimoomServiceUpdateTagProcedure,
		svc.UpdateTag,
		opts...,
	)
	simoomServiceDeleteTagHandler := connect.NewUnaryHandler(
		SimoomServiceDeleteTagProcedure,
		svc.DeleteTag,
		opts...,
	)
	return "/simoompb.v1.SimoomService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SimoomServiceCheckHealthProcedure:
			simoomServiceCheckHealthHandler.ServeHTTP(w, r)
		case SimoomServiceSignUpProcedure:
			simoomServiceSignUpHandler.ServeHTTP(w, r)
		case SimoomServiceSignInProcedure:
			simoomServiceSignInHandler.ServeHTTP(w, r)
		case SimoomServiceRefreshTokenProcedure:
			simoomServiceRefreshTokenHandler.ServeHTTP(w, r)
		case SimoomServiceCreateProjectProcedure:
			simoomServiceCreateProjectHandler.ServeHTTP(w, r)
		case SimoomServiceListProjectsProcedure:
			simoomServiceListProjectsHandler.ServeHTTP(w, r)
		case SimoomServiceUpdateProjectProcedure:
			simoomServiceUpdateProjectHandler.ServeHTTP(w, r)
		case SimoomServiceDeleteProjectProcedure:
			simoomServiceDeleteProjectHandler.ServeHTTP(w, r)
		case SimoomServiceCreateTaskProcedure:
			simoomServiceCreateTaskHandler.ServeHTTP(w, r)
		case SimoomServiceListTasksProcedure:
			simoomServiceListTasksHandler.ServeHTTP(w, r)
		case SimoomServiceUpdateTaskProcedure:
			simoomServiceUpdateTaskHandler.ServeHTTP(w, r)
		case SimoomServiceDeleteTaskProcedure:
			simoomServiceDeleteTaskHandler.ServeHTTP(w, r)
		case SimoomServiceCreateStepProcedure:
			simoomServiceCreateStepHandler.ServeHTTP(w, r)
		case SimoomServiceUpdateStepProcedure:
			simoomServiceUpdateStepHandler.ServeHTTP(w, r)
		case SimoomServiceDeleteStepProcedure:
			simoomServiceDeleteStepHandler.ServeHTTP(w, r)
		case SimoomServiceCreateTagProcedure:
			simoomServiceCreateTagHandler.ServeHTTP(w, r)
		case SimoomServiceListTagsProcedure:
			simoomServiceListTagsHandler.ServeHTTP(w, r)
		case SimoomServiceUpdateTagProcedure:
			simoomServiceUpdateTagHandler.ServeHTTP(w, r)
		case SimoomServiceDeleteTagProcedure:
			simoomServiceDeleteTagHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSimoomServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSimoomServiceHandler struct{}

func (UnimplementedSimoomServiceHandler) CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.CheckHealth is not implemented"))
}

func (UnimplementedSimoomServiceHandler) SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[v1.SignUpResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.SignUp is not implemented"))
}

func (UnimplementedSimoomServiceHandler) SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.SignIn is not implemented"))
}

func (UnimplementedSimoomServiceHandler) RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.RefreshToken is not implemented"))
}

func (UnimplementedSimoomServiceHandler) CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.Project], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.CreateProject is not implemented"))
}

func (UnimplementedSimoomServiceHandler) ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.Projects], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.ListProjects is not implemented"))
}

func (UnimplementedSimoomServiceHandler) UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.Project], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.UpdateProject is not implemented"))
}

func (UnimplementedSimoomServiceHandler) DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.DeleteProject is not implemented"))
}

func (UnimplementedSimoomServiceHandler) CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.CreateTask is not implemented"))
}

func (UnimplementedSimoomServiceHandler) ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.Tasks], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.ListTasks is not implemented"))
}

func (UnimplementedSimoomServiceHandler) UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.Task], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.UpdateTask is not implemented"))
}

func (UnimplementedSimoomServiceHandler) DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.DeleteTask is not implemented"))
}

func (UnimplementedSimoomServiceHandler) CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.Step], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.CreateStep is not implemented"))
}

func (UnimplementedSimoomServiceHandler) UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.Step], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.UpdateStep is not implemented"))
}

func (UnimplementedSimoomServiceHandler) DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.DeleteStep is not implemented"))
}

func (UnimplementedSimoomServiceHandler) CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.Tag], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.CreateTag is not implemented"))
}

func (UnimplementedSimoomServiceHandler) ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.Tags], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.ListTags is not implemented"))
}

func (UnimplementedSimoomServiceHandler) UpdateTag(context.Context, *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.Tag], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.UpdateTag is not implemented"))
}

func (UnimplementedSimoomServiceHandler) DeleteTag(context.Context, *connect.Request[v1.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.SimoomService.DeleteTag is not implemented"))
}
