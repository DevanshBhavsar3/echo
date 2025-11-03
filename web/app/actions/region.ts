'use server'

import { Region } from '@/components/pages/dashboard/monitors/data-table'
import apiClient from '@/lib/axios'
import { regionCodeSchema } from '@/lib/types'
import { AxiosError } from 'axios'
import { revalidatePath } from 'next/cache'

export async function fetchRegions() {
    try {
        const res = await apiClient.get(`/region`)

        return res.data.regions as Region[]
    } catch (error) {
        console.error('Error fetching regions:', error)
        return []
    }
}

export async function createRegion(_: unknown, data: FormData) {
    const parsedData = regionCodeSchema.safeParse({
        code: data.get('code'),
    })

    if (!parsedData.success) {
        return { error: parsedData.error.issues[0].message }
    }

    const { code } = parsedData.data

    try {
        await apiClient.post(`/region`, { code })
        revalidatePath('/admin')
    } catch (error) {
        if (error instanceof AxiosError) {
            console.error('Error creating regions:', error.message)
            return {
                error: error.response?.data.error,
            }
        }

        return {
            error: 'An unexpected error occurred',
        }
    }
}
