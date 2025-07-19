import { Button } from "../ui/button";
import { Separator } from "../ui/separator";
import { SidebarTrigger } from "../ui/sidebar";

export function DashboardHeader() {
  return (
    <header className="flex py-1 shrink-0 items-center gap-2 border-b">
      <div className="flex w-full items-center gap-1 px-4 lg:gap-2 lg:px-6">
        <SidebarTrigger />
        <Separator
          orientation="vertical"
          className="mx-2 data-[orientation=vertical]:h-4"
        />
        <h1 className="text-base font-medium font-sans">Monitors</h1>
        <div className="ml-auto flex items-center gap-2">
          <Button size="sm" className="hidden sm:flex">
            Add Monitor
          </Button>
        </div>
      </div>
    </header>
  );
}
