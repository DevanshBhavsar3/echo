export const API_URL =
    process.env.NODE_ENV === 'production'
        ? process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001/api/v1'
        : process.env.NEXT_PUBLIC_DOCKER_API_URL || 'http://api:3001/api/v1'
