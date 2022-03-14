package conf

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_NewPkgTestConfService(t *testing.T) {
    conf := NewPkgTestConfService()
    assert.NotNil(t, conf)
    assert.NotNil(t, conf.Service)
    assert.NotNil(t, conf.Server)
    assert.NotNil(t, conf.Service.Data)
    assert.NotNil(t, conf.Service.Data.Database)
    assert.NotNil(t, conf.Service.Data.Redis)
}
