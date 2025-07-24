<template>
  <a-card title="报告导出">
    <a-button type="primary" @click="onExport">导出 PDF</a-button>
  </a-card>
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router'
import { exportReport } from '../api/report'
const route = useRoute()
const onExport = async () => {
  const res = await exportReport(route.params.id as string)
  const url = window.URL.createObjectURL(new Blob([res]))
  const a = document.createElement('a')
  a.href = url
  a.download = 'report.pdf'
  a.click()
  window.URL.revokeObjectURL(url)
}
</script>