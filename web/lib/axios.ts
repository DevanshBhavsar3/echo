import axios from 'axios'
import { cookies } from 'next/headers'

const apiClient = axios.create({
    baseURL: process.env.NEXT_PUBLIC_DOCKER_API_URL,
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
