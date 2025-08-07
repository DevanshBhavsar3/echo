import { Sidebar } from '../ui/sidebar'

export function DashboardSidebar({
    ...props
}: React.ComponentProps<typeof Sidebar>) {
    return (
        <Sidebar collapsible="offcanvas" {...props}>
            Hi
        </Sidebar>
    )
}
