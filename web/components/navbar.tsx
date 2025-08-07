import { Github } from 'lucide-react'
import { Button } from './ui/button'
import Link from 'next/link'
import { ModeToggle } from './theme-toggle'
import { HeaderLine } from './header-line'
import { auth } from '@/app/auth'

export async function Navbar() {
    const user = await auth()

    return (
        <div className="bg-background fixed top-0 flex w-full flex-col items-center justify-center">
            <HeaderLine />
            <nav className="flex w-full max-w-7xl flex-col items-center justify-between gap-6 px-2 py-4 md:flex-row">
                <div className="flex w-full items-center gap-3">
                    <Link href={'/'}>
                        <Button variant={'link'} className="text-md">
                            Echo
                        </Button>
                    </Link>
                    <Link href={'/about'}>
                        <Button
                            variant={'link'}
                            className="text-muted-foreground"
                        >
                            About
                        </Button>
                    </Link>
                    <Link href={'/learn'}>
                        <Button
                            variant={'link'}
                            className="text-muted-foreground"
                        >
                            Learn
                        </Button>
                    </Link>
                </div>
                <a
                    href="https://github.com/DevanshBhavsar3/echo"
                    target="_blank"
                >
                    <Button variant={'ghost'} size={'icon'}>
                        <Github />
                    </Button>
                </a>
                <ModeToggle />

                {user?.user.id ? (
                    <Link href={'/dashboard'}>
                        <Button variant={'outline'}>Dashboard</Button>
                    </Link>
                ) : (
                    <div className="flex items-center gap-3">
                        <Link href={'/login'}>
                            <Button variant={'outline'}>Login</Button>
                        </Link>
                        <Link href={'/register'}>
                            <Button>Register</Button>
                        </Link>
                    </div>
                )}
            </nav>
        </div>
    )
}
