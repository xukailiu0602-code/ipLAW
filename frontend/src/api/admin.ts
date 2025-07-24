import api from './index'

export function adminModelConfig(data: any) {
  return api.post('/admin/model', data)
}
export function adminSliceRule(data: any) {
  return api.post('/admin/slice', data)
}
export function adminReindex() {
  return api.post('/admin/reindex', {})
}
export function adminLogs() {
  return api.get('/admin/logs')
}
export function adminListUsers() {
  return api.get('/admin/users')
}
export function adminSetRole(data: any) {
  return api.post('/admin/role', data)
}