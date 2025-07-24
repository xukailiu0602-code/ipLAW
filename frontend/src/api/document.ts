import api from './index'

export function uploadDocument(file: File) {
  const form = new FormData()
  form.append('file', file)
  return api.post('/upload', form)
}