import { DashboardHeader } from "@/components/dashboard/header";
import { DataTable, Monitors } from "./data-table";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { DashboardSidebar } from "@/components/dashboard/sidebar";
import { auth } from "../auth";
import { redirect } from "next/navigation";
import { API_URL } from "../constants";
import axios from "axios";

export default async function DashboardPage() {
  const user = await auth()

  if (!user?.user.id) {
    return redirect("/login");
  }

  let data: Monitors[] = [];

  try {
    const res = await axios.get(`${API_URL}/website`, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    })

    data = res.data as Monitors[] || [];

    data = res.data.map((item: Monitors) => ({
      ...item,
      ticks: item.ticks.sort((a, b) => new Date(b.time).getTime() - new Date(a.time).getTime())
    }));

  } catch (error) {
    console.error("Error fetching data:", error);
    redirect("/error")
  }

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

export const revalidate = 30000;
