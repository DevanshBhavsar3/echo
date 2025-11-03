'use server'

import apiClient from '@/lib/axios'
import { Tick } from '../dashboard/(user)/monitors/[monitorId]/page'

export async function getTicks(
    monitorId: string,
    timeRange: number,
    region: string,
) {
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
