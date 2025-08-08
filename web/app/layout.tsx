import type { Metadata } from 'next'
import './globals.css'
import { ThemeProvider } from '@/components/theme-provider'
import { Geist, Space_Mono } from 'next/font/google'
import { Toaster } from '@/components/ui/sonner'

const spaceMono = Space_Mono({
    weight: ['400', '700'],
    subsets: ['latin'],
    variable: '--font-space-mono',
})

const geist = Geist({
    subsets: ['latin'],
    variable: '--font-geist-sans',
})

export const metadata: Metadata = {
    title: 'Echo - Uptime Monitoring',
    description:
        'Echo is a simple uptime monitoring service that helps you keep track of your websites and services.',
}

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode
}>) {
    return (
        <html
            lang="en"
            className={`${spaceMono.className} ${geist.className}`}
            suppressHydrationWarning
        >
            <body>
                <ThemeProvider
                    attribute="class"
                    defaultTheme="system"
                    enableSystem
                    disableTransitionOnChange
                >
                    {children}
                    <Toaster richColors />
                </ThemeProvider>
            </body>
        </html>
    )
}
