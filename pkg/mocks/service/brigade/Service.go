// Code generated by mockery v1.0.0
package brigade

import brigade "github.com/Azure/brigade/pkg/brigade"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetBuild provides a mock function with given fields: buildID
func (_m *Service) GetBuild(buildID string) (*brigade.Build, error) {
	ret := _m.Called(buildID)

	var r0 *brigade.Build
	if rf, ok := ret.Get(0).(func(string) *brigade.Build); ok {
		r0 = rf(buildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(buildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildJobs provides a mock function with given fields: BuildID, desc
func (_m *Service) GetBuildJobs(BuildID string, desc bool) ([]*brigade.Job, error) {
	ret := _m.Called(BuildID, desc)

	var r0 []*brigade.Job
	if rf, ok := ret.Get(0).(func(string, bool) []*brigade.Job); ok {
		r0 = rf(BuildID, desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(BuildID, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: jobID
func (_m *Service) GetJob(jobID string) (*brigade.Job, error) {
	ret := _m.Called(jobID)

	var r0 *brigade.Job
	if rf, ok := ret.Get(0).(func(string) *brigade.Job); ok {
		r0 = rf(jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobLog provides a mock function with given fields: jobID
func (_m *Service) GetJobLog(jobID string) (string, error) {
	ret := _m.Called(jobID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(jobID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: projectID
func (_m *Service) GetProject(projectID string) (*brigade.Project, error) {
	ret := _m.Called(projectID)

	var r0 *brigade.Project
	if rf, ok := ret.Get(0).(func(string) *brigade.Project); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectBuilds provides a mock function with given fields: project, desc
func (_m *Service) GetProjectBuilds(project *brigade.Project, desc bool) ([]*brigade.Build, error) {
	ret := _m.Called(project, desc)

	var r0 []*brigade.Build
	if rf, ok := ret.Get(0).(func(*brigade.Project, bool) []*brigade.Build); ok {
		r0 = rf(project, desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Project, bool) error); ok {
		r1 = rf(project, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectLastBuild provides a mock function with given fields: projectID
func (_m *Service) GetProjectLastBuild(projectID string) (*brigade.Build, error) {
	ret := _m.Called(projectID)

	var r0 *brigade.Build
	if rf, ok := ret.Get(0).(func(string) *brigade.Build); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjects provides a mock function with given fields:
func (_m *Service) GetProjects() ([]*brigade.Project, error) {
	ret := _m.Called()

	var r0 []*brigade.Project
	if rf, ok := ret.Get(0).(func() []*brigade.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
