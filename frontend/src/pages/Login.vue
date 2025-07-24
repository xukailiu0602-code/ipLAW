<template>
  <a-row justify="center" align="middle" style="height:100vh">
    <a-col :span="6">
      <a-card title="知识产权风险评估平台登录">
        <a-form @submit.prevent="onLogin">
          <a-form-item label="用户名">
            <a-input v-model:value="form.username" />
          </a-form-item>
          <a-form-item label="密码">
            <a-input-password v-model:value="form.password" />
          </a-form-item>
          <a-form-item>
            <a-button type="primary" html-type="submit" block :loading="loading">登录</a-button>
          </a-form-item>
          <a-form-item>
            <a-button type="link" @click="showRegister=true">注册新账号</a-button>
          </a-form-item>
        </a-form>
        <a-modal v-model:open="showRegister" title="注册" @ok="onRegister">
          <a-form>
            <a-form-item label="用户名">
              <a-input v-model:value="reg.username" />
            </a-form-item>
            <a-form-item label="邮箱">
              <a-input v-model:value="reg.email" />
            </a-form-item>
            <a-form-item label="密码">
              <a-input-password v-model:value="reg.password" />
            </a-form-item>
          </a-form>
        </a-modal>
      </a-card>
    </a-col>
  </a-row>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { login, register } from '../api/auth'
import { useRouter } from 'vue-router'
const router = useRouter()
const form = ref({ username: '', password: '' })
const loading = ref(false)
const showRegister = ref(false)
const reg = ref({ username: '', password: '', email: '' })
const onLogin = async () => {
  loading.value = true
  try {
    const res = await login(form.value)
    localStorage.setItem('token', res.token)
    router.push('/qa')
  } catch (e) { }
  loading.value = false
}
const onRegister = async () => {
  await register(reg.value)
  showRegister.value = false
}
</script>