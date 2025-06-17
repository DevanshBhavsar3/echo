import { ThemeProvider } from "./components/theme-provider";
import { Navbar } from "./components/navbar";
import { Button } from "./components/ui/button";
import { HoverBorderGradient } from "./components/ui/hover-border-gradient";

function App() {
  return (
    <ThemeProvider>
      <Navbar />
      <main className="mx-auto flex min-h-screen max-w-7xl items-center justify-between">
        <div className="flex flex-col gap-2">
          <HoverBorderGradient
            as={"text"}
            className="flex items-center bg-white text-black dark:bg-black dark:text-white"
          >
            <p className="text-xs">Echo is just released! 🎉</p>
          </HoverBorderGradient>
          <h1 className="max-w-md scroll-m-20 text-left font-mono text-4xl tracking-tight">
            Your Servers Speak, We Listen.
          </h1>
          <p className="text-muted-foreground max-w-md text-lg font-semibold text-balance">
            Stop worrying about your servers. Get comprehensive uptime
            monitoring without spending single penny.
          </p>
          <div className="flex items-center gap-2">
            <Button size={"sm"}>Sign in</Button>

            <Button variant={"link"}>Learn more</Button>
          </div>
          <p className="text-muted-foreground text-xs">
            No credit card required
          </p>
        </div>

        <div></div>
      </main>
    </ThemeProvider>
  );
}

export default App;
