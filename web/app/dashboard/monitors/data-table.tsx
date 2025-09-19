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
import {
    Check,
    Disc2,
    Loader,
    MoreHorizontal,
    X,
    CircleCheck,
    CircleX,
} from 'lucide-react'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { deleteWebsite, editWebsite } from '../../actions/website'
import { DialogBox } from '@/components/dashboard/dialog'
import { DialogTrigger } from '@/components/ui/dialog'
import { cva } from 'class-variance-authority'
import { cn } from '@/lib/utils'
import {
    Tooltip,
    TooltipTrigger,
    TooltipContent,
} from '@/components/ui/tooltip'
import { Badge } from '@/components/ui/badge'

export type Status = 'up' | 'down' | 'processing'

export type Tick = {
    time: string
    status: Status
}

export type Region = {
    regionId: string
    regionName: string
}

export type Monitor = {
    id: string
    url: string
    frequency: string
    regions: Region[]
    createdAt: string
    ticks: Tick[]
}

export const statusStyles = cva('', {
    variants: {
        status: {
            up: 'text-green-500',
            down: 'text-red-400',
            processing: 'text-yellow-400',
        },
        intent: {
            text: '',
            bg: '',
        },
    },
    compoundVariants: [
        { status: 'up', intent: 'text', className: 'text-green-500' },
        { status: 'down', intent: 'text', className: 'text-red-500' },
        { status: 'processing', intent: 'text', className: 'text-yellow-500' },
        { status: 'up', intent: 'bg', className: 'bg-green-400' },
        { status: 'down', intent: 'bg', className: 'bg-red-400' },
        { status: 'processing', intent: 'bg', className: 'bg-gray-400' },
    ],
    defaultVariants: {
        status: 'processing',
        intent: 'text',
    },
})

export const columns: ColumnDef<Monitor>[] = [
    {
        accessorKey: 'url',
        header: 'Url',
        cell: ({ row }) => {
            const url = row.getValue('url') as string
            const ticks = row.getValue('ticks') as Tick[]
            const createdAt = row.original.createdAt
            let status: Status

            if (!ticks || ticks.length === 0) {
                status = 'processing'
            } else {
                status = ticks[ticks.length - 1].status
            }

            return (
                <div className="flex items-center gap-3">
                    <a
                        href={url}
                        target="_blank"
                        className="text-md w-fit hover:underline"
                    >
                        {new URL(url).hostname}
                    </a>
                    <Tooltip>
                        <TooltipTrigger>
                            <Badge
                                variant={'outline'}
                                className="text-muted-foreground items-center gap-2 rounded-full px-1.5 text-xs font-medium"
                            >
                                {status === 'up' ? (
                                    <CircleCheck
                                        size={12}
                                        className="text-green-500"
                                    />
                                ) : status == 'down' ? (
                                    <span className="text-red-500">
                                        <CircleX size={12} />
                                    </span>
                                ) : (
                                    <Loader
                                        size={12}
                                        className="animate-spin"
                                    />
                                )}
                                {status.charAt(0).toUpperCase() +
                                    status.slice(1)}
                            </Badge>
                        </TooltipTrigger>
                        <TooltipContent>
                            <LastCheckedCell
                                ticks={ticks}
                                createdAt={createdAt}
                                className="text-background gap flex items-center text-xs"
                            />
                        </TooltipContent>
                    </Tooltip>
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
                        <Tooltip>
                            <TooltipTrigger>
                                <span
                                    className={cn(
                                        'h-full w-4 p-1',
                                        statusStyles({
                                            status: 'processing',
                                            intent: 'bg',
                                        }),
                                    )}
                                ></span>
                            </TooltipTrigger>
                            <TooltipContent>
                                No uptime data available
                            </TooltipContent>
                        </Tooltip>
                    </div>
                )
            }

            return (
                <div className="flex h-full items-center gap-1">
                    {ticks.map((tick, index) => {
                        return (
                            <Tooltip key={index}>
                                <TooltipTrigger>
                                    <span
                                        className={cn(
                                            'h-full w-4 p-1',
                                            statusStyles({
                                                status: tick.status,
                                                intent: 'bg',
                                            }),
                                        )}
                                    ></span>
                                </TooltipTrigger>
                                <TooltipContent>
                                    <p>
                                        {new Date(tick.time).toLocaleDateString(
                                            'en-IN',
                                            {
                                                year: 'numeric',
                                                month: '2-digit',
                                                day: '2-digit',
                                                hour: '2-digit',
                                                minute: '2-digit',
                                                second: '2-digit',
                                            },
                                        )}
                                    </p>
                                </TooltipContent>
                            </Tooltip>
                        )
                    })}
                </div>
            )
        },
    },
    {
        accessorKey: 'frequency',
        header: 'Frequency',
        cell: ({ row }) => {
            const frequency = row.getValue('frequency') as string

            return (
                <Tooltip>
                    <TooltipTrigger>
                        <div className="text-muted-foreground flex items-center gap-2">
                            <Disc2 size={16} />
                            <span className="text-sm">{frequency}</span>
                        </div>
                    </TooltipTrigger>
                    <TooltipContent>
                        <p>Checked every {frequency}.</p>
                    </TooltipContent>
                </Tooltip>
            )
        },
    },
    {
        id: 'actions',
        cell: ({ row }) => {
            const monitor = row.original

            return (
                <div onClick={(e) => e.stopPropagation()}>
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
                </div>
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
    }, [])

    return (
        <div className="border">
            <Table>
                <TableHeader className="bg-muted sticky top-0 z-10">
                    {table.getHeaderGroups().map((headerGroup) => (
                        <TableRow key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <TableHead
                                        key={header.id}
                                        className="text-muted-foreground font-mono uppercase"
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
                            <TableRow
                                key={row.id}
                                className="cursor-pointer"
                                onClick={() =>
                                    router.push(
                                        `/dashboard/monitors/${row.original.id}`,
                                    )
                                }
                            >
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
                    (Date.now() -
                        new Date(ticks[ticks.length - 1].time).getTime()) /
                        1000,
                )
            }

            setElapsedTime(elapsedTime)
        }, 1000)

        return () => clearInterval(interval)
    }, [ticks, createdAt])

    return (
        <span className={className} {...props}>
            Last Checked{' '}
            {Math.floor(elapsedTime / 60) > 0
                ? `${Math.floor(elapsedTime / 60)} minute(s) ago`
                : `${Math.floor(elapsedTime % 3600)} second(s) ago`}
        </span>
    )
}
