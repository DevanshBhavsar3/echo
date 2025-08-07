'use client'

import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import Link from 'next/link'
import { GoogleIcon } from './assets/google'
import { GithubIcon } from './assets/github'
import { useActionState } from 'react'
import { login } from '@/app/actions/auth'

export function LoginForm({
    className,
    ...props
}: React.ComponentProps<'div'>) {
    const [state, action, pending] = useActionState(login, null)

    return (
        <div className={cn('flex flex-col gap-6', className)} {...props}>
            <Card className="rounded-none">
                <CardHeader className="text-center font-sans">
                    <CardTitle className="text-2xl font-bold">
                        Welcome Back
                    </CardTitle>
                    <CardDescription className="text-muted-foreground text-balance">
                        Login to your Echo account.
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <form action={action}>
                        <div className="flex flex-col gap-6">
                            <div className="grid gap-3">
                                <Label htmlFor="email">Email</Label>
                                <Input
                                    id="email"
                                    type="email"
                                    name="email"
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
                            <div className="after:border-border relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t">
                                <span className="bg-card text-muted-foreground relative z-10 px-2">
                                    Or continue with
                                </span>
                            </div>
                            <div className="flex flex-col gap-4">
                                <Button variant="outline" className="w-full">
                                    <GoogleIcon />
                                    Login with Google
                                </Button>
                                <Button variant="outline" className="w-full">
                                    <GithubIcon />
                                    Login with GitHub
                                </Button>
                            </div>
                        </div>
                        <div className="mt-4 text-center text-sm">
                            Don&apos;t have an account?{' '}
                            <Link
                                href="/register"
                                className="underline underline-offset-4"
                            >
                                Register
                            </Link>
                        </div>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}
