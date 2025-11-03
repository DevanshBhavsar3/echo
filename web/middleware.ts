import { cookies } from 'next/headers'
import { NextRequest, NextResponse } from 'next/server'

const protectedRoutes = ['/dashboard']
const publicRoutes = ['/login', '/register', '/admin']

export async function middleware(request: NextRequest) {
    const { pathname } = request.nextUrl
    const isProtectedRoute = protectedRoutes.some((route) =>
        pathname.startsWith(route),
    )
    const isPublicRoute = publicRoutes.includes(pathname)

    const token = (await cookies()).get('token')?.value

    if (!token && isProtectedRoute) {
        return NextResponse.redirect(new URL('/login', request.url))
    }

    if (token && isPublicRoute && !request.url.startsWith('/dashboard')) {
        return NextResponse.redirect(
            new URL('/dashboard/monitors', request.url),
        )
    }

    return NextResponse.next()
}

export const config = {
    matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}
