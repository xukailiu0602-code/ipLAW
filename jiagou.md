# 知识产权侵权风险评估平台 启动流程说明

## 一、环境依赖

- Node.js >= 16（推荐 18+）
- npm >= 8
- Go >= 1.20
- MongoDB（本地或远程）
- Milvus（或 Redis Vector，推荐 Milvus）

## 二、后端服务启动（Go）

1. 进入后端目录

```bash
cd backend
```

2. 配置数据库和向量库（可修改 `config/config.yaml`）

3. 启动后端服务

```bash
go run main.go
```

- 默认监听端口：`8080`
- API 路径前缀：`/api`

4. 可选：编译二进制

```bash
go build -o iprisk-server main.go
./iprisk-server
```

## 三、前端服务启动（Vite）

1. 进入前端目录

```bash
cd frontend
```

2. 安装依赖

```bash
npm install
```

3. 启动开发服务器

```bash
npm run dev
```

- 默认监听端口：`5173`
- 访问地址：http://localhost:5173
- 已配置 API 代理到后端 `/api`

## 四、测试流程

1. 浏览器访问 http://localhost:5173
2. 注册/登录账号
3. 上传文档、发起问答、导出报告、进入后台管理

## 五、常见问题

- 若端口冲突，可在 `.env` 或 `vite.config.ts`/`main.go` 中修改端口
- MongoDB/Milvus 未启动会导致后端报错，请先启动依赖服务
- 前后端均支持热更新，修改代码后自动生效

---

如需生产部署，建议使用 Docker Compose 或 K8s，详见 `deploy/` 目录。