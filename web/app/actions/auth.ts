"use server"

import axios, { AxiosError } from "axios";
import { loginSchema, registerSchema } from "@/lib/types";
import { AuthError } from "next-auth";
import { API_URL } from "../constants";
import { redirect } from "next/navigation";
import { signIn } from "../auth";

export async function register(_: unknown, formData: FormData) {
  const parsedData = registerSchema.safeParse({
    name: formData.get("name"),
    email: formData.get("email"),
    password: formData.get("password"),
  })

  if (!parsedData.success) {
    return {
      data: Object.fromEntries(formData.entries()),
      errors: parsedData.error.flatten().fieldErrors,
    }
  }

  try {
    await axios.post(`${API_URL}/auth/register`, {
      ...parsedData.data,
      avatar: "https://api.dicebear.com/6.x/initials/svg?seed=" + parsedData.data.name,
    })
  } catch (error) {
    if (error instanceof AxiosError) {
      return {
        error: error.response?.data?.error || "An error occurred during registration.",
      }
    }

    return {
      error: "An unexpected error occurred.",
    }
  }


  redirect("/login")
}

export async function login(_: unknown, formData: FormData) {
  const values = Object.fromEntries(formData.entries());

  const parsedData = loginSchema.safeParse({
    email: values["email"],
    password: values["password"],
  })

  if (!parsedData.success) {
    return { error: parsedData.error.issues[0].message }
  }

  const { email, password } = parsedData.data;

  try {
    await signIn("credentials", {
      email,
      password,
      redirect: false,
    });
  } catch (error) {
    if (error instanceof AuthError) {
      return { error: error.cause?.err?.message }
    }

    return { error: "An unexpected error occurred during login." }
  }


  redirect("/")
}
