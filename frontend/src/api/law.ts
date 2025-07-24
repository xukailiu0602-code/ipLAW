import api from './index'

export function listLaws() {
  return api.get('/laws')
}