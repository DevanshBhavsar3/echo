import { clsx, type ClassValue } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

export function getDateBeforeDays(day: number) {
    return new Date(
        new Date().setDate(new Date().getDate() - day),
    ).toISOString()
}
