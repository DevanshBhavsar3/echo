'use server'

import { AxiosError } from 'axios'
import { loginSchema, registerSchema, User } from '@/lib/types'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'
import apiClient from '@/lib/axios'

export async function register(_: unknown, formData: FormData) {
    const cookieStore = await cookies()

    const parsedData = registerSchema.safeParse({
        name: formData.get('name'),
        email: formData.get('email'),
        password: formData.get('password'),
    })

    if (!parsedData.success) {
        return {
            data: Object.fromEntries(formData.entries()),
            errors: parsedData.error.flatten().fieldErrors,
        }
    }

    try {
        const res = await apiClient.post(`/auth/register`, {
            ...parsedData.data,
            image:
                'https://api.dicebear.com/6.x/initials/svg?seed=' +
                parsedData.data.name,
        })

        cookieStore.set('token', res.data.token)
    } catch (error) {
        if (error instanceof AxiosError) {
            return {
                error:
                    error.response?.data?.error ||
                    'An error occurred during registration.',
            }
        }

        return {
            error: 'An unexpected error occurred.',
        }
    }

    redirect('/dashboard/monitors')
}

export async function login(_: unknown, formData: FormData) {
    const cookieStore = await cookies()
    const values = Object.fromEntries(formData.entries())

    const parsedData = loginSchema.safeParse({
        email: values['email'],
        password: values['password'],
    })

    if (!parsedData.success) {
        return { error: parsedData.error.issues[0].message }
    }

    const { email, password } = parsedData.data

    try {
        const res = await apiClient.post(`/auth/login`, {
            email,
            password,
        })

        cookieStore.set('token', res.data.token)
    } catch (error) {
        if (error instanceof AxiosError) {
            return {
                error:
                    error.response?.data?.error ||
                    'An error occurred during login.',
            }
        }

        return { error: 'An unexpected error occurred during login.' }
    }

    redirect('/dashboard/monitors')
}

export async function oauth(provider: string) {
    redirect(`${process.env.NEXT_PUBLIC_API_URL}/oauth/${provider}`)
}

export async function logout() {
    const cookieStore = await cookies()
    cookieStore.delete('token')

    redirect('/login')
}

export async function getUser(): Promise<User | null> {
    try {
        const res = await apiClient.get(`/auth/me`)

        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            console.error(error.response?.data?.error)
        }
    }

    return null
}
