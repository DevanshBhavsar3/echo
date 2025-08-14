import { auth } from '@/app/auth'
import { DashboardHeader } from '@/components/dashboard/header'
import { redirect } from 'next/navigation'
import { Monitor } from '../data-table'
import axios from 'axios'
import { API_URL } from '@/app/constants'
import { MonitorInfo } from '@/components/dashboard/monitors/monitor-info'
import { getTimeRange } from '@/lib/utils'

export type Tick = {
    time: string
    responseTime: string
    status: string
    regionName: string
}

export default async function MonitorPage({
    params,
}: {
    params: Promise<{ monitorId: string }>
}) {
    const { monitorId } = await params
    const user = await auth()

    if (!user?.user.id) {
        return redirect('/login')
    }

    let monitor: Monitor | null = null
    let ticks: Tick | null = null

    try {
        const monitorRes = await axios.get(`${API_URL}/website/${monitorId}`, {
            headers: {
                Authorization: `Bearer ${user.token}`,
            },
        })

        monitor = monitorRes.data as Monitor

        const ticksRes = await axios.get(
            `${API_URL}/website/ticks/${monitorId}?start=${getTimeRange(3)}&end=${getTimeRange(0)}`,
            {
                headers: {
                    Authorization: `Bearer ${user.token}`,
                },
            },
        )

        ticks = (ticksRes.data as Tick[]) || []
    } catch (error) {
        console.error('Error fetching data:', error)
        redirect('/error')
    }

    return (
        <div>
            <DashboardHeader title="Monitor Name" />
            <div className="flex flex-1 flex-col p-2">
                <MonitorInfo monitor={monitor} ticks={ticks} />
            </div>
        </div>
    )
}
