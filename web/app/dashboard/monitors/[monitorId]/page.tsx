import { auth } from '@/app/auth'
import { DashboardHeader } from '@/components/dashboard/header'
import { redirect } from 'next/navigation'
import { MonitorInfo } from '@/components/dashboard/monitors/monitor-info'
import { getMonitorDetails } from '@/app/actions/website'

export type Tick = {
    time: string
    responseTime: string
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

    const monitor = await getMonitorDetails(monitorId)

    if (!monitor) {
        return (
            <div className="flex flex-1 flex-col">
                <DashboardHeader
                    title="Invalid Monitor"
                    breadcrumb={['Monitors']}
                />
                <div className="flex flex-1 items-center justify-center p-4">
                    <p className="text-muted-foreground">
                        The monitor you are looking for does not exist.
                    </p>
                </div>
            </div>
        )
    }

    return (
        <div>
            <DashboardHeader
                title={new URL(monitor.url).hostname}
                breadcrumb={['Monitors']}
            />
            <div className="flex flex-1 flex-col p-2">
                <MonitorInfo monitor={monitor} />
            </div>
        </div>
    )
}
