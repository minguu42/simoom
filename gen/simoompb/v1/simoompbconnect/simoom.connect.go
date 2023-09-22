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
	v1 "github.com/minguu42/simoom/gen/simoompb/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// MonitoringServiceName is the fully-qualified name of the MonitoringService service.
	MonitoringServiceName = "simoompb.v1.MonitoringService"
	// ProjectServiceName is the fully-qualified name of the ProjectService service.
	ProjectServiceName = "simoompb.v1.ProjectService"
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "simoompb.v1.TaskService"
	// StepServiceName is the fully-qualified name of the StepService service.
	StepServiceName = "simoompb.v1.StepService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MonitoringServiceCheckHealthProcedure is the fully-qualified name of the MonitoringService's
	// CheckHealth RPC.
	MonitoringServiceCheckHealthProcedure = "/simoompb.v1.MonitoringService/CheckHealth"
	// ProjectServiceCreateProjectProcedure is the fully-qualified name of the ProjectService's
	// CreateProject RPC.
	ProjectServiceCreateProjectProcedure = "/simoompb.v1.ProjectService/CreateProject"
	// ProjectServiceListProjectsProcedure is the fully-qualified name of the ProjectService's
	// ListProjects RPC.
	ProjectServiceListProjectsProcedure = "/simoompb.v1.ProjectService/ListProjects"
	// ProjectServiceUpdateProjectProcedure is the fully-qualified name of the ProjectService's
	// UpdateProject RPC.
	ProjectServiceUpdateProjectProcedure = "/simoompb.v1.ProjectService/UpdateProject"
	// ProjectServiceDeleteProjectProcedure is the fully-qualified name of the ProjectService's
	// DeleteProject RPC.
	ProjectServiceDeleteProjectProcedure = "/simoompb.v1.ProjectService/DeleteProject"
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/simoompb.v1.TaskService/CreateTask"
	// TaskServiceListTasksProcedure is the fully-qualified name of the TaskService's ListTasks RPC.
	TaskServiceListTasksProcedure = "/simoompb.v1.TaskService/ListTasks"
	// TaskServiceUpdateTaskProcedure is the fully-qualified name of the TaskService's UpdateTask RPC.
	TaskServiceUpdateTaskProcedure = "/simoompb.v1.TaskService/UpdateTask"
	// TaskServiceDeleteTaskProcedure is the fully-qualified name of the TaskService's DeleteTask RPC.
	TaskServiceDeleteTaskProcedure = "/simoompb.v1.TaskService/DeleteTask"
	// StepServiceCreateStepProcedure is the fully-qualified name of the StepService's CreateStep RPC.
	StepServiceCreateStepProcedure = "/simoompb.v1.StepService/CreateStep"
	// StepServiceUpdateStepProcedure is the fully-qualified name of the StepService's UpdateStep RPC.
	StepServiceUpdateStepProcedure = "/simoompb.v1.StepService/UpdateStep"
	// StepServiceDeleteStepProcedure is the fully-qualified name of the StepService's DeleteStep RPC.
	StepServiceDeleteStepProcedure = "/simoompb.v1.StepService/DeleteStep"
)

// MonitoringServiceClient is a client for the simoompb.v1.MonitoringService service.
type MonitoringServiceClient interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
}

// NewMonitoringServiceClient constructs a client for the simoompb.v1.MonitoringService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMonitoringServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MonitoringServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &monitoringServiceClient{
		checkHealth: connect.NewClient[v1.CheckHealthRequest, v1.CheckHealthResponse](
			httpClient,
			baseURL+MonitoringServiceCheckHealthProcedure,
			opts...,
		),
	}
}

// monitoringServiceClient implements MonitoringServiceClient.
type monitoringServiceClient struct {
	checkHealth *connect.Client[v1.CheckHealthRequest, v1.CheckHealthResponse]
}

// CheckHealth calls simoompb.v1.MonitoringService.CheckHealth.
func (c *monitoringServiceClient) CheckHealth(ctx context.Context, req *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return c.checkHealth.CallUnary(ctx, req)
}

// MonitoringServiceHandler is an implementation of the simoompb.v1.MonitoringService service.
type MonitoringServiceHandler interface {
	CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error)
}

// NewMonitoringServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMonitoringServiceHandler(svc MonitoringServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	monitoringServiceCheckHealthHandler := connect.NewUnaryHandler(
		MonitoringServiceCheckHealthProcedure,
		svc.CheckHealth,
		opts...,
	)
	return "/simoompb.v1.MonitoringService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MonitoringServiceCheckHealthProcedure:
			monitoringServiceCheckHealthHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMonitoringServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMonitoringServiceHandler struct{}

func (UnimplementedMonitoringServiceHandler) CheckHealth(context.Context, *connect.Request[v1.CheckHealthRequest]) (*connect.Response[v1.CheckHealthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.MonitoringService.CheckHealth is not implemented"))
}

// ProjectServiceClient is a client for the simoompb.v1.ProjectService service.
type ProjectServiceClient interface {
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.ProjectResponse], error)
	ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ProjectsResponse], error)
	UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.ProjectResponse], error)
	DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewProjectServiceClient constructs a client for the simoompb.v1.ProjectService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &projectServiceClient{
		createProject: connect.NewClient[v1.CreateProjectRequest, v1.ProjectResponse](
			httpClient,
			baseURL+ProjectServiceCreateProjectProcedure,
			opts...,
		),
		listProjects: connect.NewClient[v1.ListProjectsRequest, v1.ProjectsResponse](
			httpClient,
			baseURL+ProjectServiceListProjectsProcedure,
			opts...,
		),
		updateProject: connect.NewClient[v1.UpdateProjectRequest, v1.ProjectResponse](
			httpClient,
			baseURL+ProjectServiceUpdateProjectProcedure,
			opts...,
		),
		deleteProject: connect.NewClient[v1.DeleteProjectRequest, emptypb.Empty](
			httpClient,
			baseURL+ProjectServiceDeleteProjectProcedure,
			opts...,
		),
	}
}

// projectServiceClient implements ProjectServiceClient.
type projectServiceClient struct {
	createProject *connect.Client[v1.CreateProjectRequest, v1.ProjectResponse]
	listProjects  *connect.Client[v1.ListProjectsRequest, v1.ProjectsResponse]
	updateProject *connect.Client[v1.UpdateProjectRequest, v1.ProjectResponse]
	deleteProject *connect.Client[v1.DeleteProjectRequest, emptypb.Empty]
}

// CreateProject calls simoompb.v1.ProjectService.CreateProject.
func (c *projectServiceClient) CreateProject(ctx context.Context, req *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.ProjectResponse], error) {
	return c.createProject.CallUnary(ctx, req)
}

// ListProjects calls simoompb.v1.ProjectService.ListProjects.
func (c *projectServiceClient) ListProjects(ctx context.Context, req *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ProjectsResponse], error) {
	return c.listProjects.CallUnary(ctx, req)
}

// UpdateProject calls simoompb.v1.ProjectService.UpdateProject.
func (c *projectServiceClient) UpdateProject(ctx context.Context, req *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.ProjectResponse], error) {
	return c.updateProject.CallUnary(ctx, req)
}

// DeleteProject calls simoompb.v1.ProjectService.DeleteProject.
func (c *projectServiceClient) DeleteProject(ctx context.Context, req *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteProject.CallUnary(ctx, req)
}

// ProjectServiceHandler is an implementation of the simoompb.v1.ProjectService service.
type ProjectServiceHandler interface {
	CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.ProjectResponse], error)
	ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ProjectsResponse], error)
	UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.ProjectResponse], error)
	DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewProjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectServiceHandler(svc ProjectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	projectServiceCreateProjectHandler := connect.NewUnaryHandler(
		ProjectServiceCreateProjectProcedure,
		svc.CreateProject,
		opts...,
	)
	projectServiceListProjectsHandler := connect.NewUnaryHandler(
		ProjectServiceListProjectsProcedure,
		svc.ListProjects,
		opts...,
	)
	projectServiceUpdateProjectHandler := connect.NewUnaryHandler(
		ProjectServiceUpdateProjectProcedure,
		svc.UpdateProject,
		opts...,
	)
	projectServiceDeleteProjectHandler := connect.NewUnaryHandler(
		ProjectServiceDeleteProjectProcedure,
		svc.DeleteProject,
		opts...,
	)
	return "/simoompb.v1.ProjectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProjectServiceCreateProjectProcedure:
			projectServiceCreateProjectHandler.ServeHTTP(w, r)
		case ProjectServiceListProjectsProcedure:
			projectServiceListProjectsHandler.ServeHTTP(w, r)
		case ProjectServiceUpdateProjectProcedure:
			projectServiceUpdateProjectHandler.ServeHTTP(w, r)
		case ProjectServiceDeleteProjectProcedure:
			projectServiceDeleteProjectHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectServiceHandler struct{}

func (UnimplementedProjectServiceHandler) CreateProject(context.Context, *connect.Request[v1.CreateProjectRequest]) (*connect.Response[v1.ProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.ProjectService.CreateProject is not implemented"))
}

func (UnimplementedProjectServiceHandler) ListProjects(context.Context, *connect.Request[v1.ListProjectsRequest]) (*connect.Response[v1.ProjectsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.ProjectService.ListProjects is not implemented"))
}

func (UnimplementedProjectServiceHandler) UpdateProject(context.Context, *connect.Request[v1.UpdateProjectRequest]) (*connect.Response[v1.ProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.ProjectService.UpdateProject is not implemented"))
}

func (UnimplementedProjectServiceHandler) DeleteProject(context.Context, *connect.Request[v1.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.ProjectService.DeleteProject is not implemented"))
}

// TaskServiceClient is a client for the simoompb.v1.TaskService service.
type TaskServiceClient interface {
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.TaskResponse], error)
	ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.TasksResponse], error)
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.TaskResponse], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTaskServiceClient constructs a client for the simoompb.v1.TaskService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		createTask: connect.NewClient[v1.CreateTaskRequest, v1.TaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			opts...,
		),
		listTasks: connect.NewClient[v1.ListTasksRequest, v1.TasksResponse](
			httpClient,
			baseURL+TaskServiceListTasksProcedure,
			opts...,
		),
		updateTask: connect.NewClient[v1.UpdateTaskRequest, v1.TaskResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskProcedure,
			opts...,
		),
		deleteTask: connect.NewClient[v1.DeleteTaskRequest, emptypb.Empty](
			httpClient,
			baseURL+TaskServiceDeleteTaskProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	createTask *connect.Client[v1.CreateTaskRequest, v1.TaskResponse]
	listTasks  *connect.Client[v1.ListTasksRequest, v1.TasksResponse]
	updateTask *connect.Client[v1.UpdateTaskRequest, v1.TaskResponse]
	deleteTask *connect.Client[v1.DeleteTaskRequest, emptypb.Empty]
}

// CreateTask calls simoompb.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.TaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// ListTasks calls simoompb.v1.TaskService.ListTasks.
func (c *taskServiceClient) ListTasks(ctx context.Context, req *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.TasksResponse], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// UpdateTask calls simoompb.v1.TaskService.UpdateTask.
func (c *taskServiceClient) UpdateTask(ctx context.Context, req *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.TaskResponse], error) {
	return c.updateTask.CallUnary(ctx, req)
}

// DeleteTask calls simoompb.v1.TaskService.DeleteTask.
func (c *taskServiceClient) DeleteTask(ctx context.Context, req *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the simoompb.v1.TaskService service.
type TaskServiceHandler interface {
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.TaskResponse], error)
	ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.TasksResponse], error)
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.TaskResponse], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	taskServiceCreateTaskHandler := connect.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	)
	taskServiceListTasksHandler := connect.NewUnaryHandler(
		TaskServiceListTasksProcedure,
		svc.ListTasks,
		opts...,
	)
	taskServiceUpdateTaskHandler := connect.NewUnaryHandler(
		TaskServiceUpdateTaskProcedure,
		svc.UpdateTask,
		opts...,
	)
	taskServiceDeleteTaskHandler := connect.NewUnaryHandler(
		TaskServiceDeleteTaskProcedure,
		svc.DeleteTask,
		opts...,
	)
	return "/simoompb.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceCreateTaskProcedure:
			taskServiceCreateTaskHandler.ServeHTTP(w, r)
		case TaskServiceListTasksProcedure:
			taskServiceListTasksHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskProcedure:
			taskServiceUpdateTaskHandler.ServeHTTP(w, r)
		case TaskServiceDeleteTaskProcedure:
			taskServiceDeleteTaskHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.TaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.TaskService.CreateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) ListTasks(context.Context, *connect.Request[v1.ListTasksRequest]) (*connect.Response[v1.TasksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.TaskService.ListTasks is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.TaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.TaskService.UpdateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.TaskService.DeleteTask is not implemented"))
}

// StepServiceClient is a client for the simoompb.v1.StepService service.
type StepServiceClient interface {
	CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.StepResponse], error)
	UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.StepResponse], error)
	DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewStepServiceClient constructs a client for the simoompb.v1.StepService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStepServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StepServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &stepServiceClient{
		createStep: connect.NewClient[v1.CreateStepRequest, v1.StepResponse](
			httpClient,
			baseURL+StepServiceCreateStepProcedure,
			opts...,
		),
		updateStep: connect.NewClient[v1.UpdateStepRequest, v1.StepResponse](
			httpClient,
			baseURL+StepServiceUpdateStepProcedure,
			opts...,
		),
		deleteStep: connect.NewClient[v1.DeleteStepRequest, emptypb.Empty](
			httpClient,
			baseURL+StepServiceDeleteStepProcedure,
			opts...,
		),
	}
}

// stepServiceClient implements StepServiceClient.
type stepServiceClient struct {
	createStep *connect.Client[v1.CreateStepRequest, v1.StepResponse]
	updateStep *connect.Client[v1.UpdateStepRequest, v1.StepResponse]
	deleteStep *connect.Client[v1.DeleteStepRequest, emptypb.Empty]
}

// CreateStep calls simoompb.v1.StepService.CreateStep.
func (c *stepServiceClient) CreateStep(ctx context.Context, req *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.StepResponse], error) {
	return c.createStep.CallUnary(ctx, req)
}

// UpdateStep calls simoompb.v1.StepService.UpdateStep.
func (c *stepServiceClient) UpdateStep(ctx context.Context, req *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.StepResponse], error) {
	return c.updateStep.CallUnary(ctx, req)
}

// DeleteStep calls simoompb.v1.StepService.DeleteStep.
func (c *stepServiceClient) DeleteStep(ctx context.Context, req *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteStep.CallUnary(ctx, req)
}

// StepServiceHandler is an implementation of the simoompb.v1.StepService service.
type StepServiceHandler interface {
	CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.StepResponse], error)
	UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.StepResponse], error)
	DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewStepServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStepServiceHandler(svc StepServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	stepServiceCreateStepHandler := connect.NewUnaryHandler(
		StepServiceCreateStepProcedure,
		svc.CreateStep,
		opts...,
	)
	stepServiceUpdateStepHandler := connect.NewUnaryHandler(
		StepServiceUpdateStepProcedure,
		svc.UpdateStep,
		opts...,
	)
	stepServiceDeleteStepHandler := connect.NewUnaryHandler(
		StepServiceDeleteStepProcedure,
		svc.DeleteStep,
		opts...,
	)
	return "/simoompb.v1.StepService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StepServiceCreateStepProcedure:
			stepServiceCreateStepHandler.ServeHTTP(w, r)
		case StepServiceUpdateStepProcedure:
			stepServiceUpdateStepHandler.ServeHTTP(w, r)
		case StepServiceDeleteStepProcedure:
			stepServiceDeleteStepHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStepServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStepServiceHandler struct{}

func (UnimplementedStepServiceHandler) CreateStep(context.Context, *connect.Request[v1.CreateStepRequest]) (*connect.Response[v1.StepResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.StepService.CreateStep is not implemented"))
}

func (UnimplementedStepServiceHandler) UpdateStep(context.Context, *connect.Request[v1.UpdateStepRequest]) (*connect.Response[v1.StepResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.StepService.UpdateStep is not implemented"))
}

func (UnimplementedStepServiceHandler) DeleteStep(context.Context, *connect.Request[v1.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("simoompb.v1.StepService.DeleteStep is not implemented"))
}
