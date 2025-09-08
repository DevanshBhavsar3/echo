import { DashboardSidebar } from '@/components/dashboard/sidebar'
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar'

export default function DashboardLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <SidebarProvider>
            <DashboardSidebar variant="inset" />
            <SidebarInset className="space-y-12 px-18 pt-10">
                {children}
            </SidebarInset>
        </SidebarProvider>
    )
}
