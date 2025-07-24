<template>
  <a-card title="文档上传">
    <a-upload :before-upload="beforeUpload" :show-upload-list="false">
      <a-button type="primary">选择文件</a-button>
    </a-upload>
    <a-button type="primary" @click="onUpload" :loading="loading" style="margin-top:10px">上传</a-button>
    <div v-if="docId">上传成功，文档ID: {{docId}}</div>
  </a-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { uploadDocument } from '../api/document'
let file = null as File | null
const docId = ref('')
const loading = ref(false)
const beforeUpload = (f: File) => { file = f; return false }
const onUpload = async () => {
  if (!file) return
  loading.value = true
  const res = await uploadDocument(file)
  docId.value = res.doc_id
  loading.value = false
}
</script>