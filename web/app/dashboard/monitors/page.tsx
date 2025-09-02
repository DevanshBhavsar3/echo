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
        <div>
            <DashboardHeader title={'Monitors'}>
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
            </DashboardHeader>

            <div className="flex flex-1 flex-col p-2">
                <DataTable data={data} />
            </div>
        </div>
    )
}
