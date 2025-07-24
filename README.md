你是资深产品架构与 AI 工程专家，请帮助我设计“知识产权侵权风险评估平台”的 AI 模块和整体架构。要求技术栈如下：

- 前端：Vue3 + Vite + Ant Design Vue / Element Plus
- 后端：Go + Gin 或 Echo
- 存储：MongoDB（原文、类案、问答记录等）
- AI 问答：现代 RAG 流程，使用千问 API 服务（Qwen3‑Embedding、Qwen3‑Reranker、Qwen‑Plus）
- Orchestration：采用 ByteDance 开源 Go 框架 Eino 进行节点编排

---

## 一、平台愿景与交互路径

- 支持**文档上传**（PDF/DOCX/TXT）+问答，也支持**纯文本输入提问**
- 输出层次包括：
  - **answer**：简洁专业回答
  - **citations**：引用上下文片段
  - **laws**：关联法律条款解析
  - **cases**：历史类案匹配列表
  - **risk_score**：0–100 风险评分
  - **rationale**：评分解释
- 支持报告导出（PDF/Word），包含问答记录、法条解析、类案列表、评分说明
- 后台管理支持：模型配置、切片规则、索引重建、权限控制（访客/评估者/审核者/管理员）

---

## 二、三阶段 RAG AI 流程

1. **Embedding 召回阶段**  
   - 使用 Qwen3‑Embedding 模型（0.6B / 4B / 8B）生成文本向量，参考其多语言泛化与定制维度能力 1  
   - 文档切片后存入向量库（推荐 Milvus / Weaviate / Redis Vector）

2. **Reranker 精排阶段**  
   - 使用 Qwen3‑Reranker 模型对 top‑50 召回结果精排，选出 top‑10 最相关段落做到高相关性 2

3. **生成阶段（Qwen‑Plus）**  
   - 使用 DashScope 或 OpenAI 接口的 Qwen‑Plus 模型，支持 131k 上下文长度，可生成专业级回答 3

### 输出 JSON 格式

```json
{
  "query": "...",
  "contexts": [{"id":"C1","text":"..."}, ...],
  "answer": "...",
  "citations":["C1","C3"],
  "laws":[{"article":"著作权法第**条","description":"..."}],
  "cases":[{"caseId":"2022XYZ","court":"北京知识产权法院","year":2023,"summary":"...","similarity":0.85}, ...],
  "risk_score":85,
  "rationale":"因类案匹配数量 ≥3，且相似度均 ≥0.8，风险较高……"
}


---

三、Eino Orchestration 编排模型（Go）

graph := NewGraph[map[string]any, EvalResult]()
graph.AddNode("embed_retrieve", retrieverNode)     // Qwen3‑Embedding
graph.AddNode("rerank", rerankerNode)              // Qwen3‑Reranker
graph.AddNode("llm_answer", chatModelNode)         // Qwen‑Plus
graph.AddNode("post_process", postProcessNode)     // 构建 JSON + 风险评分 + 类案检索

graph.AddEdge(START, "embed_retrieve")
graph.AddEdge("embed_retrieve", "rerank")
graph.AddEdge("rerank", "llm_answer")
graph.AddEdge("llm_answer", "post_process")
graph.AddEdge("post_process", END)

res := runner.Invoke(ctx, {"query": question, "history": history, "userId": uid})

RetrieverNode 调用 Qwen3‑Embedding API

RerankerNode 调用 Qwen3‑Reranker API

ChatModelNode 调用 Qwen‑Plus API 服务  

PostProcessNode 计算风险、扩展类案，整理结构 JSON



---

四、核心功能模块设计

4.1 文档上传与切片

支持 PDF/DOCX/TXT，抽取文本后按语义段、法律条款分切片

每个切片做 embedding，存向量库；MongoDB 存原文、id、meta 信息


4.2 RAG 问答交互

用户提问或上传文档提问

启动 AI 流程：Embedding → Reranker → Qwen‑Plus → PostProcess

持久化记录问答历史、评分、类案、解析等内容


4.3 类案检索

判例库存插入每案向量和元数据

PostProcessNode 对上下文片段发起相似度检索（阈值如 ≥0.8）

向前端返回案号、法院、时间、相似度及概要、支持点击查看原文


4.4 知识点解析

Qwen‑Plus 输出 laws 数组，含法条编号与解释

前端树状展示：法条 → 判例支持 → 注意事项，可展开细看


4.5 风险评分与报告导出

评分策略示例：

类案 ≥3 且相似度 ≥0.8 → 高风险（80–100）

类案 1–2 或中等相似度 → 中风险（40–79）

类案缺失 → 低风险（0–39）


报告导出格式：PDF（GoFPDF）或 Word（Unidoc）


4.6 后台管理与权限系统

UI：模型配置、切片规则管理、索引重建、日志监控

权限：访客、评估者、审核者、管理员

支持监控询问次数、LLM 调用次数、评分准确度等



---

五、部署与非功能需求

容器化：Docker Compose 或 Kubernetes 部署 Go 服务、向量库、MongoDB、前端

CI/CD：GitHub Actions + 单元测试 + 接口测试 + 代码质量扫描

监控：OpenTelemetry + Langfuse 跟踪 Eino 各节点指标、错误率、延迟

性能目标：单次问答延迟 <1s，系统支持千 QPS

可扩展性：支持扩展图文 OCR、音视频输入、外语类案、多模型融合等



---

六、Prompt 调用示例

你是法律 AI 模型，输入如下：
{
 "query":"{{user_question}}",
 "contexts":[{"id":"C1","text":"..."},...],
 "maxLaws":3,
 "maxCases":5
}
请输出 JSON，字段 keys：answer, citations, laws, cases, risk_score, rationale。


---

七、方案亮点总结

千问三段式流程：召回（embedding）→ 精排（reranker）→ 高质量生成（qwen‑plus），可控高效  

Go/Eino Orchestration：强类型、流式处理、组件化、高并发，源自 ByteDance 实战  

结构化输出：法律知识解析、类案支持、评分机制，便于解释和业务落地

全链路支持：前端交互、后端服务、后台管理、报告生成与导出、CI/CD、监控
