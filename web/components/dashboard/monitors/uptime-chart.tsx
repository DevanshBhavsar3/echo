'use client'

import { Tick } from '@/app/dashboard/monitors/[monitorId]/page'
import { Monitor, statusStyles } from '@/app/dashboard/monitors/data-table'
import {
    Card,
    CardAction,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/ui/card'
import {
    ChartConfig,
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
} from '@/components/ui/chart'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { ToggleGroup, ToggleGroupItem } from '@/components/ui/toggle-group'
import { cn, getDateBeforeDays } from '@/lib/utils'
import { useState } from 'react'
import ReactCountryFlag from 'react-country-flag'
import { CartesianGrid, Line, LineChart, XAxis } from 'recharts'

const chartConfig = {
    responseTime: {
        label: 'Response Time',
        color: 'var(--chart-1)',
    },
} satisfies ChartConfig

interface UptimeChartProps {
    ticks: Tick[]
    monitor: Monitor
}

export function UptimeChart({ ticks, monitor }: UptimeChartProps) {
    const [timeRange, setTimeRange] = useState('1d')
    const [region, setRegion] = useState(monitor.regions[0].regionName)

    const filteredData = ticks.filter((tick) => {
        const date = new Date(tick.time)

        let daysToSubtract = 1
        if (timeRange === '7d') {
            daysToSubtract = 7
        } else if (timeRange === '30d') {
            daysToSubtract = 30
        }

        const startDate = new Date(getDateBeforeDays(daysToSubtract))

        return date >= startDate
    })

    return (
        <Card className="@container/card">
            <CardHeader>
                <CardTitle>Uptime</CardTitle>
                <CardDescription>
                    <Select value={region} onValueChange={setRegion}>
                        <SelectTrigger
                            className="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate"
                            size="sm"
                            aria-label="Select a value"
                        >
                            <SelectValue placeholder="Select Region" />
                        </SelectTrigger>
                        <SelectContent>
                            {monitor.regions.map((r) => (
                                <SelectItem
                                    key={r.regionId}
                                    value={r.regionName}
                                    className="flex items-center"
                                >
                                    <div className="flex items-center gap-2">
                                        <ReactCountryFlag
                                            countryCode={r.regionName.toUpperCase()}
                                            svg
                                        />
                                        {r.regionName}
                                    </div>
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                </CardDescription>
                <CardAction>
                    <ToggleGroup
                        type="single"
                        value={timeRange}
                        onValueChange={(val) => {
                            if (val) setTimeRange(val)
                        }}
                        variant="outline"
                        className="hidden *:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
                    >
                        <ToggleGroupItem value="1d">Day</ToggleGroupItem>
                        <ToggleGroupItem value="7d">Week</ToggleGroupItem>
                        <ToggleGroupItem value="30d">Month</ToggleGroupItem>
                    </ToggleGroup>
                    <Select value={timeRange} onValueChange={setTimeRange}>
                        <SelectTrigger
                            className="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate @[767px]/card:hidden"
                            size="sm"
                            aria-label="Select a value"
                        >
                            <SelectValue placeholder="Day" />
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="1d">Day</SelectItem>
                            <SelectItem value="7d">Week</SelectItem>
                            <SelectItem value="30d">Month</SelectItem>
                        </SelectContent>
                    </Select>
                </CardAction>
            </CardHeader>
            <CardContent className="px-2 pt-4 sm:px-6 sm:pt-6">
                <ChartContainer
                    config={chartConfig}
                    className="aspect-auto h-[250px] w-full"
                >
                    <LineChart
                        data={filteredData}
                        margin={{
                            left: 12,
                            right: 12,
                        }}
                    >
                        <CartesianGrid vertical={false} />
                        <XAxis
                            dataKey="time"
                            tickLine={false}
                            axisLine={false}
                            tickMargin={8}
                            minTickGap={32}
                            tickFormatter={(value) => {
                                const date = new Date(value)
                                return date.toLocaleDateString('en-US', {
                                    month: 'short',
                                    day: 'numeric',
                                })
                            }}
                        />
                        <ChartTooltip
                            cursor={false}
                            content={
                                <ChartTooltipContent
                                    labelFormatter={(value) => {
                                        return new Date(
                                            value,
                                        ).toLocaleDateString('en-IN', {
                                            year: 'numeric',
                                            month: 'short',
                                            day: 'numeric',
                                            hour: '2-digit',
                                            minute: '2-digit',
                                            second: '2-digit',
                                        })
                                    }}
                                    formatter={(value, name, item) => (
                                        <>
                                            <div
                                                className="h-2.5 w-2.5 shrink-0 rounded-[2px] bg-(--color-bg)"
                                                style={
                                                    {
                                                        '--color-bg': `var(--color-${name})`,
                                                    } as React.CSSProperties
                                                }
                                            />
                                            {chartConfig[
                                                name as keyof typeof chartConfig
                                            ]?.label || name}
                                            <div className="text-foreground ml-auto flex items-baseline gap-0.5 font-mono font-medium tabular-nums">
                                                {value}
                                                <span className="text-muted-foreground font-normal">
                                                    ms
                                                </span>
                                            </div>
                                            <div className="text-foreground mt-1.5 flex basis-full items-center border-t pt-1.5 text-xs font-medium">
                                                Status
                                                <div
                                                    className={cn(
                                                        'text-foreground ml-auto flex items-baseline gap-0.5 font-mono font-medium tabular-nums',
                                                        statusStyles({
                                                            status: item.payload
                                                                .status,
                                                            intent: 'text',
                                                        }),
                                                    )}
                                                >
                                                    {item.payload.status
                                                        .charAt(0)
                                                        .toUpperCase() +
                                                        item.payload.status.slice(
                                                            1,
                                                        )}
                                                </div>
                                            </div>
                                        </>
                                    )}
                                    indicator="dot"
                                />
                            }
                        />
                        <Line
                            dataKey="responseTime"
                            type="step"
                            stroke="var(--chart-1)"
                            dot={false}
                        />
                    </LineChart>
                </ChartContainer>
            </CardContent>
        </Card>
    )
}
