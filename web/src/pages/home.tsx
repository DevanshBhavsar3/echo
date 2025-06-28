import { Navbar } from "@/components/navbar";
import { Button } from "@/components/ui/button";
import { HoverBorderGradient } from "@/components/ui/hover-border-gradient";
import { NavLink } from "react-router-dom";

export default function HomePage() {
  return (
    <div className="min-h-svh flex flex-col justify-center items-center p-6 md:p-10 relative">
      <Navbar />
      <main className="w-full max-w-md md:max-w-7xl grid gap-6">
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
            <NavLink to={"/signup"}>
              <Button size={"sm"}>Sign up</Button>
            </NavLink>
            <NavLink to={"/learn"} className="hover:underline underline-offset-4">
              Learn more
            </NavLink>
          </div>
          <span className="text-muted-foreground text-sm">
            No credit card required
          </span>
        </div>
      </main>
    </div>
  )
}
