import { ProgressChart } from "@/components/assets/progess-chart";
import { LoginForm } from "@/components/signin-form";
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb";

export default function LoginPage() {
  return (
    <div className="min-h-svh w-full flex flex-col justify-center items-center text-left p-6 md:p-10">

      <Breadcrumb className="absolute top-0 left-0 w-full px-4 py-2 z-10">
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink href="/">Home</BreadcrumbLink>
          </BreadcrumbItem>
          <BreadcrumbSeparator />
          <BreadcrumbItem>
            <BreadcrumbPage>Login</BreadcrumbPage>
          </BreadcrumbItem>
        </BreadcrumbList>
      </Breadcrumb>

      <div className="w-full max-w-sm z-10">
        <LoginForm />
      </div>

      <ProgressChart />

      <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 text-sm text-muted-foreground">
        Echo is a free and open-source project.
      </div>
    </div >
  );
}


