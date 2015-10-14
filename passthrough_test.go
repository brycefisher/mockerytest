package test

import (
	"github.com/brycefisher/mockerytest/mocks"
	"testing"
)

func TestReturnValueProviders(t *testing.T) {
	p := new(mocks.Proxy)
	p.On("Passthrough").Return(func(s string) string {
		return s
	})
	ret := p.Passthrough("hi")
	if ret != "hi" {
		t.Errorf("Expected \"hi\" but found \"%s\"", ret)
	}
}
