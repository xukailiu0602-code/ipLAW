<template>
  <div class="qa-bg">
    <Banner />
    <div class="qa-card-wrap">
      <QAForm @on-ask="handleAsk" :loading="loading" />
      <transition name="fade-slide">
        <AIAnswer v-if="showAnswer" :answer="answer" :citations="citations" :laws="laws" :cases="cases" :riskScore="riskScore" :rationale="rationale" />
      </transition>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { ask } from '../api/qa'
import Banner from './components/Banner.vue'
import QAForm from './components/QAForm.vue'
import AIAnswer from './components/AIAnswer.vue'
import anime from 'animejs'

interface Law { article: string; description: string }
interface Case { caseId: string; court: string; year: string; similarity: number }

const answer = ref('')
const citations = ref<string[]>([])
const laws = ref<Law[]>([])
const cases = ref<Case[]>([])
const riskScore = ref(0)
const rationale = ref('')
const loading = ref(false)
const showAnswer = ref(false)

const handleAsk = async (query: string) => {
  loading.value = true
  showAnswer.value = false
  const res = await ask({ query })
  answer.value = res.answer
  citations.value = res.citations
  laws.value = res.laws
  cases.value = res.cases
  riskScore.value = res.risk_score
  rationale.value = res.rationale
  loading.value = false
  showAnswer.value = true
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