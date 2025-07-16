"use client"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import Link from "next/link"
import { GoogleIcon } from "./assets/google"
import { GithubIcon } from "./assets/github"
import { useActionState } from "react"
import { register } from "@/app/actions/auth"

export function RegisterForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const [state, action, pending] = useActionState(register, undefined)

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card className="rounded-none">
        <CardHeader className="text-center font-sans">
          <CardTitle className="text-2xl font-bold">Welcome to Echo</CardTitle>
          <CardDescription className="text-muted-foreground text-balance">
            Register your new Echo account.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form action={action}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-3">
                <Label htmlFor="name">Name</Label>
                <Input
                  id="name"
                  name="name"
                  type="text"
                  placeholder="John Doe"
                  defaultValue={state?.data?.name.toString() || ""}
                  required
                />
                {state?.errors?.name && <p className="font-sans text-muted-foreground text-sm">{state.errors.name}</p>}
              </div>
              <div className="grid gap-3">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  placeholder="me@example.com"
                  defaultValue={state?.data?.email.toString() || ""}
                  required
                />
                {state?.errors?.email && <p className="font-sans text-muted-foreground text-sm">{state.errors.email}</p>}
              </div>
              <div className="grid gap-3">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  name="password"
                  type="password"
                  defaultValue={state?.data?.password.toString() || ""}
                  required />
                {state?.errors?.password && (
                  <div>
                    <p className="font-sans text-muted-foreground text-sm">Password must:</p>
                    <ul>
                      {state.errors.password.map((error) => (
                        <li key={error} className="font-sans text-muted-foreground text-sm">- {error}</li>
                      ))}
                    </ul>
                  </div>
                )}
              </div>
              {
                state?.error && (
                  <p className="text-sm font-sans text-muted-foreground">
                    {state.error}
                  </p>
                )
              }
              <Button type="submit" disabled={pending} className="w-full">
                Register
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
              Already have an account?{" "}
              <Link href="/login" className="underline underline-offset-4">
                Login
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
