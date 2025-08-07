'use server'

import { websiteSchema } from '@/lib/types'
import axios, { AxiosError } from 'axios'
import { API_URL } from '../constants'
import { auth } from '../auth'
import { redirect } from 'next/navigation'
import { revalidatePath } from 'next/cache'

export async function createWebsite(_: unknown, formData: FormData) {
  const user = await auth()

  if (!user?.token) {
    redirect('/login')
  }

  const parsedData = websiteSchema.safeParse({
    url: formData.get('url'),
    frequency: formData.get('frequency'),
    regions: formData.getAll('regions') as string[],
  })

  if (!parsedData.success) {
    return {
      errors: parsedData.error.flatten().fieldErrors,
    }
  }

  try {
    await axios.post(`${API_URL}/website`, parsedData.data, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    })

    revalidatePath('/dashboard')
  } catch (error) {
    if (error instanceof AxiosError) {
      return {
        error:
          error.response?.data?.error ||
          'An error occurred while creating the website.',
      }
    }

    return {
      error: 'An unexpected error occurred while creating the website.',
    }
  }
}

export async function deleteWebsite(id: string) {
  const user = await auth()

  if (!user?.token) {
    redirect('/login')
  }

  try {
    await axios.delete(`${API_URL}/website/${id}`, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    })

    revalidatePath('/dashboard')
  } catch (error) {
    console.error('Error deleting website:', error)
  }
}

export async function editWebsite(websiteId: string, formData: FormData) {
  const user = await auth()

  if (!user?.token) {
    redirect('/login')
  }

  const parsedData = websiteSchema.safeParse({
    url: formData.get('url'),
    frequency: formData.get('frequency'),
    regions: formData.getAll('regions') as string[],
  })

  if (!parsedData.success) {
    return {
      errors: parsedData.error.flatten().fieldErrors,
    }
  }

  try {
    await axios.put(`${API_URL}/website/${websiteId}`, parsedData.data, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    })

    revalidatePath('/dashboard')
  } catch (error) {
    if (error instanceof AxiosError) {
      return {
        error:
          error.response?.data?.error ||
          'An error occurred while editing the website.',
      }
    }

    return {
      error: 'An unexpected error occurred while editing the website.',
    }
  }
}

export async function pingWebsite(_: unknown, url: string) {
  try {
    await axios.head(url, {
      withCredentials: false,
      timeout: 5000,
    })

    return {
      status: true,
    }
  } catch (e) {
    if (e instanceof AxiosError) {
      return {
        status: false,
        error: e.response?.statusText || 'Failed to ping the website.',
      }
    }

    return {
      status: false,
      error: 'An unexpected error occurred while pinging the website.',
    }
  }
}
