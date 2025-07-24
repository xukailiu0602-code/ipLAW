import api from './index'

export function exportReport(id: string) {
  return api.get(`/report/${id}`, { responseType: 'blob' })
}