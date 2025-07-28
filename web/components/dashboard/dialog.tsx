"use client";

import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "../ui/dialog";
import { Input } from "../ui/input";
import { Checkbox } from "../ui/checkbox";
import ReactCountryFlag from "react-country-flag";
import { createWebsite, pingWebsite } from "@/app/actions/website";
import { startTransition, useActionState, useEffect, useState } from "react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select";
import { useDebounce } from "@/hooks/use-debounce";
import { CircleCheckBig, CircleX, LoaderCircle } from "lucide-react";

const frequencies = [
  { value: "30s", label: "30 Seconds" },
  { value: "1m", label: "1 Minute" },
  { value: "3m", label: "3 Minutes", default: true },
  { value: "5m", label: "5 Minutes" },
];

export function AddMonitorDialog() {
  const [state, action, pending] = useActionState(createWebsite, null);
  const [urlState, urlAction, urlPending] = useActionState(pingWebsite, null);

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" className="hidden sm:flex">
          Add Monitor
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader className="font-sans">
          <DialogTitle>Add Monitor</DialogTitle>
          <DialogDescription>
            Add a new monitor to track the uptime of your website.
          </DialogDescription>
        </DialogHeader>
        <form action={action} className="grid gap-6">
          <div className="grid gap-6">
            <div className="grid gap-3">
              <Label htmlFor="url">Url</Label>

              <div className="flex items-center relative">
                <Input
                  id="url"
                  type="text"
                  name="url"
                  defaultValue={"https://"}
                  placeholder="https://example.com"
                  required
                  onBlur={(e) => startTransition(() => {
                    const url = e.target.value;
                    if (url) {
                      urlAction(url);
                    }
                  })}
                />

                <span className="absolute right-3">
                  {
                    urlPending ? (
                      <LoaderCircle size={18} className="animate-spin opacity-50" />
                    ) :
                      urlState?.status ? (
                        <CircleCheckBig size={18} className="text-green-500" />
                      ) : (
                        <CircleX size={18} className="text-red-500" />
                      )
                  }

                </span>

              </div>
              {state?.errors?.url && (
                <p className="font-sans text-muted-foreground text-sm">
                  {state.errors.url}
                </p>
              )}
            </div>
            <div className="grid gap-3">
              <Label htmlFor="frequencies">Frequency</Label>
              <Select defaultValue="3m" name="frequency">
                <SelectTrigger className="w-50">
                  <SelectValue placeholder="Select frequency" />
                </SelectTrigger>
                <SelectContent>
                  {frequencies.map((frequency) => (
                    <SelectItem
                      key={frequency.value}
                      value={frequency.value}
                    >
                      {frequency.label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div className="grid gap-3">
              <Label htmlFor="regions">Regions</Label>
              <div className="flex items-start gap-3">
                <Checkbox name="regions" value="IN" defaultChecked />
                <ReactCountryFlag
                  countryCode={'IN'}
                  svg
                />
              </div>
            </div>
            {state?.error && (
              <p className="text-sm font-sans text-muted-foreground">
                {state.error}
              </p>
            )}
          </div>
          <DialogFooter className="sm:justify-start">
            <Button type="submit" disabled={pending}>
              Add Monitor
            </Button>
            <DialogClose asChild>
              <Button variant="outline" disabled={pending}>
                Cancel
              </Button>
            </DialogClose>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog >
  );
}
