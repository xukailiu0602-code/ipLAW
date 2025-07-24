<template>
  <motion-div class="qa-banner" :initial="{opacity:0, y:-60}" :enter="{opacity:1, y:0, transition:{duration:1.2, type:'spring'}}">
    <div class="banner-title">知识产权智能问答</div>
    <div class="banner-desc">AI赋能 · 法律合规 · 智能风控</div>
  </motion-div>
</template>
<script setup lang="ts">
import { Motion } from '@vueuse/motion'
</script>
<style scoped>
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
</style>