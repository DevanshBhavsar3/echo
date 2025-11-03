'use client'

import { getUser } from '@/app/actions/auth'
import { User } from '@/lib/types'
import { createContext, useContext, useEffect, useState } from 'react'

export const AuthContext = createContext<{
    user: User | null
    updateUser: () => void
}>({
    user: null,
    updateUser: () => {},
})

export const useAuth = () => {
    return useContext(AuthContext)
}

// Clear user when log out.
export function AuthProvider({ children }: { children: React.ReactNode }) {
    const [user, setUser] = useState<User | null>(null)

    async function updateUser() {
        const user = await getUser()

        if (user) {
            setUser(user)
        }
    }

    useEffect(() => {
        updateUser()

        return () => {
            setUser(null)
        }
    }, [])

    return <AuthContext value={{ user, updateUser }}>{children}</AuthContext>
}
