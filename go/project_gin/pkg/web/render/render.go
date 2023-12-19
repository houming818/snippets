package render

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Render is the interface that wraps the basic Render method.

type Render interface {
	Setup(c *gin.Context) error
	GetDataH(c *gin.Context) (gin.H, error)
}

// typeRegistry is the registry of all registered render types.
var typeRegistry = map[string]reflect.Type{
	"console": reflect.TypeOf(ConsoleRender{}),
	"debug":   reflect.TypeOf(DebugRender{}),
}

func RenderFactory(typeName string, ctx *gin.Context) (Render, error) {
	// 通过类型名称创建实例
	if _, ok := typeRegistry[typeName]; !ok {
		return nil, fmt.Errorf("function %v not found", typeName)
	}
	objType := typeRegistry[typeName]
	objValue := reflect.New(objType)

	objI := objValue.Interface().(Render)

	// 调用实例的Setup方法
	err := objI.Setup(ctx)
	if err != nil {
		fmt.Errorf("function %v setup failed", typeName)
		return nil, fmt.Errorf("action %v setup failed", typeName)
	}
	return objI, nil
}

type ConsoleRender struct {
	Render
}

func (c *ConsoleRender) Setup(ctx *gin.Context) error {
	return nil
}

func (c *ConsoleRender) GetDataH(ctx *gin.Context) (gin.H, error) {
	fmt.Println("console")
	return gin.H{"title": "console"}, nil
}

type DebugRender struct {
	Render
}

func (c *DebugRender) Setup(ctx *gin.Context) error {
	return nil
}

func (c *DebugRender) GetDataH(ctx *gin.Context) (gin.H, error) {
	return gin.H{"title": "debugaaa"}, nil
}
