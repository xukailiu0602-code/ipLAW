<template>
  <div class="qa-bg">
    <motion-div class="qa-banner" :initial="{opacity:0, y:-60}" :enter="{opacity:1, y:0, transition:{duration:1.2, type:'spring'}}">
      <div class="banner-title">知识产权智能问答</div>
      <div class="banner-desc">AI赋能 · 法律合规 · 智能风控</div>
    </motion-div>
    <motion-div class="qa-card-wrap" :initial="{opacity:0, scale:0.95}" :enter="{opacity:1, scale:1, transition:{duration:0.8}}">
      <a-card class="qa-card" title="智能问答">
        <a-form @submit.prevent="onAsk">
          <a-form-item label="问题">
            <a-input v-model:value="query" placeholder="请输入您的知识产权问题" />
          </a-form-item>
          <a-form-item>
            <a-button type="primary" html-type="submit" :loading="loading" class="ask-btn">
              <span v-if="!loading">提交</span>
              <span v-else class="dot-flashing"></span>
            </a-button>
          </a-form-item>
        </a-form>
        <transition name="fade-slide">
          <div v-if="answer" class="ai-answer-area">
            <a-divider>AI 回答</a-divider>
            <motion-div :initial="{opacity:0, y:30}" :enter="{opacity:1, y:0, transition:{duration:0.7}}">
              <div class="answer-main">{{answer}}</div>
            </motion-div>
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
        </transition>
      </a-card>
    </motion-div>
  </div>
</template>
<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { ask } from '../api/qa'
import { useMotion } from '@vueuse/motion'
import anime from 'animejs'
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
  await nextTick()
  anime({ targets: '.answer-main', opacity: [0,1], translateY: [40,0], duration: 800, easing: 'easeOutExpo' })
}
</script>
<style scoped>
.qa-bg {
  min-height: 100vh;
  background: linear-gradient(120deg, #e0e7ff 0%, #f0f4ff 100%);
  padding-bottom: 60px;
}
.qa-banner {
  width: 100vw;
  height: 220px;
  background: linear-gradient(90deg, #2e4cff 0%, #00c6fb 100%);
  color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 32px 0 rgba(46,76,255,0.12);
  margin-bottom: -80px;
  z-index: 2;
  position: relative;
}
.banner-title {
  font-size: 2.6rem;
  font-weight: 700;
  letter-spacing: 2px;
  margin-bottom: 10px;
}
.banner-desc {
  font-size: 1.2rem;
  opacity: 0.85;
}
.qa-card-wrap {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 500px;
  padding-top: 0;
  z-index: 3;
  position: relative;
}
.qa-card {
  width: 600px;
  margin-top: 100px;
  border-radius: 18px;
  box-shadow: 0 8px 32px 0 rgba(46,76,255,0.10);
  background: linear-gradient(120deg, #fff 60%, #f0f4ff 100%);
  border: none;
}
.ask-btn {
  width: 120px;
  height: 40px;
  font-size: 1.1rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px 0 rgba(46,76,255,0.08);
  transition: box-shadow 0.3s;
}
.ask-btn:hover {
  box-shadow: 0 4px 16px 0 rgba(46,76,255,0.16);
}
.dot-flashing {
  position: relative;
  width: 1.5em;
  height: 1em;
}
.dot-flashing:before, .dot-flashing:after, .dot-flashing {
  content: '';
  display: inline-block;
  border-radius: 50%;
  width: 0.5em;
  height: 0.5em;
  background: #2e4cff;
  animation: dotFlashing 1s infinite linear alternate;
  margin: 0 0.1em;
}
.dot-flashing:after {
  animation-delay: 0.5s;
}
@keyframes dotFlashing {
  0% { opacity: 0.2; }
  100% { opacity: 1; }
}
.fade-slide-enter-active {
  transition: all 0.7s cubic-bezier(.55,0,.1,1);
}
.fade-slide-leave-active {
  transition: all 0.4s cubic-bezier(.55,0,.1,1);
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(40px);
}
.fade-slide-enter-to {
  opacity: 1;
  transform: translateY(0);
}
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