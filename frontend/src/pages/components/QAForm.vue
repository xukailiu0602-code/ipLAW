<template>
  <a-card class="qa-card" title="智能问答">
    <a-form @submit.prevent="emitAsk">
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
  </a-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const query = ref('')
const emit = defineEmits(['on-ask'])
const props = defineProps<{ loading: boolean }>()
const emitAsk = () => {
  if (!query.value) return
  emit('on-ask', query.value)
}
</script>
<style scoped>
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
</style>