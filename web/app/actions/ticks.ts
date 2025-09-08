'use server'

import { API_URL } from '../constants'
import { Tick } from '../dashboard/monitors/[monitorId]/page'
import { auth } from '../auth'
import axios from 'axios'

export async function getTicks(
    monitorId: string,
    timeRange: number,
    region: string,
) {
    const user = await auth()
    if (!user?.token) {
        return []
    }

    try {
        const ticksRes = await axios.get(
            `${API_URL}/website/ticks/${monitorId}?days=${timeRange}&region=${region}`,
            {
                headers: {
                    Authorization: `Bearer ${user.token}`,
                },
            },
        )

        return (ticksRes.data as Tick[]) || []
    } catch (error) {
        console.error('Error fetching data:', error)
        return []
    }
}
