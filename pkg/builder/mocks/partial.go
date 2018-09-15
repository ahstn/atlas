package mocks

import (
	"fmt"
	"strings"

	mock "github.com/stretchr/testify/mock"
)

// BuilderPartial mock with the aim of retaining arg modification
type BuilderPartial struct {
	mock.Mock
	args []string
}

// Args provides a mock function with given fields:
func (_m *BuilderPartial) Args() string {
	return fmt.Sprintf("mvn --batch-mode %s", strings.Join(_m.args, " "))
}

// ModifyArgs provides a mock function with given fields: _a0
func (_m *BuilderPartial) ModifyArgs(a []string) {
	_m.args = a
}

// Run provides a mock function with given fields: _a0
func (_m *BuilderPartial) Run(_a0 bool) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
