"use client";

import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { Button } from "@/components/ui/button";
import { DropdownMenuTrigger, DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator } from "@/components/ui/dropdown-menu";
import { MoreHorizontal } from "lucide-react";
import ReactCountryFlag from "react-country-flag";
import { useEffect, useState } from "react";

export type Tick = {
  time: string
  status: string;
}

export type Monitors = {
  id: string;
  url: string;
  frequency: number;
  regions: string[];
  createdAt: string;
  ticks: Tick[];
}

export const columns: ColumnDef<Monitors>[] = [
  {
    accessorKey: "url",
    header: "Url",
  },
  {
    id: "status",
    header: "Status",
    cell: ({ row }) => {
      const ticks = row.getValue("ticks") as Tick[];
      const status = ticks[ticks.length - 1].status;

      return (
        <div>
          {status.charAt(0).toUpperCase() + status.slice(1)}
        </div>
      );
    },
  },
  {
    accessorKey: "ticks",
    header: "Uptime",
    cell: ({ row }) => {
      const ticks = row.getValue("ticks") as Tick[];

      return (
        <div className="flex items-center gap-1">
          {
            ticks.map((tick, index) => {
              if (tick.status == "up") {
                return <span className="bg-green-400 h-full w-4 p-1" key={index}></span>
              } else if (tick.status == "down") {
                return <span className="bg-red-400 h-full w-4 p-1" key={index}></span>
              } else if (tick.status == "unknown") {
                return <span className="bg-yellow-400 h-full w-4 p-1" key={index}></span>
              }
              return <span className="bg-gray-400 h-full w-4 p-1" key={index}></span>
            })
          }
        </div >
      );
    },
  },
  {
    accessorKey: "frequency",
    header: "Frequency",
  },
  {
    accessorKey: "regions",
    header: "Regions",
    cell: ({ row }) => {
      const regions = row.getValue("regions") as string[];


      return (
        <div className="flex items-center gap-1">
          {regions.map((region, index) => (
            <ReactCountryFlag
              key={index}
              countryCode={region.toUpperCase()}
              svg
            />
          ))}
        </div>
      );
    },
  },
  {
    accessorKey: "lastChecked",
    header: "Last Checked",
    cell: ({ row }) => {
      const [currentTime, setCurrentTime] = useState(Date.now());

      useEffect(() => {
        const interval = setInterval(() => {
          setCurrentTime(Date.now());
        }, 1000);

        return () => clearInterval(interval);
      }, []);

      const ticks = row.getValue("ticks") as Tick[];

      if (ticks.length === 0) {
        return <div>No data</div>;
      }

      const lastTick = ticks[0];

      const lastTickTime = Date.parse(lastTick.time);
      const elapsedMilliseconds = currentTime - lastTickTime;

      const elapsedSeconds = Math.floor(elapsedMilliseconds / 1000);
      const elapsedMinutes = Math.floor(elapsedSeconds / 60);

      let displayTime;
      if (elapsedMinutes > 0) {
        displayTime = `${elapsedMinutes} minute(s) ago`;
      } else {
        displayTime = `${elapsedSeconds} second(s) ago`;
      }

      return (
        <div>
          {displayTime}
        </div>
      );
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem>Edit</DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem variant="destructive">Delete</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu >
      )
    },
  }
]

export function DataTable({
  data,
}: {
  data: Monitors[];
}) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  })

  return (
    <div className="border">
      <Table>
        <TableHeader className="bg-muted sticky top-0 z-10">
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                  </TableHead>
                )
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No Monitors.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
