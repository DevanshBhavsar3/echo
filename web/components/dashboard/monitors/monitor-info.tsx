import { Tick } from '@/app/dashboard/monitors/[monitorId]/page'
import { Monitor } from '@/app/dashboard/monitors/data-table'
import { UptimeChart } from './uptime-chart'

interface MonitorProps {
    monitor: Monitor
    ticks: Tick[]
}

export function MonitorInfo({ monitor, ticks }: MonitorProps) {
    return (
        <div className="flex flex-col p-4">
            <UptimeChart monitor={monitor} ticks={ticks} />
        </div>
    )
}
