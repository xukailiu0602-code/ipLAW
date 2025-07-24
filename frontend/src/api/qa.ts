import api from './index'

export function ask(data: { query: string; doc_ids?: string[]; history?: any[] }) {
  return api.post('/ask', data)
}