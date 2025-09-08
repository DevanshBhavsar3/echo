import type { Metadata } from 'next'
import './globals.css'
import { ThemeProvider } from '@/components/theme-provider'
import { Geist_Mono, Inter } from 'next/font/google'
import { Toaster } from '@/components/ui/sonner'

const inter = Inter({
    weight: ['500', '700'],
    subsets: ['latin'],
    variable: '--font-inter',
})

const geistMono = Geist_Mono({
    subsets: ['latin'],
    variable: '--font-geist-mono',
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
            className={`${inter.variable} ${geistMono.variable}`}
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
