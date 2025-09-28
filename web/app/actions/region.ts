'use server'

import { Region } from '../dashboard/monitors/data-table'
import apiClient from '@/lib/axios'

export async function fetchRegions() {
    try {
        const res = await apiClient.get(`/region`)

        return res.data.regions as Region[]
    } catch (error) {
        console.error('Error fetching regions:', error)
        return []
    }
}
