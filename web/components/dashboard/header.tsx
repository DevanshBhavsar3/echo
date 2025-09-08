'use client'

interface DashboardHeaderProps {
    breadcrumb?: string[]
    children?: React.ReactNode
}

export function DashboardHeader(props: DashboardHeaderProps) {
    return (
        <header className="flex shrink-0 items-center gap-2">
            {/* <Breadcrumb>
                    <BreadcrumbList>
                        {props.breadcrumb?.map((item, index) => (
                            <Fragment key={index}>
                                <BreadcrumbItem>
                                    <BreadcrumbLink asChild>
                                        <Link
                                            href={`/dashboard/${item.toLowerCase()}`}
                                            className="text-foreground hover:text-foreground/80 text-base font-medium hover:underline"
                                        >
                                            {item}
                                        </Link>
                                    </BreadcrumbLink>
                                </BreadcrumbItem>
                                <BreadcrumbSeparator />
                            </Fragment>
                        ))}
                        <BreadcrumbItem>
                            <BreadcrumbPage className="text-foreground font-sans text-4xl font-medium">
                                {props.title}
                            </BreadcrumbPage>
                        </BreadcrumbItem>
                    </BreadcrumbList>
                </Breadcrumb> */}

            {props.children}
        </header>
    )
}
