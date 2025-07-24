import api from './index'

export function listCases() {
  return api.get('/cases')
}