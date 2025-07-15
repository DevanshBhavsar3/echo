import axios, { AxiosError } from "axios";
import { FormState, RegisterFormSchema } from "../lib/definitions";
import { API_URL } from "../lib/constant";
import { toast } from "sonner";

export async function register(state: FormState, formData: FormData) {
  const parsedData = RegisterFormSchema.safeParse({
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
    const repsonse = await axios.post(`${API_URL}/auth/register`, {
      ...parsedData.data,
      avatar: "https://api.dicebear.com/6.x/initials/svg?seed=" + parsedData.data.name,
    })

    toast.success("Registration successful!")
  } catch (error) {
    if (error instanceof AxiosError) {
      toast.error(error.response?.data?.error)
      return
    }

    toast.error("Registration failed.")
  }
}

export async function login(state: FormState, formData: FormData) {
  const data = Object.fromEntries(formData.entries())

  try {
    const repsonse = await axios.post(`${API_URL}/auth/login`, {
      email: data.email,
      password: data.password,
    })

    toast.success("Login successful!")
  } catch (error) {
    if (error instanceof AxiosError) {
      toast.error(error.response?.data?.error)
      return
    }

    toast.error("Login failed.")
  }
}
