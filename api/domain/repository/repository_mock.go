// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"context"
	"sync"

	"github.com/minguu42/simoom/api/domain/model"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			CreateProjectFunc: func(ctx context.Context, p model.Project) error {
//				panic("mock out the CreateProject method")
//			},
//			CreateStepFunc: func(ctx context.Context, s model.Step) error {
//				panic("mock out the CreateStep method")
//			},
//			CreateTagFunc: func(ctx context.Context, t model.Tag) error {
//				panic("mock out the CreateTag method")
//			},
//			CreateTaskFunc: func(ctx context.Context, t model.Task) error {
//				panic("mock out the CreateTask method")
//			},
//			CreateUserFunc: func(ctx context.Context, u model.User) error {
//				panic("mock out the CreateUser method")
//			},
//			DeleteProjectFunc: func(ctx context.Context, id model.ProjectID) error {
//				panic("mock out the DeleteProject method")
//			},
//			DeleteStepFunc: func(ctx context.Context, id model.StepID) error {
//				panic("mock out the DeleteStep method")
//			},
//			DeleteTagFunc: func(ctx context.Context, id model.TagID) error {
//				panic("mock out the DeleteTag method")
//			},
//			DeleteTaskFunc: func(ctx context.Context, id model.TaskID) error {
//				panic("mock out the DeleteTask method")
//			},
//			GetProjectByIDFunc: func(ctx context.Context, id model.ProjectID) (model.Project, error) {
//				panic("mock out the GetProjectByID method")
//			},
//			GetStepByIDFunc: func(ctx context.Context, id model.StepID) (model.Step, error) {
//				panic("mock out the GetStepByID method")
//			},
//			GetTagByIDFunc: func(ctx context.Context, id model.TagID) (model.Tag, error) {
//				panic("mock out the GetTagByID method")
//			},
//			GetTaskByIDFunc: func(ctx context.Context, id model.TaskID) (model.Task, error) {
//				panic("mock out the GetTaskByID method")
//			},
//			GetUserByEmailFunc: func(ctx context.Context, email string) (model.User, error) {
//				panic("mock out the GetUserByEmail method")
//			},
//			GetUserByIDFunc: func(ctx context.Context, id model.UserID) (model.User, error) {
//				panic("mock out the GetUserByID method")
//			},
//			GetUserByNameFunc: func(ctx context.Context, name string) (model.User, error) {
//				panic("mock out the GetUserByName method")
//			},
//			ListProjectsByUserIDFunc: func(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Project, error) {
//				panic("mock out the ListProjectsByUserID method")
//			},
//			ListTagsByUserIDFunc: func(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Tag, error) {
//				panic("mock out the ListTagsByUserID method")
//			},
//			ListTasksByUserIDFunc: func(ctx context.Context, userID model.UserID, limit uint, offset uint, projectID *model.ProjectID, tagID *model.TagID) ([]model.Task, error) {
//				panic("mock out the ListTasksByUserID method")
//			},
//			TransactionFunc: func(ctx context.Context, fn func(ctxWithTx context.Context) error) error {
//				panic("mock out the Transaction method")
//			},
//			UpdateProjectFunc: func(ctx context.Context, p model.Project) error {
//				panic("mock out the UpdateProject method")
//			},
//			UpdateStepFunc: func(ctx context.Context, s model.Step) error {
//				panic("mock out the UpdateStep method")
//			},
//			UpdateTagFunc: func(ctx context.Context, t model.Tag) error {
//				panic("mock out the UpdateTag method")
//			},
//			UpdateTaskFunc: func(ctx context.Context, t model.Task) error {
//				panic("mock out the UpdateTask method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// CreateProjectFunc mocks the CreateProject method.
	CreateProjectFunc func(ctx context.Context, p model.Project) error

	// CreateStepFunc mocks the CreateStep method.
	CreateStepFunc func(ctx context.Context, s model.Step) error

	// CreateTagFunc mocks the CreateTag method.
	CreateTagFunc func(ctx context.Context, t model.Tag) error

	// CreateTaskFunc mocks the CreateTask method.
	CreateTaskFunc func(ctx context.Context, t model.Task) error

	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(ctx context.Context, u model.User) error

	// DeleteProjectFunc mocks the DeleteProject method.
	DeleteProjectFunc func(ctx context.Context, id model.ProjectID) error

	// DeleteStepFunc mocks the DeleteStep method.
	DeleteStepFunc func(ctx context.Context, id model.StepID) error

	// DeleteTagFunc mocks the DeleteTag method.
	DeleteTagFunc func(ctx context.Context, id model.TagID) error

	// DeleteTaskFunc mocks the DeleteTask method.
	DeleteTaskFunc func(ctx context.Context, id model.TaskID) error

	// GetProjectByIDFunc mocks the GetProjectByID method.
	GetProjectByIDFunc func(ctx context.Context, id model.ProjectID) (model.Project, error)

	// GetStepByIDFunc mocks the GetStepByID method.
	GetStepByIDFunc func(ctx context.Context, id model.StepID) (model.Step, error)

	// GetTagByIDFunc mocks the GetTagByID method.
	GetTagByIDFunc func(ctx context.Context, id model.TagID) (model.Tag, error)

	// GetTaskByIDFunc mocks the GetTaskByID method.
	GetTaskByIDFunc func(ctx context.Context, id model.TaskID) (model.Task, error)

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(ctx context.Context, email string) (model.User, error)

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(ctx context.Context, id model.UserID) (model.User, error)

	// GetUserByNameFunc mocks the GetUserByName method.
	GetUserByNameFunc func(ctx context.Context, name string) (model.User, error)

	// ListProjectsByUserIDFunc mocks the ListProjectsByUserID method.
	ListProjectsByUserIDFunc func(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Project, error)

	// ListTagsByUserIDFunc mocks the ListTagsByUserID method.
	ListTagsByUserIDFunc func(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Tag, error)

	// ListTasksByUserIDFunc mocks the ListTasksByUserID method.
	ListTasksByUserIDFunc func(ctx context.Context, userID model.UserID, limit uint, offset uint, projectID *model.ProjectID, tagID *model.TagID) ([]model.Task, error)

	// TransactionFunc mocks the Transaction method.
	TransactionFunc func(ctx context.Context, fn func(ctxWithTx context.Context) error) error

	// UpdateProjectFunc mocks the UpdateProject method.
	UpdateProjectFunc func(ctx context.Context, p model.Project) error

	// UpdateStepFunc mocks the UpdateStep method.
	UpdateStepFunc func(ctx context.Context, s model.Step) error

	// UpdateTagFunc mocks the UpdateTag method.
	UpdateTagFunc func(ctx context.Context, t model.Tag) error

	// UpdateTaskFunc mocks the UpdateTask method.
	UpdateTaskFunc func(ctx context.Context, t model.Task) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateProject holds details about calls to the CreateProject method.
		CreateProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P model.Project
		}
		// CreateStep holds details about calls to the CreateStep method.
		CreateStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// S is the s argument value.
			S model.Step
		}
		// CreateTag holds details about calls to the CreateTag method.
		CreateTag []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// T is the t argument value.
			T model.Tag
		}
		// CreateTask holds details about calls to the CreateTask method.
		CreateTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// T is the t argument value.
			T model.Task
		}
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// U is the u argument value.
			U model.User
		}
		// DeleteProject holds details about calls to the DeleteProject method.
		DeleteProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.ProjectID
		}
		// DeleteStep holds details about calls to the DeleteStep method.
		DeleteStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.StepID
		}
		// DeleteTag holds details about calls to the DeleteTag method.
		DeleteTag []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.TagID
		}
		// DeleteTask holds details about calls to the DeleteTask method.
		DeleteTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.TaskID
		}
		// GetProjectByID holds details about calls to the GetProjectByID method.
		GetProjectByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.ProjectID
		}
		// GetStepByID holds details about calls to the GetStepByID method.
		GetStepByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.StepID
		}
		// GetTagByID holds details about calls to the GetTagByID method.
		GetTagByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.TagID
		}
		// GetTaskByID holds details about calls to the GetTaskByID method.
		GetTaskByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.TaskID
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID model.UserID
		}
		// GetUserByName holds details about calls to the GetUserByName method.
		GetUserByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// ListProjectsByUserID holds details about calls to the ListProjectsByUserID method.
		ListProjectsByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID model.UserID
			// Limit is the limit argument value.
			Limit uint
			// Offset is the offset argument value.
			Offset uint
		}
		// ListTagsByUserID holds details about calls to the ListTagsByUserID method.
		ListTagsByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID model.UserID
			// Limit is the limit argument value.
			Limit uint
			// Offset is the offset argument value.
			Offset uint
		}
		// ListTasksByUserID holds details about calls to the ListTasksByUserID method.
		ListTasksByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID model.UserID
			// Limit is the limit argument value.
			Limit uint
			// Offset is the offset argument value.
			Offset uint
			// ProjectID is the projectID argument value.
			ProjectID *model.ProjectID
			// TagID is the tagID argument value.
			TagID *model.TagID
		}
		// Transaction holds details about calls to the Transaction method.
		Transaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Fn is the fn argument value.
			Fn func(ctxWithTx context.Context) error
		}
		// UpdateProject holds details about calls to the UpdateProject method.
		UpdateProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P model.Project
		}
		// UpdateStep holds details about calls to the UpdateStep method.
		UpdateStep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// S is the s argument value.
			S model.Step
		}
		// UpdateTag holds details about calls to the UpdateTag method.
		UpdateTag []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// T is the t argument value.
			T model.Tag
		}
		// UpdateTask holds details about calls to the UpdateTask method.
		UpdateTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// T is the t argument value.
			T model.Task
		}
	}
	lockCreateProject        sync.RWMutex
	lockCreateStep           sync.RWMutex
	lockCreateTag            sync.RWMutex
	lockCreateTask           sync.RWMutex
	lockCreateUser           sync.RWMutex
	lockDeleteProject        sync.RWMutex
	lockDeleteStep           sync.RWMutex
	lockDeleteTag            sync.RWMutex
	lockDeleteTask           sync.RWMutex
	lockGetProjectByID       sync.RWMutex
	lockGetStepByID          sync.RWMutex
	lockGetTagByID           sync.RWMutex
	lockGetTaskByID          sync.RWMutex
	lockGetUserByEmail       sync.RWMutex
	lockGetUserByID          sync.RWMutex
	lockGetUserByName        sync.RWMutex
	lockListProjectsByUserID sync.RWMutex
	lockListTagsByUserID     sync.RWMutex
	lockListTasksByUserID    sync.RWMutex
	lockTransaction          sync.RWMutex
	lockUpdateProject        sync.RWMutex
	lockUpdateStep           sync.RWMutex
	lockUpdateTag            sync.RWMutex
	lockUpdateTask           sync.RWMutex
}

// CreateProject calls CreateProjectFunc.
func (mock *RepositoryMock) CreateProject(ctx context.Context, p model.Project) error {
	if mock.CreateProjectFunc == nil {
		panic("RepositoryMock.CreateProjectFunc: method is nil but Repository.CreateProject was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   model.Project
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockCreateProject.Lock()
	mock.calls.CreateProject = append(mock.calls.CreateProject, callInfo)
	mock.lockCreateProject.Unlock()
	return mock.CreateProjectFunc(ctx, p)
}

// CreateProjectCalls gets all the calls that were made to CreateProject.
// Check the length with:
//
//	len(mockedRepository.CreateProjectCalls())
func (mock *RepositoryMock) CreateProjectCalls() []struct {
	Ctx context.Context
	P   model.Project
} {
	var calls []struct {
		Ctx context.Context
		P   model.Project
	}
	mock.lockCreateProject.RLock()
	calls = mock.calls.CreateProject
	mock.lockCreateProject.RUnlock()
	return calls
}

// CreateStep calls CreateStepFunc.
func (mock *RepositoryMock) CreateStep(ctx context.Context, s model.Step) error {
	if mock.CreateStepFunc == nil {
		panic("RepositoryMock.CreateStepFunc: method is nil but Repository.CreateStep was just called")
	}
	callInfo := struct {
		Ctx context.Context
		S   model.Step
	}{
		Ctx: ctx,
		S:   s,
	}
	mock.lockCreateStep.Lock()
	mock.calls.CreateStep = append(mock.calls.CreateStep, callInfo)
	mock.lockCreateStep.Unlock()
	return mock.CreateStepFunc(ctx, s)
}

// CreateStepCalls gets all the calls that were made to CreateStep.
// Check the length with:
//
//	len(mockedRepository.CreateStepCalls())
func (mock *RepositoryMock) CreateStepCalls() []struct {
	Ctx context.Context
	S   model.Step
} {
	var calls []struct {
		Ctx context.Context
		S   model.Step
	}
	mock.lockCreateStep.RLock()
	calls = mock.calls.CreateStep
	mock.lockCreateStep.RUnlock()
	return calls
}

// CreateTag calls CreateTagFunc.
func (mock *RepositoryMock) CreateTag(ctx context.Context, t model.Tag) error {
	if mock.CreateTagFunc == nil {
		panic("RepositoryMock.CreateTagFunc: method is nil but Repository.CreateTag was just called")
	}
	callInfo := struct {
		Ctx context.Context
		T   model.Tag
	}{
		Ctx: ctx,
		T:   t,
	}
	mock.lockCreateTag.Lock()
	mock.calls.CreateTag = append(mock.calls.CreateTag, callInfo)
	mock.lockCreateTag.Unlock()
	return mock.CreateTagFunc(ctx, t)
}

// CreateTagCalls gets all the calls that were made to CreateTag.
// Check the length with:
//
//	len(mockedRepository.CreateTagCalls())
func (mock *RepositoryMock) CreateTagCalls() []struct {
	Ctx context.Context
	T   model.Tag
} {
	var calls []struct {
		Ctx context.Context
		T   model.Tag
	}
	mock.lockCreateTag.RLock()
	calls = mock.calls.CreateTag
	mock.lockCreateTag.RUnlock()
	return calls
}

// CreateTask calls CreateTaskFunc.
func (mock *RepositoryMock) CreateTask(ctx context.Context, t model.Task) error {
	if mock.CreateTaskFunc == nil {
		panic("RepositoryMock.CreateTaskFunc: method is nil but Repository.CreateTask was just called")
	}
	callInfo := struct {
		Ctx context.Context
		T   model.Task
	}{
		Ctx: ctx,
		T:   t,
	}
	mock.lockCreateTask.Lock()
	mock.calls.CreateTask = append(mock.calls.CreateTask, callInfo)
	mock.lockCreateTask.Unlock()
	return mock.CreateTaskFunc(ctx, t)
}

// CreateTaskCalls gets all the calls that were made to CreateTask.
// Check the length with:
//
//	len(mockedRepository.CreateTaskCalls())
func (mock *RepositoryMock) CreateTaskCalls() []struct {
	Ctx context.Context
	T   model.Task
} {
	var calls []struct {
		Ctx context.Context
		T   model.Task
	}
	mock.lockCreateTask.RLock()
	calls = mock.calls.CreateTask
	mock.lockCreateTask.RUnlock()
	return calls
}

// CreateUser calls CreateUserFunc.
func (mock *RepositoryMock) CreateUser(ctx context.Context, u model.User) error {
	if mock.CreateUserFunc == nil {
		panic("RepositoryMock.CreateUserFunc: method is nil but Repository.CreateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		U   model.User
	}{
		Ctx: ctx,
		U:   u,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(ctx, u)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//
//	len(mockedRepository.CreateUserCalls())
func (mock *RepositoryMock) CreateUserCalls() []struct {
	Ctx context.Context
	U   model.User
} {
	var calls []struct {
		Ctx context.Context
		U   model.User
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// DeleteProject calls DeleteProjectFunc.
func (mock *RepositoryMock) DeleteProject(ctx context.Context, id model.ProjectID) error {
	if mock.DeleteProjectFunc == nil {
		panic("RepositoryMock.DeleteProjectFunc: method is nil but Repository.DeleteProject was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.ProjectID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteProject.Lock()
	mock.calls.DeleteProject = append(mock.calls.DeleteProject, callInfo)
	mock.lockDeleteProject.Unlock()
	return mock.DeleteProjectFunc(ctx, id)
}

// DeleteProjectCalls gets all the calls that were made to DeleteProject.
// Check the length with:
//
//	len(mockedRepository.DeleteProjectCalls())
func (mock *RepositoryMock) DeleteProjectCalls() []struct {
	Ctx context.Context
	ID  model.ProjectID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.ProjectID
	}
	mock.lockDeleteProject.RLock()
	calls = mock.calls.DeleteProject
	mock.lockDeleteProject.RUnlock()
	return calls
}

// DeleteStep calls DeleteStepFunc.
func (mock *RepositoryMock) DeleteStep(ctx context.Context, id model.StepID) error {
	if mock.DeleteStepFunc == nil {
		panic("RepositoryMock.DeleteStepFunc: method is nil but Repository.DeleteStep was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.StepID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteStep.Lock()
	mock.calls.DeleteStep = append(mock.calls.DeleteStep, callInfo)
	mock.lockDeleteStep.Unlock()
	return mock.DeleteStepFunc(ctx, id)
}

// DeleteStepCalls gets all the calls that were made to DeleteStep.
// Check the length with:
//
//	len(mockedRepository.DeleteStepCalls())
func (mock *RepositoryMock) DeleteStepCalls() []struct {
	Ctx context.Context
	ID  model.StepID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.StepID
	}
	mock.lockDeleteStep.RLock()
	calls = mock.calls.DeleteStep
	mock.lockDeleteStep.RUnlock()
	return calls
}

// DeleteTag calls DeleteTagFunc.
func (mock *RepositoryMock) DeleteTag(ctx context.Context, id model.TagID) error {
	if mock.DeleteTagFunc == nil {
		panic("RepositoryMock.DeleteTagFunc: method is nil but Repository.DeleteTag was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.TagID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteTag.Lock()
	mock.calls.DeleteTag = append(mock.calls.DeleteTag, callInfo)
	mock.lockDeleteTag.Unlock()
	return mock.DeleteTagFunc(ctx, id)
}

// DeleteTagCalls gets all the calls that were made to DeleteTag.
// Check the length with:
//
//	len(mockedRepository.DeleteTagCalls())
func (mock *RepositoryMock) DeleteTagCalls() []struct {
	Ctx context.Context
	ID  model.TagID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.TagID
	}
	mock.lockDeleteTag.RLock()
	calls = mock.calls.DeleteTag
	mock.lockDeleteTag.RUnlock()
	return calls
}

// DeleteTask calls DeleteTaskFunc.
func (mock *RepositoryMock) DeleteTask(ctx context.Context, id model.TaskID) error {
	if mock.DeleteTaskFunc == nil {
		panic("RepositoryMock.DeleteTaskFunc: method is nil but Repository.DeleteTask was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.TaskID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteTask.Lock()
	mock.calls.DeleteTask = append(mock.calls.DeleteTask, callInfo)
	mock.lockDeleteTask.Unlock()
	return mock.DeleteTaskFunc(ctx, id)
}

// DeleteTaskCalls gets all the calls that were made to DeleteTask.
// Check the length with:
//
//	len(mockedRepository.DeleteTaskCalls())
func (mock *RepositoryMock) DeleteTaskCalls() []struct {
	Ctx context.Context
	ID  model.TaskID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.TaskID
	}
	mock.lockDeleteTask.RLock()
	calls = mock.calls.DeleteTask
	mock.lockDeleteTask.RUnlock()
	return calls
}

// GetProjectByID calls GetProjectByIDFunc.
func (mock *RepositoryMock) GetProjectByID(ctx context.Context, id model.ProjectID) (model.Project, error) {
	if mock.GetProjectByIDFunc == nil {
		panic("RepositoryMock.GetProjectByIDFunc: method is nil but Repository.GetProjectByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.ProjectID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetProjectByID.Lock()
	mock.calls.GetProjectByID = append(mock.calls.GetProjectByID, callInfo)
	mock.lockGetProjectByID.Unlock()
	return mock.GetProjectByIDFunc(ctx, id)
}

// GetProjectByIDCalls gets all the calls that were made to GetProjectByID.
// Check the length with:
//
//	len(mockedRepository.GetProjectByIDCalls())
func (mock *RepositoryMock) GetProjectByIDCalls() []struct {
	Ctx context.Context
	ID  model.ProjectID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.ProjectID
	}
	mock.lockGetProjectByID.RLock()
	calls = mock.calls.GetProjectByID
	mock.lockGetProjectByID.RUnlock()
	return calls
}

// GetStepByID calls GetStepByIDFunc.
func (mock *RepositoryMock) GetStepByID(ctx context.Context, id model.StepID) (model.Step, error) {
	if mock.GetStepByIDFunc == nil {
		panic("RepositoryMock.GetStepByIDFunc: method is nil but Repository.GetStepByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.StepID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetStepByID.Lock()
	mock.calls.GetStepByID = append(mock.calls.GetStepByID, callInfo)
	mock.lockGetStepByID.Unlock()
	return mock.GetStepByIDFunc(ctx, id)
}

// GetStepByIDCalls gets all the calls that were made to GetStepByID.
// Check the length with:
//
//	len(mockedRepository.GetStepByIDCalls())
func (mock *RepositoryMock) GetStepByIDCalls() []struct {
	Ctx context.Context
	ID  model.StepID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.StepID
	}
	mock.lockGetStepByID.RLock()
	calls = mock.calls.GetStepByID
	mock.lockGetStepByID.RUnlock()
	return calls
}

// GetTagByID calls GetTagByIDFunc.
func (mock *RepositoryMock) GetTagByID(ctx context.Context, id model.TagID) (model.Tag, error) {
	if mock.GetTagByIDFunc == nil {
		panic("RepositoryMock.GetTagByIDFunc: method is nil but Repository.GetTagByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.TagID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetTagByID.Lock()
	mock.calls.GetTagByID = append(mock.calls.GetTagByID, callInfo)
	mock.lockGetTagByID.Unlock()
	return mock.GetTagByIDFunc(ctx, id)
}

// GetTagByIDCalls gets all the calls that were made to GetTagByID.
// Check the length with:
//
//	len(mockedRepository.GetTagByIDCalls())
func (mock *RepositoryMock) GetTagByIDCalls() []struct {
	Ctx context.Context
	ID  model.TagID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.TagID
	}
	mock.lockGetTagByID.RLock()
	calls = mock.calls.GetTagByID
	mock.lockGetTagByID.RUnlock()
	return calls
}

// GetTaskByID calls GetTaskByIDFunc.
func (mock *RepositoryMock) GetTaskByID(ctx context.Context, id model.TaskID) (model.Task, error) {
	if mock.GetTaskByIDFunc == nil {
		panic("RepositoryMock.GetTaskByIDFunc: method is nil but Repository.GetTaskByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.TaskID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetTaskByID.Lock()
	mock.calls.GetTaskByID = append(mock.calls.GetTaskByID, callInfo)
	mock.lockGetTaskByID.Unlock()
	return mock.GetTaskByIDFunc(ctx, id)
}

// GetTaskByIDCalls gets all the calls that were made to GetTaskByID.
// Check the length with:
//
//	len(mockedRepository.GetTaskByIDCalls())
func (mock *RepositoryMock) GetTaskByIDCalls() []struct {
	Ctx context.Context
	ID  model.TaskID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.TaskID
	}
	mock.lockGetTaskByID.RLock()
	calls = mock.calls.GetTaskByID
	mock.lockGetTaskByID.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *RepositoryMock) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	if mock.GetUserByEmailFunc == nil {
		panic("RepositoryMock.GetUserByEmailFunc: method is nil but Repository.GetUserByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	mock.lockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(ctx, email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//
//	len(mockedRepository.GetUserByEmailCalls())
func (mock *RepositoryMock) GetUserByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	mock.lockGetUserByEmail.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *RepositoryMock) GetUserByID(ctx context.Context, id model.UserID) (model.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("RepositoryMock.GetUserByIDFunc: method is nil but Repository.GetUserByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  model.UserID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(ctx, id)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//
//	len(mockedRepository.GetUserByIDCalls())
func (mock *RepositoryMock) GetUserByIDCalls() []struct {
	Ctx context.Context
	ID  model.UserID
} {
	var calls []struct {
		Ctx context.Context
		ID  model.UserID
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// GetUserByName calls GetUserByNameFunc.
func (mock *RepositoryMock) GetUserByName(ctx context.Context, name string) (model.User, error) {
	if mock.GetUserByNameFunc == nil {
		panic("RepositoryMock.GetUserByNameFunc: method is nil but Repository.GetUserByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetUserByName.Lock()
	mock.calls.GetUserByName = append(mock.calls.GetUserByName, callInfo)
	mock.lockGetUserByName.Unlock()
	return mock.GetUserByNameFunc(ctx, name)
}

// GetUserByNameCalls gets all the calls that were made to GetUserByName.
// Check the length with:
//
//	len(mockedRepository.GetUserByNameCalls())
func (mock *RepositoryMock) GetUserByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockGetUserByName.RLock()
	calls = mock.calls.GetUserByName
	mock.lockGetUserByName.RUnlock()
	return calls
}

// ListProjectsByUserID calls ListProjectsByUserIDFunc.
func (mock *RepositoryMock) ListProjectsByUserID(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Project, error) {
	if mock.ListProjectsByUserIDFunc == nil {
		panic("RepositoryMock.ListProjectsByUserIDFunc: method is nil but Repository.ListProjectsByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID model.UserID
		Limit  uint
		Offset uint
	}{
		Ctx:    ctx,
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
	mock.lockListProjectsByUserID.Lock()
	mock.calls.ListProjectsByUserID = append(mock.calls.ListProjectsByUserID, callInfo)
	mock.lockListProjectsByUserID.Unlock()
	return mock.ListProjectsByUserIDFunc(ctx, userID, limit, offset)
}

// ListProjectsByUserIDCalls gets all the calls that were made to ListProjectsByUserID.
// Check the length with:
//
//	len(mockedRepository.ListProjectsByUserIDCalls())
func (mock *RepositoryMock) ListProjectsByUserIDCalls() []struct {
	Ctx    context.Context
	UserID model.UserID
	Limit  uint
	Offset uint
} {
	var calls []struct {
		Ctx    context.Context
		UserID model.UserID
		Limit  uint
		Offset uint
	}
	mock.lockListProjectsByUserID.RLock()
	calls = mock.calls.ListProjectsByUserID
	mock.lockListProjectsByUserID.RUnlock()
	return calls
}

// ListTagsByUserID calls ListTagsByUserIDFunc.
func (mock *RepositoryMock) ListTagsByUserID(ctx context.Context, userID model.UserID, limit uint, offset uint) ([]model.Tag, error) {
	if mock.ListTagsByUserIDFunc == nil {
		panic("RepositoryMock.ListTagsByUserIDFunc: method is nil but Repository.ListTagsByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID model.UserID
		Limit  uint
		Offset uint
	}{
		Ctx:    ctx,
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
	mock.lockListTagsByUserID.Lock()
	mock.calls.ListTagsByUserID = append(mock.calls.ListTagsByUserID, callInfo)
	mock.lockListTagsByUserID.Unlock()
	return mock.ListTagsByUserIDFunc(ctx, userID, limit, offset)
}

// ListTagsByUserIDCalls gets all the calls that were made to ListTagsByUserID.
// Check the length with:
//
//	len(mockedRepository.ListTagsByUserIDCalls())
func (mock *RepositoryMock) ListTagsByUserIDCalls() []struct {
	Ctx    context.Context
	UserID model.UserID
	Limit  uint
	Offset uint
} {
	var calls []struct {
		Ctx    context.Context
		UserID model.UserID
		Limit  uint
		Offset uint
	}
	mock.lockListTagsByUserID.RLock()
	calls = mock.calls.ListTagsByUserID
	mock.lockListTagsByUserID.RUnlock()
	return calls
}

// ListTasksByUserID calls ListTasksByUserIDFunc.
func (mock *RepositoryMock) ListTasksByUserID(ctx context.Context, userID model.UserID, limit uint, offset uint, projectID *model.ProjectID, tagID *model.TagID) ([]model.Task, error) {
	if mock.ListTasksByUserIDFunc == nil {
		panic("RepositoryMock.ListTasksByUserIDFunc: method is nil but Repository.ListTasksByUserID was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		UserID    model.UserID
		Limit     uint
		Offset    uint
		ProjectID *model.ProjectID
		TagID     *model.TagID
	}{
		Ctx:       ctx,
		UserID:    userID,
		Limit:     limit,
		Offset:    offset,
		ProjectID: projectID,
		TagID:     tagID,
	}
	mock.lockListTasksByUserID.Lock()
	mock.calls.ListTasksByUserID = append(mock.calls.ListTasksByUserID, callInfo)
	mock.lockListTasksByUserID.Unlock()
	return mock.ListTasksByUserIDFunc(ctx, userID, limit, offset, projectID, tagID)
}

// ListTasksByUserIDCalls gets all the calls that were made to ListTasksByUserID.
// Check the length with:
//
//	len(mockedRepository.ListTasksByUserIDCalls())
func (mock *RepositoryMock) ListTasksByUserIDCalls() []struct {
	Ctx       context.Context
	UserID    model.UserID
	Limit     uint
	Offset    uint
	ProjectID *model.ProjectID
	TagID     *model.TagID
} {
	var calls []struct {
		Ctx       context.Context
		UserID    model.UserID
		Limit     uint
		Offset    uint
		ProjectID *model.ProjectID
		TagID     *model.TagID
	}
	mock.lockListTasksByUserID.RLock()
	calls = mock.calls.ListTasksByUserID
	mock.lockListTasksByUserID.RUnlock()
	return calls
}

// Transaction calls TransactionFunc.
func (mock *RepositoryMock) Transaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error {
	if mock.TransactionFunc == nil {
		panic("RepositoryMock.TransactionFunc: method is nil but Repository.Transaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Fn  func(ctxWithTx context.Context) error
	}{
		Ctx: ctx,
		Fn:  fn,
	}
	mock.lockTransaction.Lock()
	mock.calls.Transaction = append(mock.calls.Transaction, callInfo)
	mock.lockTransaction.Unlock()
	return mock.TransactionFunc(ctx, fn)
}

// TransactionCalls gets all the calls that were made to Transaction.
// Check the length with:
//
//	len(mockedRepository.TransactionCalls())
func (mock *RepositoryMock) TransactionCalls() []struct {
	Ctx context.Context
	Fn  func(ctxWithTx context.Context) error
} {
	var calls []struct {
		Ctx context.Context
		Fn  func(ctxWithTx context.Context) error
	}
	mock.lockTransaction.RLock()
	calls = mock.calls.Transaction
	mock.lockTransaction.RUnlock()
	return calls
}

// UpdateProject calls UpdateProjectFunc.
func (mock *RepositoryMock) UpdateProject(ctx context.Context, p model.Project) error {
	if mock.UpdateProjectFunc == nil {
		panic("RepositoryMock.UpdateProjectFunc: method is nil but Repository.UpdateProject was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   model.Project
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockUpdateProject.Lock()
	mock.calls.UpdateProject = append(mock.calls.UpdateProject, callInfo)
	mock.lockUpdateProject.Unlock()
	return mock.UpdateProjectFunc(ctx, p)
}

// UpdateProjectCalls gets all the calls that were made to UpdateProject.
// Check the length with:
//
//	len(mockedRepository.UpdateProjectCalls())
func (mock *RepositoryMock) UpdateProjectCalls() []struct {
	Ctx context.Context
	P   model.Project
} {
	var calls []struct {
		Ctx context.Context
		P   model.Project
	}
	mock.lockUpdateProject.RLock()
	calls = mock.calls.UpdateProject
	mock.lockUpdateProject.RUnlock()
	return calls
}

// UpdateStep calls UpdateStepFunc.
func (mock *RepositoryMock) UpdateStep(ctx context.Context, s model.Step) error {
	if mock.UpdateStepFunc == nil {
		panic("RepositoryMock.UpdateStepFunc: method is nil but Repository.UpdateStep was just called")
	}
	callInfo := struct {
		Ctx context.Context
		S   model.Step
	}{
		Ctx: ctx,
		S:   s,
	}
	mock.lockUpdateStep.Lock()
	mock.calls.UpdateStep = append(mock.calls.UpdateStep, callInfo)
	mock.lockUpdateStep.Unlock()
	return mock.UpdateStepFunc(ctx, s)
}

// UpdateStepCalls gets all the calls that were made to UpdateStep.
// Check the length with:
//
//	len(mockedRepository.UpdateStepCalls())
func (mock *RepositoryMock) UpdateStepCalls() []struct {
	Ctx context.Context
	S   model.Step
} {
	var calls []struct {
		Ctx context.Context
		S   model.Step
	}
	mock.lockUpdateStep.RLock()
	calls = mock.calls.UpdateStep
	mock.lockUpdateStep.RUnlock()
	return calls
}

// UpdateTag calls UpdateTagFunc.
func (mock *RepositoryMock) UpdateTag(ctx context.Context, t model.Tag) error {
	if mock.UpdateTagFunc == nil {
		panic("RepositoryMock.UpdateTagFunc: method is nil but Repository.UpdateTag was just called")
	}
	callInfo := struct {
		Ctx context.Context
		T   model.Tag
	}{
		Ctx: ctx,
		T:   t,
	}
	mock.lockUpdateTag.Lock()
	mock.calls.UpdateTag = append(mock.calls.UpdateTag, callInfo)
	mock.lockUpdateTag.Unlock()
	return mock.UpdateTagFunc(ctx, t)
}

// UpdateTagCalls gets all the calls that were made to UpdateTag.
// Check the length with:
//
//	len(mockedRepository.UpdateTagCalls())
func (mock *RepositoryMock) UpdateTagCalls() []struct {
	Ctx context.Context
	T   model.Tag
} {
	var calls []struct {
		Ctx context.Context
		T   model.Tag
	}
	mock.lockUpdateTag.RLock()
	calls = mock.calls.UpdateTag
	mock.lockUpdateTag.RUnlock()
	return calls
}

// UpdateTask calls UpdateTaskFunc.
func (mock *RepositoryMock) UpdateTask(ctx context.Context, t model.Task) error {
	if mock.UpdateTaskFunc == nil {
		panic("RepositoryMock.UpdateTaskFunc: method is nil but Repository.UpdateTask was just called")
	}
	callInfo := struct {
		Ctx context.Context
		T   model.Task
	}{
		Ctx: ctx,
		T:   t,
	}
	mock.lockUpdateTask.Lock()
	mock.calls.UpdateTask = append(mock.calls.UpdateTask, callInfo)
	mock.lockUpdateTask.Unlock()
	return mock.UpdateTaskFunc(ctx, t)
}

// UpdateTaskCalls gets all the calls that were made to UpdateTask.
// Check the length with:
//
//	len(mockedRepository.UpdateTaskCalls())
func (mock *RepositoryMock) UpdateTaskCalls() []struct {
	Ctx context.Context
	T   model.Task
} {
	var calls []struct {
		Ctx context.Context
		T   model.Task
	}
	mock.lockUpdateTask.RLock()
	calls = mock.calls.UpdateTask
	mock.lockUpdateTask.RUnlock()
	return calls
}
