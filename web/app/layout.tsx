import type { Metadata } from 'next'
import { Geist_Mono, Inter } from 'next/font/google'
import './globals.css'
import { Providers } from './Providers'

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
                <Providers>{children}</Providers>
            </body>
        </html>
    )
}
