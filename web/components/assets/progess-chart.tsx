'use client'

import { Area, AreaChart } from 'recharts'
import { ChartConfig, ChartContainer } from '../ui/chart'

export function ProgressChart() {
    const chartData = [
        { month: 'January', uptime: 20 },
        { month: 'February', uptime: 32 },
        { month: 'March', uptime: 36 },
        { month: 'April', uptime: 42 },
        { month: 'May', uptime: 44 },
        { month: 'June', uptime: 49 },
        { month: 'July', uptime: 50 },
        { month: 'August', uptime: 70 },
        { month: 'September', uptime: 75 },
        { month: 'October', uptime: 71 },
        { month: 'November', uptime: 67 },
        { month: 'December', uptime: 87 },
    ]
    const chartConfig = {
        uptime: {
            label: 'Uptime',
            color: 'var(--chart-1)',
        },
    } satisfies ChartConfig

    return (
        <ChartContainer
            config={chartConfig}
            className="fixed bottom-0 left-0 z-0 hidden h-full w-full opacity-50 md:block"
        >
            <AreaChart accessibilityLayer data={chartData}>
                <Area
                    dataKey="uptime"
                    fill="var(--color-primary)"
                    type="natural"
                    fillOpacity={0.3}
                    stroke="var(--color-chart-2)"
                />
            </AreaChart>
        </ChartContainer>
    )
}
