import { BACKEND_URL } from "@/lib/env"
import { useMutation } from "@tanstack/react-query"
import axios from "axios"
import { useNavigate } from "react-router-dom"

interface User {
  first_name: string
  last_name: string
  email: string
  phone_number: string
  avatar: string
  password: string
}

export function useSignup() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: (user: User) => {
      return axios.post(`${BACKEND_URL}/auth/register`, user);
    },
    onSuccess: () => {
      navigate("/dashboard")
    }
  });
}
