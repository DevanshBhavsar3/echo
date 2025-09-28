'use server'

import { cookies } from 'next/headers'
import { Tick } from '../dashboard/monitors/[monitorId]/page'
import apiClient from '@/lib/axios'

export async function getTicks(
    monitorId: string,
    timeRange: number,
    region: string,
) {
    const cookieStore = await cookies()
    const token = cookieStore.get('token')?.value

    if (!token) {
        return { error: 'No token found' }
    }

    try {
        const ticksRes = await apiClient.get(
            `/website/ticks/${monitorId}?days=${timeRange}&region=${region}`,
        )

        return (ticksRes.data as Tick[]) || []
    } catch (error) {
        console.error('Error fetching data:', error)
        return []
    }
}
