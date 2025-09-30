import { clsx, type ClassValue } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

// Convert monitor.frequency (e.g., "30s", "1m", "3m", "5m") to milliseconds
export function frequencyToMs(freq: string): number {
    const match = freq.match(/^(\d+)([sm])$/)
    if (!match) {
        return 30000
    }

    const value = parseInt(match[1], 10)

    const unit = match[2]
    if (unit === 's') {
        return value * 1000
    }

    if (unit === 'm') {
        return value * 60 * 1000
    }

    return 30000
}
