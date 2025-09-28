import { API_URL } from '@/app/constants'
import axios from 'axios'
import { cookies } from 'next/headers'

const apiClient = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
})

apiClient.interceptors.request.use(async (config) => {
    const token = (await cookies()).get('token')

    if (token) {
        config.headers.Authorization = `Bearer ${token.value}`
    }

    return config
})

export default apiClient
