<template>
  <a-card title="后台管理">
    <a-tabs>
      <a-tab-pane key="model" tab="模型配置">
        <a-button @click="onModelConfig">保存模型配置</a-button>
      </a-tab-pane>
      <a-tab-pane key="slice" tab="切片规则">
        <a-button @click="onSliceRule">保存切片规则</a-button>
      </a-tab-pane>
      <a-tab-pane key="reindex" tab="索引重建">
        <a-button @click="onReindex">重建索引</a-button>
      </a-tab-pane>
      <a-tab-pane key="logs" tab="日志监控">
        <a-button @click="onLogs">刷新日志</a-button>
        <ul>
          <li v-for="l in logs" :key="l">{{l}}</li>
        </ul>
      </a-tab-pane>
      <a-tab-pane key="users" tab="用户管理">
        <a-button @click="onListUsers">刷新用户</a-button>
        <ul>
          <li v-for="u in users" :key="u">{{u}}</li>
        </ul>
      </a-tab-pane>
    </a-tabs>
  </a-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { adminModelConfig, adminSliceRule, adminReindex, adminLogs, adminListUsers } from '../api/admin'
const logs = ref<string[]>([])
const users = ref<string[]>([])
const onModelConfig = async () => { await adminModelConfig({}) }
const onSliceRule = async () => { await adminSliceRule({}) }
const onReindex = async () => { await adminReindex() }
const onLogs = async () => { logs.value = (await adminLogs()).logs }
const onListUsers = async () => { users.value = (await adminListUsers()).users }
</script>