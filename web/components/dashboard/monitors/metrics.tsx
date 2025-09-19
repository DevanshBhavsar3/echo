'use client'

import { getMonitorMetrics } from '@/app/actions/website'
import { MetricCard } from './metric-card'
import { useEffect, useState } from 'react'
import { Metrics } from '@/lib/types'
import { Monitor, Region } from '@/app/dashboard/monitors/data-table'
import { RegionSelect } from '@/components/region-select'

interface MetricsProps {
    monitor: Monitor
}

export function MetricsSection({ monitor }: MetricsProps) {
    const [data, setData] = useState<Metrics>()
    const [region, setRegion] = useState<Region>(monitor.regions[0])

    useEffect(() => {
        async function fetchMetrics() {
            const metric = await getMonitorMetrics(
                monitor.id,
                region.regionName,
            )

            setData(metric)
        }

        fetchMetrics()
    }, [region, monitor])

    if (!data) {
        return <div className="grid grid-cols-3 gap-5">Skeleton</div>
    }

    return (
        <div className="grid gap-3">
            <RegionSelect
                regions={monitor.regions}
                region={region}
                setRegion={setRegion}
            />
            <div className="grid grid-cols-3 gap-5">
                <MetricCard
                    title="Reponse Time"
                    description=" Measures the total time from request to full response."
                    data={data.response}
                    suffix="MS"
                    showTrend
                />
                <MetricCard
                    title="Status"
                    description="The operational state of your server."
                    data={data.status}
                />
                <MetricCard
                    title="Availability"
                    description=" The overall readiness of the service, including uptime and its ability to return successful responses."
                    data={data.availability}
                    suffix="%"
                    showTrend
                />
            </div>
        </div>
    )
}
