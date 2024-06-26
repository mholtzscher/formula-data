// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dal "github.com/mholtzscher/formula-data/internal/dal"
	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// CreateSeason provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateSeason(ctx context.Context, arg dal.CreateSeasonParams) (int32, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateSeason")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dal.CreateSeasonParams) (int32, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dal.CreateSeasonParams) int32); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dal.CreateSeasonParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_CreateSeason_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSeason'
type MockQuerier_CreateSeason_Call struct {
	*mock.Call
}

// CreateSeason is a helper method to define mock.On call
//   - ctx context.Context
//   - arg dal.CreateSeasonParams
func (_e *MockQuerier_Expecter) CreateSeason(ctx interface{}, arg interface{}) *MockQuerier_CreateSeason_Call {
	return &MockQuerier_CreateSeason_Call{Call: _e.mock.On("CreateSeason", ctx, arg)}
}

func (_c *MockQuerier_CreateSeason_Call) Run(run func(ctx context.Context, arg dal.CreateSeasonParams)) *MockQuerier_CreateSeason_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dal.CreateSeasonParams))
	})
	return _c
}

func (_c *MockQuerier_CreateSeason_Call) Return(_a0 int32, _a1 error) *MockQuerier_CreateSeason_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_CreateSeason_Call) RunAndReturn(run func(context.Context, dal.CreateSeasonParams) (int32, error)) *MockQuerier_CreateSeason_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllSeasons provides a mock function with given fields: ctx
func (_m *MockQuerier) GetAllSeasons(ctx context.Context) ([]dal.Season, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSeasons")
	}

	var r0 []dal.Season
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dal.Season, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dal.Season); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dal.Season)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetAllSeasons_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllSeasons'
type MockQuerier_GetAllSeasons_Call struct {
	*mock.Call
}

// GetAllSeasons is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) GetAllSeasons(ctx interface{}) *MockQuerier_GetAllSeasons_Call {
	return &MockQuerier_GetAllSeasons_Call{Call: _e.mock.On("GetAllSeasons", ctx)}
}

func (_c *MockQuerier_GetAllSeasons_Call) Run(run func(ctx context.Context)) *MockQuerier_GetAllSeasons_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_GetAllSeasons_Call) Return(_a0 []dal.Season, _a1 error) *MockQuerier_GetAllSeasons_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetAllSeasons_Call) RunAndReturn(run func(context.Context) ([]dal.Season, error)) *MockQuerier_GetAllSeasons_Call {
	_c.Call.Return(run)
	return _c
}

// GetDriver provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetDriver(ctx context.Context, id int32) (dal.Driver, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetDriver")
	}

	var r0 dal.Driver
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (dal.Driver, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) dal.Driver); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dal.Driver)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDriver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDriver'
type MockQuerier_GetDriver_Call struct {
	*mock.Call
}

// GetDriver is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerier_Expecter) GetDriver(ctx interface{}, id interface{}) *MockQuerier_GetDriver_Call {
	return &MockQuerier_GetDriver_Call{Call: _e.mock.On("GetDriver", ctx, id)}
}

func (_c *MockQuerier_GetDriver_Call) Run(run func(ctx context.Context, id int32)) *MockQuerier_GetDriver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerier_GetDriver_Call) Return(_a0 dal.Driver, _a1 error) *MockQuerier_GetDriver_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDriver_Call) RunAndReturn(run func(context.Context, int32) (dal.Driver, error)) *MockQuerier_GetDriver_Call {
	_c.Call.Return(run)
	return _c
}

// GetRace provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetRace(ctx context.Context, id int32) (dal.Race, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetRace")
	}

	var r0 dal.Race
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (dal.Race, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) dal.Race); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dal.Race)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetRace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRace'
type MockQuerier_GetRace_Call struct {
	*mock.Call
}

// GetRace is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerier_Expecter) GetRace(ctx interface{}, id interface{}) *MockQuerier_GetRace_Call {
	return &MockQuerier_GetRace_Call{Call: _e.mock.On("GetRace", ctx, id)}
}

func (_c *MockQuerier_GetRace_Call) Run(run func(ctx context.Context, id int32)) *MockQuerier_GetRace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerier_GetRace_Call) Return(_a0 dal.Race, _a1 error) *MockQuerier_GetRace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetRace_Call) RunAndReturn(run func(context.Context, int32) (dal.Race, error)) *MockQuerier_GetRace_Call {
	_c.Call.Return(run)
	return _c
}

// GetResult provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetResult(ctx context.Context, id int32) (dal.Result, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetResult")
	}

	var r0 dal.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (dal.Result, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) dal.Result); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dal.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResult'
type MockQuerier_GetResult_Call struct {
	*mock.Call
}

// GetResult is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerier_Expecter) GetResult(ctx interface{}, id interface{}) *MockQuerier_GetResult_Call {
	return &MockQuerier_GetResult_Call{Call: _e.mock.On("GetResult", ctx, id)}
}

func (_c *MockQuerier_GetResult_Call) Run(run func(ctx context.Context, id int32)) *MockQuerier_GetResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerier_GetResult_Call) Return(_a0 dal.Result, _a1 error) *MockQuerier_GetResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetResult_Call) RunAndReturn(run func(context.Context, int32) (dal.Result, error)) *MockQuerier_GetResult_Call {
	_c.Call.Return(run)
	return _c
}

// GetSeasonById provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetSeasonById(ctx context.Context, id int32) (dal.Season, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSeasonById")
	}

	var r0 dal.Season
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (dal.Season, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) dal.Season); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dal.Season)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetSeasonById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSeasonById'
type MockQuerier_GetSeasonById_Call struct {
	*mock.Call
}

// GetSeasonById is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerier_Expecter) GetSeasonById(ctx interface{}, id interface{}) *MockQuerier_GetSeasonById_Call {
	return &MockQuerier_GetSeasonById_Call{Call: _e.mock.On("GetSeasonById", ctx, id)}
}

func (_c *MockQuerier_GetSeasonById_Call) Run(run func(ctx context.Context, id int32)) *MockQuerier_GetSeasonById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerier_GetSeasonById_Call) Return(_a0 dal.Season, _a1 error) *MockQuerier_GetSeasonById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetSeasonById_Call) RunAndReturn(run func(context.Context, int32) (dal.Season, error)) *MockQuerier_GetSeasonById_Call {
	_c.Call.Return(run)
	return _c
}

// GetTeam provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetTeam(ctx context.Context, id int32) (dal.Team, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTeam")
	}

	var r0 dal.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (dal.Team, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) dal.Team); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dal.Team)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetTeam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTeam'
type MockQuerier_GetTeam_Call struct {
	*mock.Call
}

// GetTeam is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerier_Expecter) GetTeam(ctx interface{}, id interface{}) *MockQuerier_GetTeam_Call {
	return &MockQuerier_GetTeam_Call{Call: _e.mock.On("GetTeam", ctx, id)}
}

func (_c *MockQuerier_GetTeam_Call) Run(run func(ctx context.Context, id int32)) *MockQuerier_GetTeam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerier_GetTeam_Call) Return(_a0 dal.Team, _a1 error) *MockQuerier_GetTeam_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetTeam_Call) RunAndReturn(run func(context.Context, int32) (dal.Team, error)) *MockQuerier_GetTeam_Call {
	_c.Call.Return(run)
	return _c
}

// ListTeams provides a mock function with given fields: ctx
func (_m *MockQuerier) ListTeams(ctx context.Context) ([]dal.Team, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListTeams")
	}

	var r0 []dal.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dal.Team, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dal.Team); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dal.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_ListTeams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTeams'
type MockQuerier_ListTeams_Call struct {
	*mock.Call
}

// ListTeams is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) ListTeams(ctx interface{}) *MockQuerier_ListTeams_Call {
	return &MockQuerier_ListTeams_Call{Call: _e.mock.On("ListTeams", ctx)}
}

func (_c *MockQuerier_ListTeams_Call) Run(run func(ctx context.Context)) *MockQuerier_ListTeams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_ListTeams_Call) Return(_a0 []dal.Team, _a1 error) *MockQuerier_ListTeams_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_ListTeams_Call) RunAndReturn(run func(context.Context) ([]dal.Team, error)) *MockQuerier_ListTeams_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
