# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**重要：在此项目中与用户交流请使用中文。**

## 构建与开发命令

```bash
make run      # 直接启动应用 (go run ./cmd)
make build    # 完整构建：生成 Wire 代码 + 编译到 ./bin/app
make wire     # 仅生成 Wire 依赖注入代码
go test ./... # 运行所有测试
```

## 架构概览

基于 Go 的轻量级博客平台，三层架构 + 依赖注入 + 路由接口模式。

### 目录结构

```
internal/
├── biz/        # 业务逻辑层（Service）
├── data/       # 数据访问层（Repository/DAO）
├── model/      # 数据模型（Entity/DTO）
├── http/       # HTTP 层
│   ├── middleware/   # 中间件 (Recovery, Logger, Tracer)
│   ├── api/          # API 路由组 (handler.go, router.go, provider.go)
│   ├── admin/        # Admin 路由组
│   │   └── user/     # 用户模块 (UserController)
│   └── provider.go   # 路由聚合 (RouterProviderSet)
└── job/        # 后台任务 (IJob, IJobGroup)
```

### 依赖注入 (Wire)

- 主注入器：`internal/provider.go`（标记：`//go:build wireinject`）
- 各包 `provider.go` 通过 `wire.NewSet()` 导出 `ProviderSet`
- `make wire` 重新生成注入代码

### 路由注册模式

使用 `router` 包的接口模式注册路由：

```go
// router/IRouter: 单个路由
router.NewGetRouter("/path", handler)  // 自动命名
router.NewRouter("POST", "/path", "name", handler)  // 自定义命名

// router/IRouterGroup: 路由组
api.NewApiRouterProvider()      // 返回 []router.IRouterGroup
admin.NewAdminRouterProvider()  // 注入 UserController 等依赖
```

**注意**：`internal/http/provider.go` 的 `RouterProviderSet` 只聚合各路由组，不应包含具体 Handler 的 ProviderSet。

### 应用初始化

```
main.go → InitApp() → Wire 构建 → App.Run()
         ↓
Gin Engine + 基础中间件 → 路由组 → Handler → Biz → Data
```

基础中间件（顺序）：
1. Recovery - 捕获 panic，写入 `runtime/{name}-error-YYYY-MM-DD.log`
2. Logger - 访问日志，写入 `runtime/{name}-access-YYYY-MM-DD.log`
3. Tracer - 生成/传递 `X-TRACE-ID`

`App` 结构体可通过 `context.Set("app", app)` 获取，内置 `GetLogger()` 方法。

### 后台任务

- `job.IJob` 接口：`Run(ctx context.Context) error` + `Name() string`
- `job.IJobGroup` 聚合任务，使用 `sync.WaitGroup` 并发执行
- 通过上下文实现优雅关闭
- 在 `job/provider.go` 的 `NewJobGroupProvider()` 中注册

### 配置

- 入口：`config/app.yml`
- 结构体：`config.AppConfig`（Name, Version, Port, Host, ErrorPath, Env）
- 解析：`config.NewAppConfig()` 使用 Viper

### 添加新功能步骤

1. `model/` - 定义数据结构
2. `data/` - 实现数据访问
3. `biz/` - 业务逻辑，依赖 data 层
4. `http/admin/user/` - 创建模块目录，定义 Controller
5. `internal/http/admin/provider.go` - 添加 Controller 和 RouterProvider 到 ProviderSet
6. `make wire` - 重新生成依赖注入
