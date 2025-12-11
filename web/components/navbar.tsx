'use client'

import Link from 'next/link'
import { Button } from './ui/button'
import { useAuth } from './providers/auth-provider'
import { motion, useMotionValueEvent, useScroll } from 'motion/react'
import { useState } from 'react'

export function Navbar() {
    const { scrollY } = useScroll()
    const [scrolled, setScrolled] = useState(false)
    const { user } = useAuth()

    useMotionValueEvent(scrollY, 'change', (latest) => {
        if (latest > 100) {
            setScrolled(true)
        } else if (latest < 100) {
            setScrolled(false)
        }
    })

    return (
        <motion.div
            initial={{
                y: -100,
            }}
            animate={{
                y: 0,
            }}
            layoutId="nav"
            className={`border-border top-0 z-30 flex w-full flex-col items-center justify-center py-5 ${scrolled && 'bg-background sticky border-b'}`}
        >
            <nav className="z-10 flex w-full max-w-7xl flex-col items-center justify-between gap-6 md:flex-row">
                <Link
                    href="/"
                    className="flex w-full items-center gap-3 font-mono"
                >
                    Echo
                </Link>
                {/*<a
                    href="https://github.com/DevanshBhavsar3/echo"
                    target="_blank"
                >
                    <Button variant={'ghost'} size={'icon'}>
                        <Github />
                    </Button>
                </a>*/}
                {/*<ModeToggle />*/}

                {user ? (
                    <Link
                        href={
                            user.isAdmin
                                ? '/dashboard/admin'
                                : '/dashboard/monitors'
                        }
                    >
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
        </motion.div>
    )
}
