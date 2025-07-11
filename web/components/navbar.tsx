import { Github } from "lucide-react";
import { Button } from "./ui/button";
import Link from "next/link";
import { ModeToggle } from "./theme-toggle";

export function Navbar() {
  return (
    <nav className="fixed top-0 left-1/2 py-4 px-2 flex flex-col md:flex-row w-full max-w-7xl -translate-x-1/2 items-center justify-between gap-6">
      <div className="flex w-full items-center gap-3">
        <Link href={"/"}>
          <Button variant={"link"} className="text-md">
            Echo
          </Button>
        </Link>
        <Link href={"/about"}>
          <Button variant={"link"} className="text-muted-foreground">About</Button>
        </Link>
        <Link href={"/learn"}>
          <Button variant={"link"} className="text-muted-foreground">Learn</Button>
        </Link>
      </div>
      <a href="https://github.com/DevanshBhavsar3/echo" target="_blank">
        <Button variant={"ghost"} size={"icon"}>
          <Github />
        </Button>
      </a>
      <ModeToggle />

      <div className="flex items-center gap-3">
        <Link href={"/login"}>
          <Button size={"sm"}>Log in</Button>
        </Link>
        <Link href={"/signup"}>
          <Button size={"sm"} variant={"outline"}>
            Sign up
          </Button>
        </Link>
      </div>
    </nav >
  );
}
