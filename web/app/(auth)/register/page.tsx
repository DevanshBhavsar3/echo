import { ProgressChart } from '@/components/assets/progess-chart'
import { RegisterForm } from '@/components/register-form'
import {
    Breadcrumb,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbList,
    BreadcrumbPage,
    BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'

export default function RegisterPage() {
    return (
        <div className="flex min-h-svh w-full flex-col items-center justify-center p-6 text-left md:p-10">
            <Breadcrumb className="absolute left-0 top-0 z-10 w-full px-4 py-2">
                <BreadcrumbList>
                    <BreadcrumbItem>
                        <BreadcrumbLink href="/">Home</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>Register</BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>

            <div className="z-10 w-full max-w-sm">
                <RegisterForm />
            </div>

            <ProgressChart />

            <div className="text-muted-foreground absolute bottom-4 left-1/2 -translate-x-1/2 transform text-sm">
                Echo is a free and open-source project.
            </div>
        </div>
    )
}
