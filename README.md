# 知识产权侵权风险评估平台 接口文档

## 一、RAG材料存放说明

- **原始材料文件**：
  - 推荐存放于后端 `/backend/data/` 目录（如：/backend/data/laws.json, /backend/data/cases.json, /backend/data/documents/）
  - 支持 txt、json、pdf 等格式，建议结构化为 json 便于入库
- **数据库存储**：
  - 文本原文、元数据存 MongoDB
  - 向量（embedding）存 Milvus
- **上传入口**：
  - 前端“文档上传”页面（/upload），对应后端 `/api/upload` 接口

---

## 二、前后端接口一览

| 前端页面/功能         | 前端API调用           | 后端接口路径      | 主要后端处理逻辑/依赖         |
|----------------------|----------------------|------------------|------------------------------|
| 登录                 | `auth.ts: login`     | `/api/login`     | 用户认证，JWT                |
| 注册                 | `auth.ts: register`  | `/api/register`  | 用户注册，写入用户表         |
| 智能问答             | `qa.ts: ask`         | `/api/ask`       | RAG主流程，见下详细说明      |
| 文档上传             | `document.ts: upload`| `/api/upload`    | 材料入库，embedding          |
| 法条查询             | `law.ts: getLaws`    | `/api/law`       | MongoDB法条表                |
| 案例查询             | `case.ts: getCases`  | `/api/case`      | MongoDB案例表                |
| 报告导出             | `report.ts: getReport`| `/api/report/:id`| 生成/导出PDF报告             |
| 后台管理             | `admin.ts: ...`      | `/api/admin/...` | 用户/材料/日志管理           |

---

## 三、智能问答（RAG）接口详细说明

### 1. 前端调用

- 文件：`frontend/src/api/qa.ts`
- 方法：`ask({ query: string, doc_ids?: string[], history?: any[] })`
- 页面：`frontend/src/pages/QA.vue`（主页面）、`QAForm.vue`（输入组件）

### 2. 后端接口

- 路径：`/api/ask`（POST）
- 文件：`backend/api/qa.go`，方法：`Ask`
- 主要流程：
  1. 解析用户query
  2. 通过RAG编排（`eino/rag_graph.go`）依次调用：
     - `RetrieverNode`：调用千问embedding，Milvus召回相关材料
     - `RerankerNode`：调用千问rerank，精排top文档
     - `LLMAnswerNode`：调用千问LLM，生成最终答案
  3. 返回结构：
     ```json
     {
       "answer": "AI生成的答案",
       "citations": ["引用片段1", "引用片段2"],
       "laws": [{"article": "法条编号", "description": "法条内容"}],
       "cases": [{"caseId": "案例ID", "court": "法院", "year": "年份", "similarity": 0.92}],
       "risk_score": 85,
       "rationale": "风险评分理由"
     }
     ```
- 依赖：
  - `services/qwen.go`：千问API（embedding、rerank、llm）
  - `services/milvus.go`：向量召回
  - `models/`：如有自定义数据结构

### 3. RAG材料流转

- 上传/导入：前端上传 -> 后端 `/api/upload` -> 入MongoDB/Milvus
- 检索：问答时 embedding 检索 -> Milvus 召回 -> rerank -> LLM生成
- 存放目录建议：
  - `/backend/data/`：原始材料文件
  - MongoDB：结构化文本、元数据
  - Milvus：embedding向量

---

## 四、接口字段与数据结构

### 1. 问答请求
```json
POST /api/ask
{
  "query": "专利侵权如何认定？",
  "doc_ids": ["..."],
  "history": [ ... ]
}
```

### 2. 问答响应
```json
{
  "answer": "...",
  "citations": ["..."],
  "laws": [{"article": "...", "description": "..."}],
  "cases": [{"caseId": "...", "court": "...", "year": "...", "similarity": 0.91}],
  "risk_score": 80,
  "rationale": "..."
}
```

---

## 五、接口扩展建议

- 支持批量材料导入、异步embedding、材料分组管理
- 支持多轮对话（history字段）
- 支持自定义RAG召回策略

---

如需更详细的API参数、RAG材料格式或数据库结构示例，请联系开发团队。
