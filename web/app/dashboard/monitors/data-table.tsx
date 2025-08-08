'use client'

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
import { Button } from '@/components/ui/button'
import {
    DropdownMenuTrigger,
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { MoreHorizontal } from 'lucide-react'
import ReactCountryFlag from 'react-country-flag'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { deleteWebsite, editWebsite } from '../../actions/website'
import { DialogBox } from '@/components/dashboard/dialog'
import { DialogTrigger } from '@/components/ui/dialog'

export type Tick = {
    time: string
    status: string
}

export type Monitor = {
    id: string
    url: string
    frequency: string
    regions: string[]
    createdAt: string
    ticks: Tick[]
}

const statusStyles: Record<string, string> = {
    up: 'bg-green-400',
    down: 'bg-red-400',
    unknown: 'bg-yellow-400',
    default: 'bg-gray-400',
}

export const columns: ColumnDef<Monitor>[] = [
    {
        accessorKey: 'url',
        header: 'Monitor',
        cell: ({ row }) => {
            const url = row.getValue('url') as string
            const ticks = row.getValue('ticks') as Tick[]
            let status

            if (!ticks || ticks.length === 0) {
                status = 'unknown'
            } else {
                status = ticks[ticks.length - 1].status
            }

            const createdAt = row.original.createdAt

            return (
                <div className="flex items-center gap-3">
                    {
                        <span
                            className={`size-3 rounded-full ${statusStyles[status] || statusStyles.default} flex items-center justify-center opacity-70`}
                        ></span>
                    }
                    <div className="grid gap-1">
                        <a
                            href={url}
                            target="_blank"
                            className="text-md w-fit hover:underline"
                        >
                            {new URL(url).hostname}
                        </a>
                        <span className="text-muted-foreground text-xs">
                            Last checked{' '}
                            <LastCheckedCell
                                ticks={ticks}
                                createdAt={createdAt}
                            />
                        </span>
                    </div>
                </div>
            )
        },
    },
    {
        accessorKey: 'ticks',
        header: 'Uptime',
        cell: ({ row }) => {
            const ticks = row.getValue('ticks') as Tick[]

            if (ticks == null || ticks.length === 0) {
                return (
                    <div className="flex h-full items-center gap-1">
                        <span className="h-full w-4 bg-gray-400 p-1" />
                    </div>
                )
            }

            return (
                <div className="flex h-full items-center gap-1">
                    {ticks.map((tick, index) => {
                        if (tick.status == 'up') {
                            return (
                                <span
                                    className="h-full w-4 bg-green-400 p-1"
                                    key={index}
                                ></span>
                            )
                        } else if (tick.status == 'down') {
                            return (
                                <span
                                    className="h-full w-4 bg-red-400 p-1"
                                    key={index}
                                ></span>
                            )
                        } else if (tick.status == 'unknown') {
                            return (
                                <span
                                    className="h-full w-4 bg-yellow-400 p-1"
                                    key={index}
                                ></span>
                            )
                        }
                        return (
                            <span
                                className="h-full w-4 bg-gray-400 p-1"
                                key={index}
                            ></span>
                        )
                    })}
                </div>
            )
        },
    },
    {
        accessorKey: 'frequency',
        header: 'Frequency',
    },
    {
        accessorKey: 'regions',
        header: 'Regions',
        cell: ({ row }) => {
            const regions = row.getValue('regions') as string[]

            return (
                <div className="flex items-center gap-1">
                    {regions.map((region, index) => (
                        <ReactCountryFlag
                            key={index}
                            countryCode={region.toUpperCase()}
                            svg
                        />
                    ))}
                </div>
            )
        },
    },
    {
        id: 'actions',
        cell: ({ row }) => {
            const monitor = row.original

            return (
                <DialogBox
                    label={'Edit Monitor'}
                    description={'Edit your monitor configuration.'}
                    data={monitor}
                    onSubmitAction={(_: unknown, formData: FormData) =>
                        editWebsite(monitor.id, formData)
                    }
                >
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button variant="ghost" className="h-8 w-8 p-0">
                                <span className="sr-only">Open menu</span>
                                <MoreHorizontal className="h-4 w-4" />
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                            <DropdownMenuItem>
                                <DialogTrigger className="w-full text-left">
                                    Edit
                                </DialogTrigger>
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem
                                variant="destructive"
                                onClick={() => deleteWebsite(monitor.id)}
                            >
                                Delete
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </DialogBox>
            )
        },
    },
]

export function DataTable({ data }: { data: Monitor[] }) {
    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
    })
    const router = useRouter()

    useEffect(() => {
        const timer = setInterval(() => {
            router.refresh()
        }, 30000) // Refresh every 30 seconds

        return () => clearInterval(timer)
    })

    return (
        <div className="border">
            <Table>
                <TableHeader className="bg-muted sticky top-0 z-10 font-sans">
                    {table.getHeaderGroups().map((headerGroup) => (
                        <TableRow key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <TableHead key={header.id}>
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

                {/** onClick={() =>
                                    router.push(
                                        `/dashboard/monitor/${row.original.id}`,
                                    )
                                }
**/}
                <TableBody>
                    {table.getRowModel().rows?.length ? (
                        table.getRowModel().rows.map((row) => (
                            <TableRow key={row.id} className="cursor-pointer">
                                {row.getVisibleCells().map((cell) => (
                                    <TableCell key={cell.id}>
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
                                No Monitors.
                            </TableCell>
                        </TableRow>
                    )}
                </TableBody>
            </Table>
        </div>
    )
}

function LastCheckedCell({
    ticks,
    createdAt,
    className,
    ...props
}: {
    ticks?: Tick[]
    createdAt: string
    className?: string
    props?: React.ComponentProps<'span'>
}) {
    const [elapsedTime, setElapsedTime] = useState(0)

    useEffect(() => {
        const interval = setInterval(() => {
            let elapsedTime

            if (!ticks || ticks.length === 0) {
                elapsedTime = Math.floor(
                    (Date.now() - new Date(createdAt).getTime()) / 1000,
                )
            } else {
                elapsedTime = Math.floor(
                    (Date.now() - new Date(ticks[0].time).getTime()) / 1000,
                )
            }

            setElapsedTime(elapsedTime)
        }, 1000)

        return () => clearInterval(interval)
    }, [ticks, createdAt])

    return (
        <span className={className} {...props}>
            {Math.floor(elapsedTime / 60) > 0
                ? `${Math.floor(elapsedTime / 60)} minute(s) ago`
                : `${Math.floor(elapsedTime % 3600)} second(s) ago`}
        </span>
    )
}
