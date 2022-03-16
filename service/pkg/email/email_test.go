package email

import (
	"github.com/stretchr/testify/assert"
	"service/pkg/conf"
	"testing"
)

func TestNewEmailService(t *testing.T) {
	s, err := NewEmailService(&conf.Email{
		User: "mail@mailpush.thjam.cc",
		Pass: "iTvLd6f9cKwQ3yn",
		Host: "smtpdm.aliyun.com",
		Port: 80,
	})
	if err != nil {
		return
	}
	err = s.SendEmail("null122", "测试", "测试测试测试测试", "1092570726@qq.com")
	assert.NoError(t, err)
}
