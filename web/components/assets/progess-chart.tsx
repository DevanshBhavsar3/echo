"use client"

import { Area, AreaChart, XAxis } from "recharts";
import { ChartConfig, ChartContainer, ChartTooltipContent } from "../ui/chart";

export function ProgressChart() {
  const chartData = [
    { month: "January", uptime: 20 },
    { month: "February", uptime: 32 },
    { month: "March", uptime: 50 },
    { month: "April", uptime: 65 },
    { month: "May", uptime: 70 },
    { month: "June", uptime: 67 },
    { month: "July", uptime: 86 },
    { month: "August", uptime: 90 },
    { month: "September", uptime: 93 },
    { month: "October", uptime: 87 },
    { month: "November", uptime: 89 },
    { month: "December", uptime: 99 },
  ]
  const chartConfig = {
    uptime: {
      label: "Uptime",
      color: "var(--chart-1)",
    },
  } satisfies ChartConfig

  return (
    <ChartContainer config={chartConfig} className="hidden md:block fixed z-0 bottom-0 left-0 w-full h-1/2 opacity-50">
      <AreaChart
        accessibilityLayer
        data={chartData}
      >
        <Area
          dataKey="uptime"
          fill="var(--color-primary)"
          type="natural"
          fillOpacity={0.3}
          stroke="var(--color-chart-2)"
        />
      </AreaChart>
    </ChartContainer>
  );
}
