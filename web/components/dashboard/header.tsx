'use client'

import { RefreshCw } from 'lucide-react'
import { Button } from '../ui/button'
import { Separator } from '../ui/separator'
import { SidebarTrigger } from '../ui/sidebar'
import { useRouter } from 'next/navigation'

interface DashboardHeaderProps {
    title: string
    children?: React.ReactNode
}

export function DashboardHeader(props: DashboardHeaderProps) {
    const router = useRouter()

    return (
        <header className="flex shrink-0 items-center gap-2 border-b py-1">
            <div className="flex w-full items-center gap-1 px-4 lg:gap-2 lg:px-6">
                <SidebarTrigger />
                <Separator
                    orientation="vertical"
                    className="mx-2 data-[orientation=vertical]:h-4"
                />
                <h1 className="font-sans text-base font-medium">
                    {props.title}
                </h1>
                <div className="ml-auto flex items-center gap-2">
                    <Button
                        variant="ghost"
                        size="icon"
                        className="hidden sm:flex"
                        onClick={() => router.refresh()}
                    >
                        <RefreshCw />
                    </Button>

                    {props.children}
                </div>
            </div>
        </header>
    )
}
