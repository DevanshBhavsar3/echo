"use client";

import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { Button } from "@/components/ui/button";
import { DropdownMenuTrigger, DropdownMenu, DropdownMenuContent, DropdownMenuLabel, DropdownMenuItem, DropdownMenuSeparator } from "@/components/ui/dropdown-menu";
import { MoreHorizontal } from "lucide-react";

export type Monitors = {
  id: string;
  website: string;
  status: "up" | "down" | "unknown" | "processing";
  uptime: Array<"up" | "down" | "unknown" | "processing">;
  lastChecked: string;
}

export const columns: ColumnDef<Monitors>[] = [
  {
    accessorKey: "website",
    header: "Website",
  },
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => {
      const status = row.getValue("status") as string;

      return (
        <div>
          {status.charAt(0).toUpperCase() + status.slice(1)}
        </div>
      );
    },
  },
  {
    accessorKey: "uptime",
    header: "Uptime",
    cell: ({ row }) => {
      const uptime = row.getValue("uptime") as string[];

      return (
        <div className="flex items-center gap-1">
          {
            uptime.map((status, index) => {
              if (status == "up") {
                return <span className="bg-green-400 h-full w-4 p-1" key={index}></span>
              } else if (status == "down") {
                return <span className="bg-red-400 h-full w-4 p-1" key={index}></span>
              } else if (status == "unknown") {
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
    accessorKey: "lastChecked",
    header: "Last Checked",
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const monitor = row.original

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
