package config

import (
	"io/ioutil"
	"testing"
)

func TestNewCache(t *testing.T) {
	// Given
	//ctrl := gomock.NewController(t)

	// 写入Config配置文件到临时文件
	t.Log("写入Config配置文件到临时文件 /tmp/app_cache_gocache.yaml")
	configData := `
http:
  host: 0.0.0.0:5277
  debug: true

cache:
  backend: gocache
  expire: 60
  cleanup: 90
`

	ioutil.WriteFile("/tmp/app_cache_gocache.yaml", []byte(configData), 0644)
	configPath := "/tmp/app_cache_gocache.yaml"
	_, err := New(&configPath)

	if err != nil {
		t.Log("读取配置文件失败")
		t.Fail()
	}

	t.Log("读取配置文件成功")

}
