<template>
  <a-layout>
    <a-layout-header>
      <a-menu mode="horizontal" theme="dark" :selected-keys="[active]">
        <a-menu-item key="qa" @click="go('/qa')">问答</a-menu-item>
        <a-menu-item key="upload" @click="go('/upload')">文档上传</a-menu-item>
        <a-menu-item key="report" @click="go('/report/1')">报告导出</a-menu-item>
        <a-menu-item key="admin" @click="go('/admin')">后台管理</a-menu-item>
        <a-menu-item key="logout" @click="logout">退出</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding:24px">
      <router-view />
    </a-layout-content>
  </a-layout>
</template>
<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()
const go = (path: string) => router.push(path)
const logout = () => { localStorage.removeItem('token'); router.push('/login') }
const active = route.path.split('/')[1] || 'qa'
</script>
<style>
body { background: #f0f2f5; }
</style>