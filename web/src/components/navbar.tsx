import { Github } from "lucide-react";
import { ModeToggle } from "./mode-toggle";
import { Button } from "./ui/button";

export function Navbar() {
  return (
    <nav className="absolute left-1/2 my-4 flex min-w-7xl -translate-x-1/2 items-center justify-between gap-8">
      <div className="flex w-full items-center gap-5 font-mono">
        <Button variant={"link"} className="font-bold">
          Echo
        </Button>
        <Button variant={"link"}>Home</Button>
        <Button variant={"link"}>About</Button>
      </div>

      <Button variant={"ghost"} size={"icon"}>
        <Github />
      </Button>

      <ModeToggle />
      <div className="flex items-center gap-3">
        <Button size={"sm"}>Sign in</Button>
        <Button size={"sm"} variant={"outline"}>
          Sign up
        </Button>
      </div>
    </nav>
  );
}
