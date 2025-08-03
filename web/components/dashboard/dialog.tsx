"use client";

import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "../ui/dialog";
import { Input } from "../ui/input";
import { Checkbox } from "../ui/checkbox";
import ReactCountryFlag from "react-country-flag";
import { pingWebsite } from "@/app/actions/website";
import { startTransition, useActionState, useEffect, useState } from "react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select";
import { useDebounce } from "@/hooks/use-debounce";
import { CircleCheckBig, CircleX, LoaderCircle } from "lucide-react";
import { Monitors } from "@/app/dashboard/data-table";
import { fetchRegions } from "@/app/actions/region";

type Region = {
  id: string;
  name: string;
};

const frequencies = [
  { value: "30s", label: "30 Seconds" },
  { value: "1m", label: "1 Minute" },
  { value: "3m", label: "3 Minutes" },
  { value: "5m", label: "5 Minutes" },
];

interface DialogProps {
  label: string;
  description?: string;
  data?: Monitors;
  onSubmitAction: (_: unknown, formData: FormData) => Promise<any>;
  children?: React.ReactNode;
}

export function DialogBox({ label, description, data, onSubmitAction, children }: DialogProps) {
  const [open, setOpen] = useState(false);

  const [state, action, pending] = useActionState(onSubmitAction, null);
  const [urlState, urlAction, urlPending] = useActionState(pingWebsite, null);
  const [regionState, regionAction, regionPending] = useActionState(fetchRegions, null);

  useEffect(() => {
    if (open) {
      startTransition(() => {
        regionAction();
      })
    }
  }, [open]);

  const [url, setUrl] = useState(data?.url || "https://");
  const debouncedUrl = useDebounce(url, 500);

  useEffect(() => {
    if (debouncedUrl && debouncedUrl !== "https://") {
      startTransition(() => {
        urlAction(debouncedUrl);
      });
    }
  }, [debouncedUrl, urlAction]);

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      {children}
      <DialogContent>
        <DialogHeader className="font-sans">
          <DialogTitle>{label}</DialogTitle>
          <DialogDescription>
            {description}
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
                  value={url}
                  placeholder="https://example.com"
                  required
                  onChange={(e) => {
                    setUrl(e.target.value);
                  }}
                />

                {url && url.length > "https://".length ? (
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
                ) :
                  null
                }

              </div>
              {state?.errors?.url && (
                <p className="font-sans text-muted-foreground text-sm">
                  {state.errors.url}
                </p>
              )}
            </div>
            <div className="grid gap-3">
              <Label htmlFor="frequencies">Frequency</Label>
              <Select defaultValue={(data && data?.frequency) || "3m"} name="frequency">
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
              {
                regionPending ? (
                  <LoaderCircle size={18} className="animate-spin opacity-50" />
                ) : regionState?.regions && regionState.regions.length > 0 ? (
                  <div className="grid gap-3">
                    <Label htmlFor="regions">Regions</Label>
                    <div className="flex flex-wrap gap-3">
                      {regionState.regions.map((region: Region) => (
                        <div key={region.id} className="flex items-start gap-3">
                          <Checkbox
                            name="regions"
                            value={region.name}
                            defaultChecked={(data && data.regions.find(r => r == region.name)) ? true : false}
                          />
                          <ReactCountryFlag
                            countryCode={region.name.toUpperCase()}
                            svg
                          />
                        </div>
                      ))}
                    </div>
                  </div>
                ) : (
                  <p className="text-sm font-sans text-muted-foreground">
                    No regions available.
                  </p>
                )
              }
              {
                state?.errors?.regions && (
                  <p className="font-sans text-muted-foreground text-sm">
                    {state.errors.regions}
                  </p>
                )
              }
            </div>
            {state?.error && (
              <p className="text-sm font-sans text-muted-foreground">
                {state.error}
              </p>
            )}
          </div>
          <DialogFooter className="sm:justify-start">
            <Button type="submit" disabled={pending}>
              {label}
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
