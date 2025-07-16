import { Navbar } from "@/components/navbar";
import { Button } from "@/components/ui/button";
import { HoverBorderGradient } from "@/components/ui/hover-border-gradient";
import Link from "next/link";
import { auth } from "./auth";

export default async function HomePage() {
  const user = await auth()
  return (
    <div className="flex flex-col justify-center items-center px-6 md:px-10 relative">
      <Navbar />
      <main className="min-h-svh w-full max-w-md md:max-w-7xl flex flex-col justify-center items-center">
        <div className="w-full h-full grid gap-6">

          {"USERNAME: " + user?.user.name}
          <HoverBorderGradient
            as={"text"}
            className="flex items-center bg-background text-foreground"
          >
            <p className="text-xs">Echo is just released! 🎉</p>
          </HoverBorderGradient>

          <div className="grid gap-3">
            <h1 className="max-w-md scroll-m-20 text-left font-mono text-4xl tracking-tight">
              Your Servers Speak, We Listen.
            </h1>
            <p className="text-muted-foreground text-balance max-w-md">
              Stop worrying about your servers. Get comprehensive uptime
              monitoring without spending single penny.
            </p>
          </div>

          <div className="grid gap-3">
            <div className="grid grid-cols-2 gap-3 w-fit items-center">
              <Link href={"/register"}>
                <Button size={"sm"}>Register</Button>
              </Link>
              <Link href={"/learn"} className="hover:underline underline-offset-4">
                Learn more
              </Link>
            </div>
            <span className="text-muted-foreground text-sm">
              No credit card required
            </span>
          </div>
        </div>
      </main>

      <section>
        Features
      </section>
    </div>
  )
}
