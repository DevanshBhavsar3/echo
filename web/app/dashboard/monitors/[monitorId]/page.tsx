import { auth } from '@/app/auth'
import { redirect } from 'next/navigation'
import { getMonitorDetails } from '@/app/actions/website'

import {
    ChevronLeft,
    Globe,
    Link as LinkIcon,
    Settings,
    TriangleAlert,
} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { UptimeChart } from '@/components/dashboard/monitors/uptime-chart'
import Link from 'next/link'
import { MetricsSection } from '@/components/dashboard/monitors/metrics'
import { UptimeTable } from '@/components/dashboard/monitors/uptime-table'

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

    return (
        <>
            <header className="flex w-full shrink-0 items-center gap-2">
                <div className="flex w-full flex-col items-start justify-center">
                    <h2 className="text-muted-foreground flex items-center gap-1 font-mono uppercase">
                        <Globe size={16} />
                        Monitor
                    </h2>
                    <div className="flex w-full items-center justify-between">
                        <h1 className="text-foreground flex items-center gap-3 text-3xl font-medium">
                            {new URL(monitor.url).hostname}
                            <a
                                href={monitor.url}
                                target="_blank"
                                className="text-muted-foreground hover:text-foreground"
                            >
                                <LinkIcon />
                            </a>
                        </h1>
                        <Button
                            size="lg"
                            variant={'outline'}
                            className="hidden font-medium sm:flex"
                        >
                            <Settings />
                            Settings
                        </Button>
                    </div>
                </div>
            </header>
            <div className="flex flex-col gap-18">
                <UptimeChart monitor={monitor} />
                <MetricsSection monitor={monitor} />
                <UptimeTable monitor={monitor} />
            </div>
        </>
    )
}
