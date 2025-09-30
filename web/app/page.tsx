import { Navbar } from '@/components/navbar'
import { Button } from '@/components/ui/button'
import { HoverBorderGradient } from '@/components/ui/hover-border-gradient'
import { isLoggedIn } from '@/lib/auth'
import Link from 'next/link'

export default async function HomePage() {
    const loginStatus = await isLoggedIn()

    return (
        <div className="relative flex flex-col items-center justify-center px-6 md:px-10">
            <Navbar />
            <main className="flex min-h-svh w-full max-w-md flex-col items-center justify-center md:max-w-7xl">
                <div className="grid h-full w-full gap-6">
                    <HoverBorderGradient
                        as={'text'}
                        className="bg-background text-foreground flex items-center"
                    >
                        <p className="text-xs">Echo is just released! 🎉</p>
                    </HoverBorderGradient>

                    <div className="grid gap-3">
                        <h1 className="max-w-md scroll-m-20 text-left font-mono text-4xl tracking-tight">
                            Your Servers Speak, We Listen.
                        </h1>
                        <p className="text-muted-foreground max-w-md font-sans text-balance">
                            Stop worrying about your servers. Get comprehensive
                            uptime monitoring without spending single penny.
                        </p>
                    </div>

                    <div className="grid gap-3">
                        <div className="grid w-fit grid-cols-2 items-center gap-3">
                            {loginStatus ? (
                                <Link href={'/dashboard/monitors'}>
                                    <Button size={'sm'}>Dashboard</Button>
                                </Link>
                            ) : (
                                <Link href={'/register'}>
                                    <Button size={'sm'}>Register</Button>
                                </Link>
                            )}
                            <Link
                                href={'/learn'}
                                className="underline-offset-4 hover:underline"
                            >
                                Learn more
                            </Link>
                        </div>
                        <span className="text-muted-foreground text-sm">
                            No credit card required
                        </span>
                    </div>
                </div>
            </main>

            <section>Features</section>
        </div>
    )
}
