'use client'

import { getTicks } from '@/app/actions/ticks'
import { Tick } from '@/app/dashboard/monitors/[monitorId]/page'
import { Monitor } from '@/app/dashboard/monitors/data-table'
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
import { frequencyToMs } from '@/lib/utils'
import { LoaderCircle } from 'lucide-react'
import { useEffect, useState } from 'react'
import ReactCountryFlag from 'react-country-flag'
import { CartesianGrid, Line, LineChart, XAxis } from 'recharts'

const chartConfig = {
    responseTime: {
        label: 'Response Time',
        color: 'var(--chart-1)',
    },
} satisfies ChartConfig

interface UptimeChartProps {
    monitor: Monitor
}

export function UptimeChart({ monitor }: UptimeChartProps) {
    const [data, setData] = useState<Tick[]>([])
    const [timeRange, setTimeRange] = useState('1')
    const [region, setRegion] = useState(monitor.regions[0].regionName)

    useEffect(() => {
        async function fetchTick() {
            const ticks = await getTicks(
                monitor.id,
                parseInt(timeRange),
                region,
            )

            setData(ticks)
        }

        fetchTick()

        const intervalMs = frequencyToMs(monitor.frequency)
        const interval = setInterval(fetchTick, intervalMs)

        return () => clearInterval(interval)
    }, [timeRange, region, monitor])

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
                        <ToggleGroupItem value="1">Day</ToggleGroupItem>
                        <ToggleGroupItem value="7">Week</ToggleGroupItem>
                        <ToggleGroupItem value="30">Month</ToggleGroupItem>
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
                            <SelectItem value="1">Day</SelectItem>
                            <SelectItem value="7">Week</SelectItem>
                            <SelectItem value="30">Month</SelectItem>
                        </SelectContent>
                    </Select>
                </CardAction>
            </CardHeader>
            <CardContent className="px-2 pt-4 sm:px-6 sm:pt-6">
                {data.length <= 0 ? (
                    <div className="flex aspect-auto h-[250px] w-full items-center justify-center">
                        <LoaderCircle
                            size={18}
                            className="animate-spin opacity-50"
                        />
                    </div>
                ) : (
                    <ChartContainer
                        config={chartConfig}
                        className="aspect-auto h-[250px] w-full"
                    >
                        <LineChart
                            data={data}
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
                                            const dateObj = new Date(value)
                                            const time =
                                                dateObj.toLocaleTimeString(
                                                    'en-IN',
                                                    {
                                                        hour: '2-digit',
                                                        minute: '2-digit',
                                                    },
                                                )
                                            const date =
                                                dateObj.toLocaleDateString(
                                                    'en-IN',
                                                    {
                                                        year: 'numeric',
                                                        month: 'short',
                                                        day: 'numeric',
                                                    },
                                                )

                                            return `${time}, ${date}`
                                        }}
                                        formatter={(value, name) => (
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
                )}
            </CardContent>
        </Card>
    )
}
