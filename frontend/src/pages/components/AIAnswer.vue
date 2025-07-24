<template>
  <div class="ai-answer-area">
    <a-divider>AI 回答</a-divider>
    <div class="answer-main">{{answer}}</div>
    <a-divider>引用片段</a-divider>
    <ul>
      <li v-for="c in citations" :key="c">{{c}}</li>
    </ul>
    <a-divider>关联法条</a-divider>
    <ul>
      <li v-for="l in laws" :key="l.article">{{l.article}}: {{l.description}}</li>
    </ul>
    <a-divider>历史类案</a-divider>
    <ul>
      <li v-for="cs in cases" :key="cs.caseId">{{cs.caseId}} {{cs.court}} {{cs.year}} 相似度:{{cs.similarity}}</li>
    </ul>
    <a-divider>风险评分</a-divider>
    <div>风险分: {{riskScore}} ({{rationale}})</div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{
  answer: string
  citations: string[]
  laws: { article: string; description: string }[]
  cases: { caseId: string; court: string; year: string; similarity: number }[]
  riskScore: number
  rationale: string
}>()
</script>
<style scoped>
.ai-answer-area {
  margin-top: 32px;
  background: linear-gradient(120deg, #f0f4ff 60%, #e0e7ff 100%);
  border-radius: 14px;
  box-shadow: 0 2px 12px 0 rgba(46,76,255,0.06);
  padding: 24px 24px 12px 24px;
  min-height: 120px;
}
.answer-main {
  font-size: 1.18rem;
  font-weight: 500;
  color: #222;
  line-height: 1.8;
  margin-bottom: 18px;
  opacity: 0;
}
</style>