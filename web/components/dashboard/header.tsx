import { createWebsite } from '@/app/actions/website'
import { Separator } from '../ui/separator'
import { SidebarTrigger } from '../ui/sidebar'
import { DialogBox } from './dialog'
import { DialogTrigger } from '../ui/dialog'
import { Button } from '../ui/button'

export function DashboardHeader() {
    return (
        <header className="flex shrink-0 items-center gap-2 border-b py-1">
            <div className="flex w-full items-center gap-1 px-4 lg:gap-2 lg:px-6">
                <SidebarTrigger />
                <Separator
                    orientation="vertical"
                    className="mx-2 data-[orientation=vertical]:h-4"
                />
                <h1 className="font-sans text-base font-medium">Monitors</h1>
                <div className="ml-auto flex items-center gap-2">
                    <DialogBox
                        label={'Add Monitor'}
                        description={
                            'Add a new monitor to track the uptime of your website.'
                        }
                        onSubmitAction={createWebsite}
                    >
                        <DialogTrigger asChild>
                            <Button size="sm" className="hidden sm:flex">
                                Add Monitor
                            </Button>
                        </DialogTrigger>
                    </DialogBox>
                </div>
            </div>
        </header>
    )
}
