"use server"

import axios from "axios";
import { API_URL } from "../constants";

export async function fetchRegions() {
  try {
    const res = await axios.get(`${API_URL}/region`);

    return res.data;
  } catch (error) {
    console.error("Error fetching regions:", error);
    return [];
  }
}
