import { Github } from "lucide-react";
import { Button } from "./ui/button";
import Link from "next/link";
import { ModeToggle } from "./theme-toggle";
import { HeaderLine } from "./header-line";

export function Navbar() {
  return (
    <div className="w-full fixed top-0 flex flex-col justify-center items-center">
      <HeaderLine>
        🎉 Echo just launched!! Get unlimited uptime monitoring without any credit card.
      </HeaderLine>
      <nav className="py-4 px-2 flex flex-col md:flex-row w-full max-w-7xl items-center justify-between gap-6">
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
            <Button size={"sm"} variant={"outline"}>
              Login
            </Button>
          </Link>
          <Link href={"/register"}>
            <Button size={"sm"}>
              Register
            </Button>
          </Link>
        </div>
      </nav >
    </div>
  );
}
