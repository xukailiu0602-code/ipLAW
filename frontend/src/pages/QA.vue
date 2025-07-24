<template>
  <a-card title="智能问答">
    <a-form @submit.prevent="onAsk">
      <a-form-item label="问题">
        <a-input v-model:value="query" placeholder="请输入您的知识产权问题" />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit" :loading="loading">提交</a-button>
      </a-form-item>
    </a-form>
    <div v-if="answer">
      <a-divider>AI 回答</a-divider>
      <div>{{answer}}</div>
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
  </a-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { ask } from '../api/qa'
const query = ref('')
const answer = ref('')
const citations = ref<string[]>([])
const laws = ref<any[]>([])
const cases = ref<any[]>([])
const riskScore = ref(0)
const rationale = ref('')
const loading = ref(false)
const onAsk = async () => {
  loading.value = true
  const res = await ask({ query: query.value })
  answer.value = res.answer
  citations.value = res.citations
  laws.value = res.laws
  cases.value = res.cases
  riskScore.value = res.risk_score
  rationale.value = res.rationale
  loading.value = false
}
</script>