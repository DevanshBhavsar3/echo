'use client'

import dayjs from '@/lib/dayjs'
import { getUptime } from '@/app/actions/website'
import { Monitor } from '@/app/dashboard/monitors/data-table'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/ui/table'
import {
    ColumnDef,
    flexRender,
    getCoreRowModel,
    useReactTable,
} from '@tanstack/react-table'
import { CalendarIcon, LoaderCircle } from 'lucide-react'
import { useState } from 'react'
import { DateRange } from 'react-day-picker'

export const columns: ColumnDef<Uptime>[] = [
    {
        accessorKey: 'time',
        header: 'Time',
        cell: ({ row }) => {
            const [from, to] = row.original.time.split(', ')

            const fromTime = dayjs(from)
            const toTime = dayjs(to)

            if (row.original.custom) {
                return (
                    fromTime.format('D MMMM, YYYY') +
                    ' to ' +
                    toTime.format('D MMMM, YYYY')
                )
            }

            const days = toTime.diff(fromTime, 'day')

            if (days == 0) {
                return 'Today'
            }

            return 'Last ' + toTime.diff(fromTime, 'day') + ' days'
        },
    },
    { accessorKey: 'availability', header: 'Availability' },
    { accessorKey: 'avg_response_time', header: 'Avg. Response Times' },
]

export type Uptime = {
    custom?: boolean
    time: string
    availability: string
    avg_response_time: string
}

export function UptimeTable({ monitor }: { monitor: Monitor }) {
    const [data, setData] = useState(monitor.uptime)
    const [dateRange, setDateRange] = useState<DateRange>({
        from: new Date(),
        to: new Date(),
    })
    const [loading, setLoading] = useState(false)

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
    })

    return (
        <div className="grid gap-3">
            <div className="border">
                <Table>
                    <TableHeader className="bg-muted sticky top-0 z-10">
                        {table.getHeaderGroups().map((headerGroup) => (
                            <TableRow key={headerGroup.id}>
                                {headerGroup.headers.map((header) => {
                                    return (
                                        <TableHead
                                            key={header.id}
                                            className={`text-muted-foreground font-mono uppercase ${header.index === 0 ? 'w-1/2 text-left' : 'w-auto text-center'}`}
                                        >
                                            {header.isPlaceholder
                                                ? null
                                                : flexRender(
                                                      header.column.columnDef
                                                          .header,
                                                      header.getContext(),
                                                  )}
                                        </TableHead>
                                    )
                                })}
                            </TableRow>
                        ))}
                    </TableHeader>

                    <TableBody>
                        {table.getRowModel().rows?.length ? (
                            table.getRowModel().rows.map((row) => (
                                <TableRow key={row.id}>
                                    {row
                                        .getVisibleCells()
                                        .map((cell, cellIdx) => (
                                            <TableCell
                                                key={cell.id}
                                                className={
                                                    cellIdx === 0
                                                        ? 'text-left'
                                                        : 'text-center'
                                                }
                                            >
                                                {flexRender(
                                                    cell.column.columnDef.cell,
                                                    cell.getContext(),
                                                )}
                                            </TableCell>
                                        ))}
                                </TableRow>
                            ))
                        ) : (
                            <TableRow>
                                <TableCell
                                    colSpan={columns.length}
                                    className="h-24 text-center"
                                >
                                    No uptime data available.
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </div>
            <div className="flex items-center gap-3">
                <Popover>
                    <PopoverTrigger asChild>
                        <Button
                            variant={'outline'}
                            className={
                                'text-muted-foreground w-[250px] pl-3 text-left font-normal'
                            }
                        >
                            <span>
                                {dayjs(dateRange.from).format('DD-MM-YYYY')} -{' '}
                                {dayjs(dateRange.to).format('DD-MM-YYYY')}
                            </span>
                            <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                        </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-auto p-0" align="start">
                        <Calendar
                            selected={dateRange}
                            onSelect={setDateRange}
                            mode="range"
                            required
                            disabled={(date) =>
                                date > new Date() ||
                                date < new Date(monitor.createdAt)
                            }
                            captionLayout="dropdown"
                        />
                    </PopoverContent>
                </Popover>
                <Button
                    onClick={async () => {
                        setLoading(true)
                        const uptime = await getUptime(monitor.id, dateRange)
                        setLoading(false)

                        if (!uptime.success) {
                            return
                        }

                        setData((prev) => {
                            if (
                                prev.length != 0 &&
                                prev[prev.length - 1].custom
                            ) {
                                prev.pop()
                            }

                            return [...prev, uptime.data as Uptime]
                        })
                    }}
                    className="w-[100px] px-4"
                >
                    {loading ? (
                        <LoaderCircle size={18} className="animate-spin" />
                    ) : (
                        'Calculate'
                    )}
                </Button>
            </div>
        </div>
    )
}
