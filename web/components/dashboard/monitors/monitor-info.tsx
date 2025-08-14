import { Tick } from '@/app/dashboard/monitors/[monitorId]/page'
import { Monitor } from '@/app/dashboard/monitors/data-table'

interface MonitorProps {
    monitor: Monitor
    ticks: Tick
}

export function MonitorInfo({ monitor, ticks }: MonitorProps) {
    return (
        <div className="flex flex-col p-4">
            {JSON.stringify(monitor)}, {JSON.stringify(ticks)}
        </div>
    )
}
