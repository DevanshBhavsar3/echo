import { getAllWebsites } from '@/app/actions/website'
import { MonitorsPage } from '@/components/pages/dashboard/monitors/page'

export default async function Page() {
    const monitors = await getAllWebsites()

    return <MonitorsPage monitors={monitors} />
}
