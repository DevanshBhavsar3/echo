import { fetchRegions } from '@/app/actions/region'
import { AdminPage } from '@/components/pages/dashboard/admin/page'

export default async function Page() {
    const regions = await fetchRegions()

    return (
        <div className="mx-auto grid max-w-5xl gap-5 py-10">
            <AdminPage regions={regions} />
        </div>
    )
}
