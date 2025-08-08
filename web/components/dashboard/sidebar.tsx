import Link from 'next/link'
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
} from '../ui/sidebar'
import { Globe, Radio } from 'lucide-react'
import { auth } from '@/app/auth'
import { redirect } from 'next/navigation'
import { NavUser } from './nav-user'

const data = {
    main: [
        {
            title: 'Monitors',
            icon: Globe,
            href: '/dashboard/monitors',
        },
        {
            title: 'Status Pages',
            icon: Radio,
            href: '/dashboard/status-pages',
        },
    ],
}

export async function DashboardSidebar({
    ...props
}: React.ComponentProps<typeof Sidebar>) {
    const user = await auth()

    if (!user?.user.id) {
        return redirect('/login')
    }

    return (
        <Sidebar collapsible="offcanvas" {...props}>
            <SidebarHeader>
                <SidebarMenu>
                    <SidebarMenuItem>
                        <SidebarMenuButton
                            asChild
                            className="data-[slot=sidebar-menu-button]:!p-1.5"
                        >
                            <Link href="/">
                                <span className="font-sans text-base font-semibold">
                                    Echo
                                </span>
                            </Link>
                        </SidebarMenuButton>
                    </SidebarMenuItem>
                </SidebarMenu>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupContent className="flex flex-col gap-2">
                        <SidebarMenu>
                            {data.main.map((item) => (
                                <SidebarMenuItem key={item.title}>
                                    <SidebarMenuButton
                                        tooltip={item.title}
                                        asChild
                                    >
                                        <Link href={item.href}>
                                            {item.icon && <item.icon />}
                                            <span>{item.title}</span>
                                        </Link>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter>
                <NavUser user={user.user} />
            </SidebarFooter>
        </Sidebar>
    )
}
