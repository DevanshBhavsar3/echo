'use server'

import { cookies } from 'next/headers'

export async function isLoggedIn() {
    return (await cookies()).has('token')
}
