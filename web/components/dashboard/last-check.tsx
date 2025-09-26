import { Tick } from '@/app/dashboard/monitors/data-table'
import dayjs from '@/lib/dayjs'
import { useEffect, useState } from 'react'

export function LastChecked({
    ticks = [],
    createdAt,
    className,
    ...props
}: {
    ticks?: Tick[]
    createdAt: string
    className?: string
    props?: React.ComponentProps<'span'>
}) {
    const [elapsedTime, setElapsedTime] = useState('')

    const latestTime =
        ticks.length > 0 ? ticks[ticks.length - 1].time : createdAt

    useEffect(() => {
        setElapsedTime(dayjs(latestTime).fromNow())
    }, [latestTime])

    return (
        <span className={className} {...props}>
            {elapsedTime || 'Loading...'}
        </span>
    )
}
