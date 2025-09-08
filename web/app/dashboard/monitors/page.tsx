import { DataTable, Monitor } from './data-table'
import { auth } from '../../auth'
import { redirect } from 'next/navigation'
import { API_URL } from '../../constants'
import axios from 'axios'
import { DashboardHeader } from '@/components/dashboard/header'
import { DialogBox } from '@/components/dashboard/dialog'
import { createWebsite } from '../../actions/website'
import { DialogTrigger } from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Search } from 'lucide-react'

export default async function DashboardPage() {
    const user = await auth()

    if (!user?.user.id) {
        return redirect('/login')
    }

    let data: Monitor[] = []

    try {
        const res = await axios.get(`${API_URL}/website`, {
            headers: {
                Authorization: `Bearer ${user.token}`,
            },
        })

        data = (res.data as Monitor[]) || []
    } catch (error) {
        console.error('Error fetching data:', error)
        redirect('/error')
    }

    return (
        <>
            <DashboardHeader>
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
            </DashboardHeader>

            <div className="flex flex-1 flex-col">
                <DataTable data={data} />
            </div>
        </>
    )
}
