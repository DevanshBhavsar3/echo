import { getMonitorDetails } from '@/app/actions/website'
import { MonitorPage } from '@/components/pages/dashboard/monitors/monitor/page'
import { ChevronLeft, TriangleAlert } from 'lucide-react'
import Link from 'next/link'

export type Tick = {
    time: string
    responseTime: string
}

export default async function Page({
    params,
}: {
    params: Promise<{ monitorId: string }>
}) {
    const { monitorId } = await params
    const monitor = await getMonitorDetails(monitorId)

    if (!monitor) {
        return (
            <div className="flex flex-1 flex-col">
                <Link
                    href={'/dashboard/monitors'}
                    className="flex items-center hover:underline"
                >
                    <ChevronLeft size={16} />
                    Monitors
                </Link>
                <div className="text-muted-foreground flex flex-1 flex-col items-center justify-center gap-3 p-4">
                    <TriangleAlert size={52} />
                    <p>The monitor you are looking for does not exist.</p>
                </div>
            </div>
        )
    }

    return <MonitorPage monitor={monitor} />
}
