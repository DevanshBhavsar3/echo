import { BACKEND_URL } from "@/lib/env"
import {useMutation, useQueryClient} from "@tanstack/react-query"
import axios, { AxiosError } from "axios"
import { useNavigate } from "react-router-dom"
import { toast } from "sonner"

interface User {
  email: string
  password: string
}

export default function useLogin() {
  const navigate = useNavigate();
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (user: User) => {
      try {
        const response = await axios.post(`${BACKEND_URL}/auth/signin`, user);

        return response.data
      } catch (e) {
        if (e instanceof AxiosError) {
          throw new Error(e.response?.data?.error)
        }

        throw new Error("An unexpected error occurred. Please try again.")
      }
    },
    onSuccess: (data) => {
      queryClient.setQueryData(["user", "data"], data);
      toast.success("Logged in successfully.")
      navigate("/dashboard")
    },
    onError: (error) => {
      toast.error(error.message)
    }
  });
}

