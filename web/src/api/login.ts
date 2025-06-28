import { BACKEND_URL } from "@/lib/env"
import { useMutation } from "@tanstack/react-query"
import axios from "axios"
import { useNavigate } from "react-router-dom"

interface User {
  email: string
  password: string
}

export function useLogin() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: (user: User) => {
      return axios.post(`${BACKEND_URL}/auth/signin`, user);
    },
    onSuccess: () => {
      navigate("/dashboard")
    }
  });
}

