import { ThemeProvider } from "./components/theme-provider"
import { Navbar } from "./components/navbar"
import { Button } from "./components/ui/button"
import { HoverBorderGradient } from "./components/ui/hover-border-gradient"

function App() {
  return (
    <ThemeProvider>
      <Navbar />
      <main className="min-h-screen max-w-7xl mx-auto flex items-center justify-between">
        <div className="flex flex-col gap-2">
          <HoverBorderGradient as={"text"}
            className="dark:bg-black bg-white text-black dark:text-white flex items-center"
          >
            <p className="text-xs">Echo is just  released! 🎉</p>
          </HoverBorderGradient>
          <h1 className="font-mono scroll-m-20 text-4xl tracking-tight text-left max-w-md">Your Servers Speak, We Listen.</h1>
          <p className="text-lg font-semibold max-w-md text-balance text-muted-foreground">
            Stop worrying about your servers. Get comprehensive uptime monitoring without spending single penny.
          </p>
          <div className="flex items-center gap-2">
            <Button size={"sm"}>Sign in</Button>

            <Button variant={"link"}>
              Learn more
            </Button>
          </div>
          <p className="text-muted-foreground text-xs">No credit card required</p>
        </div>

        <div>
        </div>
      </main>

    </ThemeProvider>
  )
}

export default App
