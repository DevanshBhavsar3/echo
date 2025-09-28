'use server'

import { Metrics, websiteSchema } from '@/lib/types'
import axios, { AxiosError } from 'axios'
import { revalidatePath } from 'next/cache'
import { Monitor } from '../dashboard/monitors/data-table'
import { DateRange } from 'react-day-picker'
import dayjs from '@/lib/dayjs'
import apiClient from '@/lib/axios'

export async function getAllWebsites() {
    try {
        const res = await apiClient.get(`/website`)

        return res.data as Monitor[]
    } catch (error) {
        console.error('Error fetching websites:', error)
        return []
    }
}

export async function createWebsite(_: unknown, formData: FormData) {
    const parsedData = websiteSchema.safeParse({
        url: formData.get('url'),
        frequency: formData.get('frequency'),
        regions: formData.getAll('regions') as string[],
    })

    if (!parsedData.success) {
        return {
            errors: parsedData.error.flatten().fieldErrors,
        }
    }

    try {
        await apiClient.post(`/website`, parsedData.data)

        revalidatePath('/dashboard')
    } catch (error) {
        if (error instanceof AxiosError) {
            return {
                error:
                    error.response?.data?.error ||
                    'An error occurred while creating the website.',
            }
        }

        return {
            error: 'An unexpected error occurred while creating the website.',
        }
    }
}

export async function deleteWebsite(id: string) {
    try {
        await apiClient.delete(`/website/${id}`)

        revalidatePath('/dashboard/monitors')
    } catch (error) {
        console.error('Error deleting website:', error)
    }
}

export async function editWebsite(websiteId: string, formData: FormData) {
    const parsedData = websiteSchema.safeParse({
        url: formData.get('url'),
        frequency: formData.get('frequency'),
        regions: formData.getAll('regions') as string[],
    })

    if (!parsedData.success) {
        return {
            errors: parsedData.error.flatten().fieldErrors,
        }
    }

    try {
        await apiClient.put(`/website/${websiteId}`, parsedData.data)

        revalidatePath('/dashboard/monitors')
    } catch (error) {
        if (error instanceof AxiosError) {
            return {
                error:
                    error.response?.data?.error ||
                    'An error occurred while editing the website.',
            }
        }

        return {
            error: 'An unexpected error occurred while editing the website.',
        }
    }
}

export async function pingWebsite(_: unknown, url: string) {
    try {
        await axios.head(url, {
            withCredentials: false,
            timeout: 5000,
        })

        return {
            status: true,
        }
    } catch (e) {
        if (e instanceof AxiosError) {
            return {
                status: false,
                error: e.response?.statusText || 'Failed to ping the website.',
            }
        }

        return {
            status: false,
            error: 'An unexpected error occurred while pinging the website.',
        }
    }
}

export async function getMonitorDetails(monitorId: string) {
    try {
        const monitorRes = await apiClient.get(`/website/${monitorId}`)

        return monitorRes.data as Monitor
    } catch (error) {
        console.error('Error fetching monitor details:', error)
    }
}

export async function getMonitorMetrics(monitorId: string, region: string) {
    try {
        const metricRes = await apiClient.get(
            `/website/metrics/${monitorId}?region=${region}`,
        )

        return metricRes.data as Metrics
    } catch (error) {
        console.error('Error fetching monitor metrics:', error)
    }
}

export async function getUptime(monitorId: string, range: DateRange) {
    try {
        const res = await apiClient.get(
            `/website/uptime/${monitorId}?from=${dayjs(range.from).format('YYYY-MM-DD')}&to=${dayjs(range.to).format('YYYY-MM-DD')}`,
        )

        return {
            success: true,
            data: {
                ...res.data,
                custom: true,
            },
        }
    } catch (error) {
        console.error('Error fetching monitor uptime:', error)
        return { success: false }
    }
}
