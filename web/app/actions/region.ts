'use server'

import axios from 'axios'
import { API_URL } from '../constants'
import { Region } from '../dashboard/monitors/data-table'

export async function fetchRegions() {
    try {
        const res = await axios.get(`${API_URL}/region`)

        return res.data.regions as Region[]
    } catch (error) {
        console.error('Error fetching regions:', error)
        return []
    }
}
