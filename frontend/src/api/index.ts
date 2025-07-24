import axios from 'axios'

const instance = axios.create({
  baseURL: '/api',
  timeout: 10000
})

instance.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) config.headers['Authorization'] = 'Bearer ' + token
  return config
})

instance.interceptors.response.use(
  res => res.data,
  err => {
    if (err.response && err.response.status === 401) {
      window.location.href = '/login'
    }
    return Promise.reject(err)
  }
)

export default instance