'use client'

import { Background } from '@/components/background'
import { Navbar } from '@/components/navbar'
import { useAuth } from '@/components/providers/auth-provider'
import { Button } from '@/components/ui/button'
import { motion } from 'motion/react'
import Link from 'next/link'

export default function HomePage() {
    const { user } = useAuth()

    return (
        <div className="relative flex flex-col items-center justify-center space-y-12">
            <Navbar />
            <main className="relative flex w-full max-w-md flex-col items-center md:max-w-7xl">
                <div
                    className="absolute inset-0 z-0"
                    style={{
                        background: '#ffffff',
                        backgroundImage: `
                      radial-gradient(
                        circle at top center,
                        rgba(70, 130, 180, 0.5),
                        transparent 70%
                      )
                    `,
                        filter: 'blur(80px)',
                        backgroundRepeat: 'no-repeat',
                    }}
                />
                <motion.div
                    initial={{
                        opacity: 0,
                        y: 50,
                    }}
                    animate={{
                        opacity: 1,
                        y: 0,
                    }}
                    className="z-10 flex w-full flex-col items-center justify-center gap-8"
                >
                    <div className="bg-primary/85 flex items-center rounded-full border border-sky-800 px-4 py-1">
                        <p className="text-primary-foreground text-xs font-medium">
                            Echo is just released! 🎉
                        </p>
                    </div>

                    <h1 className="text-primary-foreground max-w-4xl scroll-m-20 text-center text-7xl font-medium text-balance">
                        Your Servers Speak, We Listen.
                    </h1>
                    <p className="max-w-lg text-center text-sm font-medium text-balance text-neutral-600">
                        Stop worrying about your servers. Get comprehensive
                        uptime monitoring without spending a single penny.
                    </p>

                    <div className="grid gap-3">
                        <div className="grid w-fit grid-cols-2 items-center gap-3">
                            {user ? (
                                <Link
                                    href={
                                        user.isAdmin
                                            ? '/dashboard/admin'
                                            : '/dashboard/monitors'
                                    }
                                >
                                    <Button size={'sm'}>Dashboard</Button>
                                </Link>
                            ) : (
                                <Link href={'/register'}>
                                    <Button>Register</Button>
                                </Link>
                            )}
                            <Link
                                href={'/learn'}
                                className="text-sm underline-offset-4 hover:underline"
                            >
                                Learn more
                            </Link>
                        </div>
                    </div>
                </motion.div>
                <div className="h-80">
                    <Background />
                </div>
            </main>

            <section className="mb-1000">Features</section>
        </div>
    )
}
