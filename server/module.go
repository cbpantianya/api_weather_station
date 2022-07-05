package server

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
	"sync"
)

var (
	modules   = make(map[string]ModuleInfo)
	modulesMu sync.RWMutex
)

type Module interface {
	GetModuleInfo() ModuleInfo
	// Module 的生命周期

	// Init 初始化
	// 待所有 Module 初始化完成后
	// 进行服务注册 Serve
	Init()

	// Serve 向Bot注册服务函数
	// 结束后调用 Start
	Serve(server *Server)
}

// RegisterModule - 向全局添加 Module
func RegisterModule(instance Module) {
	mod := instance.GetModuleInfo()

	if mod.ID == "" {
		panic("module ID missing")
	}
	if mod.Instance == nil {
		panic("missing ModuleInfo.Instance")
	}

	modulesMu.Lock()
	defer modulesMu.Unlock()

	if _, ok := modules[string(mod.ID)]; ok {
		panic(fmt.Sprintf("module already registered: %s", mod.ID))
	}
	modules[string(mod.ID)] = mod
}

// ModuleID 模块ID
// 请使用 小写 并用 _ 代替空格
// Example:
// - atom.pong
type ModuleID string

// Namespace - 获取一个 Module 的 Namespace
func (id ModuleID) Namespace() string {
	lastDot := strings.LastIndex(string(id), ".")
	if lastDot < 0 {
		return ""
	}
	return string(id)[:lastDot]
}

// Name - 获取一个 Module 的 Name
func (id ModuleID) Name() string {
	if id == "" {
		return ""
	}
	parts := strings.Split(string(id), ".")
	return parts[len(parts)-1]
}

type ModuleInfo struct {
	// ID 模块的名称
	// 应全局唯一
	ID ModuleID

	// Instance 返回 Module
	Instance Module
}

func (mi ModuleInfo) String() string {
	return string(mi.ID)
}

func StartService() {
	for _, mi := range modules {
		mi.Instance.Init()
	}

	for _, mi := range modules {
		mi.Instance.Serve(Instance)
	}

	log.Logger.Info().Msg("Tasks started")
}
