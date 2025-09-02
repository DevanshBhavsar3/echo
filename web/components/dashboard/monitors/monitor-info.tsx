import { Monitor } from '@/app/dashboard/monitors/data-table'
import { UptimeChart } from './uptime-chart'

interface MonitorProps {
    monitor: Monitor
}

export function MonitorInfo({ monitor }: MonitorProps) {
    return (
        <div className="flex flex-col p-4">
            <UptimeChart monitor={monitor} />
        </div>
    )
}
