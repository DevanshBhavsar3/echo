import { MetricsSection } from '@/components/dashboard/monitors/metrics'
import { UptimeChart } from '@/components/dashboard/monitors/uptime-chart'
import { UptimeTable } from '@/components/dashboard/monitors/uptime-table'
import { Button } from '@/components/ui/button'
import { Globe, LinkIcon, Settings } from 'lucide-react'
import { Monitor } from '../data-table'

export function MonitorPage({ monitor }: { monitor: Monitor }) {
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
