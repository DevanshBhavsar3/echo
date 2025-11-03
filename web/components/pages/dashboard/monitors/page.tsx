'use client'

import { createWebsite } from '@/app/actions/website'
import { DialogBox } from '@/components/dashboard/dialog'
import { useAuth } from '@/components/providers/auth-provider'
import { Button } from '@/components/ui/button'
import { DialogTrigger } from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Search } from 'lucide-react'
import { redirect } from 'next/navigation'
import { DataTable, Monitor } from './data-table'

export function MonitorsPage({ monitors }: { monitors: Monitor[] }) {
    const { user } = useAuth()

    if (user && user.isAdmin) {
        redirect('/dashboard/admin')
    }

    return (
        <>
            <header className="flex w-full shrink-0 items-center gap-2">
                <h1 className="text-foreground text-4xl font-medium">
                    Monitors
                </h1>
                <div className="ml-auto flex items-center gap-2">
                    <Input
                        type="text"
                        placeholder="Search"
                        icon={<Search size={18} />}
                    />

                    <DialogBox
                        label={'Add Monitor'}
                        description={
                            'Add a new monitor to track the uptime of your website.'
                        }
                        onSubmitAction={createWebsite}
                    >
                        <DialogTrigger asChild>
                            <Button className="hidden text-sm font-medium sm:flex">
                                Add Monitor
                            </Button>
                        </DialogTrigger>
                    </DialogBox>
                </div>
            </header>

            <div className="flex flex-1 flex-col">
                <DataTable data={monitors} />
            </div>
        </>
    )
}
