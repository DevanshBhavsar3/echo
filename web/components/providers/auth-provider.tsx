'use client'

import { getUser } from '@/app/actions/auth'
import { User } from '@/lib/types'
import { usePathname } from 'next/navigation'
import { createContext, useContext, useEffect, useState } from 'react'

export const AuthContext = createContext<{
    user: User | null
    updateUser: () => void
    clearUser: () => void
}>({
    user: null,
    updateUser: () => {},
    clearUser: () => {},
})

export const useAuth = () => {
    return useContext(AuthContext)
}

// Clear user when log out.
export function AuthProvider({ children }: { children: React.ReactNode }) {
    const pathname = usePathname()
    const [user, setUser] = useState<User | null>(null)

    async function updateUser() {
        const isLoggedIn = document.cookie.includes('token')

        if (isLoggedIn && !user?.isAdmin) {
            const user = await getUser()

            if (user) {
                setUser(user)
            }
        }
    }

    function clearUser() {
        setUser(null)
    }

    useEffect(() => {
        updateUser()

        return () => {
            setUser(null)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [pathname])

    return (
        <AuthContext value={{ user, updateUser, clearUser }}>
            {children}
        </AuthContext>
    )
}
