import { DashboardHeader } from "@/components/dashboard/header";
import { DataTable, Monitors } from "./data-table";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { DashboardSidebar } from "@/components/dashboard/sidebar";

export default function DashboardPage() {
  const data: Monitors[] = [
    {
      id: "1",
      website: "https://example.com",
      status: "up",
      uptime: ["up", "up", "down", "up", "unknown"],
      lastChecked: "2023-10-01 12:00",
    },
    {
      id: "2",
      website: "https://another-example.com",
      status: "down",
      uptime: ["down", "down", "up", "unknown", "processing"],
      lastChecked: "2023-10-01 12:05",
    },
    {
      id: "3",
      website: "https://yet-another-example.com",
      status: "unknown",
      uptime: ["unknown", "up", "down", "up", "processing"],
      lastChecked: "2023-10-01 12:10",
    },
  ];

  return (
    <SidebarProvider>
      <DashboardSidebar variant="sidebar" />
      <SidebarInset>
        <DashboardHeader />
        <div className="flex flex-1 flex-col p-2">
          <DataTable data={data} />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
