'use client'

import { Card, CardContent, CardHeader } from '@/components/ui/card'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import {
    Tooltip,
    TooltipContent,
    TooltipTrigger,
} from '@/components/ui/tooltip'
import { Latencies, Latency } from '@/lib/types'
import { Info } from 'lucide-react'
import { useState } from 'react'

interface MetricCardProps {
    title: string
    description: string
    data: Latencies
}

export function MetricCard({ title, description, data }: MetricCardProps) {
    const [latency, setLatency] = useState<Latency>(Latency.P99)

    return (
        <div className="grid gap-3">
            <div className="flex items-center gap-2">
                {title}
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Info size={16} className="text-muted-foreground" />
                    </TooltipTrigger>
                    <TooltipContent>{description}</TooltipContent>
                </Tooltip>
            </div>
            <Card>
                <CardHeader>
                    <Select
                        value={latency.toString()}
                        onValueChange={(value) =>
                            setLatency(value as unknown as Latency)
                        }
                    >
                        <SelectTrigger
                            className="flex w-24 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate"
                            size="sm"
                            aria-label="Select a value"
                        >
                            <SelectValue placeholder="Select Latency" />
                        </SelectTrigger>
                        <SelectContent>
                            {Object.values(Latency).map((l, idx) => (
                                <SelectItem key={idx} value={l}>
                                    {l}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                </CardHeader>
                <CardContent>
                    <div className="mb-12 flex h-24 flex-1 flex-col items-center justify-center gap-2 font-mono">
                        <p className="text-2xl font-medium">
                            {data[latency].current}
                        </p>
                        {data[latency].trend && (
                            <p className="text-sm">{data[latency].trend}</p>
                        )}
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
