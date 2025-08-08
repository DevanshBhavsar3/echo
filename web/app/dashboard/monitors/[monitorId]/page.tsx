import { auth } from '@/app/auth'
import { DashboardHeader } from '@/components/dashboard/header'
import { redirect } from 'next/navigation'
import { Monitor } from '../data-table'
import axios from 'axios'
import { API_URL } from '@/app/constants'
import { MonitorInfo } from '@/components/dashboard/monitor/monitor-info'

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
    try {
        const res = await axios.get(`${API_URL}/website/${monitorId}`, {
            headers: {
                Authorization: `Bearer ${user.token}`,
            },
        })

        monitor = res.data as Monitor
    } catch (error) {
        console.error('Error fetching data:', error)
        redirect('/error')
    }

    return (
        <>
            <DashboardHeader title="Monitor Name" />
            <div className="flex flex-1 flex-col p-2">
                <MonitorInfo data={monitor} />
            </div>
        </>
    )
}
