import { Github } from "lucide-react";
import { ModeToggle } from "./mode-toggle";
import { Button } from "./ui/button";
import { NavLink } from "react-router-dom";

export function Navbar() {
  return (
    <nav className="fixed top-0 left-1/2 py-4 px-2 flex flex-col md:flex-row w-full max-w-7xl -translate-x-1/2 items-center justify-between gap-6">
      <div className="flex w-full items-center gap-3">
        <NavLink to={"/"}>
          <Button variant={"link"} className="text-md">
            Echo
          </Button>
        </NavLink>
        <NavLink to={"/about"}>
          <Button variant={"link"} className="text-muted-foreground">About</Button>
        </NavLink>
        <NavLink to={"/learn"}>
          <Button variant={"link"} className="text-muted-foreground">Learn</Button>
        </NavLink>
      </div>
      <a href="https://github.com/DevanshBhavsar3/echo" target="_blank">
        <Button variant={"ghost"} size={"icon"}>
          <Github />
        </Button>
      </a>
      <ModeToggle />
      <div className="flex items-center gap-3">
        <NavLink to={"/login"}>
          <Button size={"sm"}>Log in</Button>
        </NavLink>
        <NavLink to={"/signup"}>
          <Button size={"sm"} variant={"outline"}>
            Sign up
          </Button>
        </NavLink>
      </div>
    </nav >
  );
}
