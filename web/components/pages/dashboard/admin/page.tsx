'use client'

import { logout } from '@/app/actions/auth'
import { createRegion } from '@/app/actions/region'
import { redirect } from 'next/navigation'
import { useActionState, useRef } from 'react'
import ReactCountryFlag from 'react-country-flag'
import { useAuth } from '../../../providers/auth-provider'
import {
    Breadcrumb,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbList,
    BreadcrumbPage,
    BreadcrumbSeparator,
} from '../../../ui/breadcrumb'
import { Button } from '../../../ui/button'
import {
    Field,
    FieldDescription,
    FieldError,
    FieldLabel,
} from '../../../ui/field'
import { Input } from '../../../ui/input'
import { Region } from '../monitors/data-table'

export function AdminPage({ regions }: { regions: Region[] }) {
    const { user } = useAuth()

    const [regionstate, regionAction, regionPending] = useActionState(
        createRegion,
        null,
    )

    const regionRef = useRef<HTMLInputElement>(null)

    if (user && !user.isAdmin) {
        redirect('/dashboard/monitors')
    }

    return (
        <>
            <Breadcrumb>
                <BreadcrumbList>
                    <BreadcrumbItem>
                        <BreadcrumbLink href="/">Home</BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>Admin</BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>
            <div className="grid gap-5">
                <form className="grid gap-3" action={regionAction}>
                    <Field>
                        <FieldLabel htmlFor="code">Region Name</FieldLabel>
                        <Input
                            name="code"
                            type="text"
                            placeholder="IN"
                            className="w-fit"
                            ref={regionRef}
                        />
                        <FieldDescription>
                            Add new ISO 2 code for the region.
                        </FieldDescription>
                        {regionstate?.error && (
                            <FieldError>{regionstate.error}</FieldError>
                        )}
                    </Field>
                    <Button
                        type="submit"
                        className="w-fit"
                        disabled={regionPending}
                    >
                        Submit
                    </Button>
                </form>
                <div className="grid gap-3">
                    <h1 className="text-sm">Available Regions: </h1>
                    {regions?.map((r) => (
                        <div
                            key={r.regionId}
                            className="flex items-center gap-2"
                        >
                            <ReactCountryFlag
                                countryCode={r.regionName.toUpperCase()}
                                svg
                            />
                            <p className="text-muted-foreground text-xs">
                                {r.regionName}
                            </p>
                        </div>
                    ))}
                </div>
                <Button className="w-fit" onClick={logout}>
                    Log out
                </Button>
            </div>
        </>
    )
}
