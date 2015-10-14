package mocks

import "github.com/stretchr/testify/mock"

type Proxy struct {
	mock.Mock
}

func (_m *Proxy) Passthrough(s string) string {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
