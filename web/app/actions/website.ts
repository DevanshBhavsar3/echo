"use server";

import { createWebsiteSchema } from "@/lib/types";
import axios, { AxiosError } from "axios";
import { API_URL } from "../constants";
import { auth } from "../auth";
import { redirect } from "next/navigation";
import { revalidatePath } from "next/cache";

export async function createWebsite(_: unknown, formData: FormData) {
  const user = await auth();

  if (!user?.token) {
    redirect("/login");
  }

  const parsedData = createWebsiteSchema.safeParse({
    url: formData.get("url"),
    frequency: formData.get("frequency"),
    regions: formData.getAll("regions") as string[],
  });

  if (!parsedData.success) {
    console.error("Validation Errors:", parsedData.error.flatten().fieldErrors);
    return {
      errors: parsedData.error.flatten().fieldErrors,
    };
  }

  console.log("Parsed Data:", parsedData.data);

  try {
    await axios.post(`${API_URL}/website`, parsedData.data, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    });

    revalidatePath("/dashboard");
  } catch (error) {
    if (error instanceof AxiosError) {
      return {
        error: error.response?.data?.error || "An error occurred while creating the website.",
      }
    }

    return {
      error: "An unexpected error occurred while creating the website.",
    }
  }
}

export async function pingWebsite(_: unknown, url: string) {
  try {
    const res = await axios.head(url, {
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
        error: e.response?.statusText || "Failed to ping the website."
      }
    }

    return {
      status: false,
      error: "An unexpected error occurred while pinging the website."
    }
  }
}
