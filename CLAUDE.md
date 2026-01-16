# CLAUDE.md

本文件用于指导 Claude Code (claude.ai/code) 在此代码库中的工作。

**重要：在此项目中与用户交流请使用中文。**

## 构建与开发命令

```bash
make run      # 直接启动应用
make build    # 完整构建：生成 Wire 代码 + 编译到 ./bin/app
make wire     # 仅生成 Wire 依赖注入代码
go test ./config/...  # 运行测试（可添加其他包路径）
```

启动应用：构建后运行 `./bin/app`，或开发时使用 `make run`。

## 架构概览

这是一个基于 Go 的轻量级博客平台，采用标准三层架构（biz/data/model），使用 Gin 进行 HTTP 路由，Wire 进行编译期依赖注入，Viper 进行配置管理。

### 依赖注入 (Wire)

项目使用 Google Wire 进行编译期依赖注入。关键模式：

- 使用 `//go:build wireinject` 构建标记定义 providers
- 在每个包的 `provider.go` 中通过 `wire.NewSet()` 导出 `ProviderSet`
- 生成的代码位于 `wire_gen.go`（请勿编辑）
- 修改 providers 后运行 `make wire` 重新生成

主注入器：`internal/provider.go`（构建标记：`wireinject`）

### 三层架构

```
internal/
├── biz/        # 业务逻辑层 - Service
├── data/       # 数据访问层 - Repository/DAO
├── model/      # 数据模型 - Entity/DTO
├── http/       # HTTP 层 - Controller/Handler
│   ├── middleware/   # 中间件 (Recovery, Logger, Tracer)
│   ├── api/          # API 路由组
│   └── admin/        # Admin 路由组
└── job/        # 后台任务
```

### HTTP 层结构

每个路由组（api/admin）包含三个文件：
- `handler.go` - Handler 函数定义
- `router.go` - 路由注册逻辑（使用 `router.New*Router()` 辅助函数）
- `provider.go` - Wire ProviderSet

**注意**：`internal/http/provider.go` 的 `RouterProviderSet` 只包含 `NewRouterProvider`，不应包含 api/admin 的 ProviderSet（避免 Wire 绑定冲突）。

### 应用初始化流程

```
main.go → internal.InitApp() → Wire 构建依赖 → App.Run()
     ↓
Gin Engine + 基础中间件 → 路由组 → Handler → Biz → Data
```

基础中间件（全局应用，按顺序）：
1. Recovery - 捕获 panic，记录到 runtime/{name}-error-YYYY-MM-DD.log
2. Logger - 访问日志记录到 runtime/{name}-access-YYYY-MM-DD.log
3. Tracer - 生成/传递 X-TRACE-ID 请求头

### 添加新功能

1. `model/` - 定义数据结构
2. `data/` - 实现数据访问（database、cache 等）
3. `biz/` - 实现业务逻辑，依赖 data 层
4. `http/` - 定义 Handler，依赖 biz 层，在 router.go 中注册路由
5. 修改对应 `provider.go` 添加 ProviderSet
6. 运行 `make wire` 重新生成依赖注入代码

### 后台任务

通过 `job.ProviderSet` 注册的任务使用 `sync.WaitGroup` 并发运行。自定义任务需实现 `job.IJob` 接口。任务支持基于上下文的优雅关闭。

### 配置

运行时配置通过 Viper 从 `config/app.yml` 加载。关键路径通过 `config.AppConfig` 结构体访问。

### 当前路由

- `GET /api/health` - API 健康检查
- `GET /admin/health` - 管理后台健康检查
- `GET /admin/demoPanic` - 演示 panic 恢复中间件
