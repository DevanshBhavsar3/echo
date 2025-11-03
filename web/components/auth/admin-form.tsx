'use client'

import { adminLogin } from '@/app/actions/auth'
import { cn } from '@/lib/utils'
import { useActionState } from 'react'
import { Button } from '../ui/button'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '../ui/card'
import { Input } from '../ui/input'
import { Label } from '../ui/label'

export function AdminForm({
    className,
    ...props
}: React.ComponentProps<'div'>) {
    const [state, action, pending] = useActionState(adminLogin, null)

    return (
        <div className={cn('flex flex-col gap-6', className)} {...props}>
            <Card className="rounded-none">
                <CardHeader className="text-center font-sans">
                    <CardTitle className="text-2xl font-bold">
                        Hello, Admin!
                    </CardTitle>
                    <CardDescription className="text-muted-foreground text-balance">
                        Login to your Echo admin account.
                    </CardDescription>
                </CardHeader>
                <CardContent className="grid gap-4">
                    <form action={action}>
                        <div className="flex flex-col gap-6">
                            <div className="grid gap-3">
                                <Label htmlFor="username">Username</Label>
                                <Input
                                    id="username"
                                    type="text"
                                    name="username"
                                    placeholder="me@example.com"
                                    required
                                />
                            </div>
                            <div className="grid gap-3">
                                <Label htmlFor="password">Password</Label>
                                <Input
                                    id="password"
                                    name="password"
                                    type="password"
                                    required
                                />
                            </div>
                            {state?.error && (
                                <p className="text-muted-foreground font-sans text-sm">
                                    {state.error}
                                </p>
                            )}
                            <Button
                                type="submit"
                                disabled={pending}
                                className="w-full"
                            >
                                Login
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}
