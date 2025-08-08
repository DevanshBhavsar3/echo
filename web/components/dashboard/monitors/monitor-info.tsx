import { Monitor } from '@/app/dashboard/monitors/data-table'

interface MonitorProps {
    data: Monitor
}

export function MonitorInfo({ data }: MonitorProps) {
    return <div className="flex flex-col p-4">{JSON.stringify(data)}</div>
}
