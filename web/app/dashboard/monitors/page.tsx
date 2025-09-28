import { DataTable } from './data-table'
import { DialogBox } from '@/components/dashboard/dialog'
import { createWebsite, getAllWebsites } from '../../actions/website'
import { DialogTrigger } from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Search } from 'lucide-react'

export default async function DashboardPage() {
    const data = await getAllWebsites()

    return (
        <>
            <header className="flex w-full shrink-0 items-center gap-2">
                <h1 className="text-foreground text-4xl font-medium">
                    Monitor
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
                            <Button
                                size="lg"
                                className="hidden text-sm font-medium sm:flex"
                            >
                                Add Monitor
                            </Button>
                        </DialogTrigger>
                    </DialogBox>
                </div>
            </header>

            <div className="flex flex-1 flex-col">
                <DataTable data={data} />
            </div>
        </>
    )
}
