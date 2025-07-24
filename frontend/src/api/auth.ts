import api from './index'

export function login(data: { username: string; password: string }) {
  return api.post('/login', data)
}

export function register(data: { username: string; password: string; email: string }) {
  return api.post('/register', data)
}

export function getMe() {
  return api.get('/user/me')
}