import { BACKEND_URL } from "@/lib/env"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import axios, { AxiosError } from "axios"
import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

interface User {
  id: string
  name: string
  email: string
  avatar: string
  password: string
  created_at: string
  updated_at: string
}

export default function useUser() {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const { data: user } = useQuery({
    queryKey: ["user", "data"],
    queryFn: async () => {
      try {
        const response = await axios.get(`${BACKEND_URL}/auth/user`, { withCredentials: true })

        return response.data
      } catch (e) {
        if (e instanceof AxiosError) {
          throw new Error(e.response?.data?.error)
        }

        throw new Error("An unexpected error occurred. Please try again.")
      }
    },
    refetchOnMount: false,
    refetchOnWindowFocus: false,
    refetchOnReconnect: false,
    initialData: getLocalUser()
  })


  useEffect(() => {
    if (!user) {
      removeLocalUser();
    } else {
      saveLocalUser(user);
    }
  }, [user]);


  function logout() {
    queryClient.setQueryData(["user", "data"], null);
    navigate('/login');
  }

  return {
    user: user || null,
    logout
  }
}


function saveLocalUser(user: User): void {
  localStorage.setItem("USER", JSON.stringify(user));
}

function getLocalUser(): User | undefined {
  try {
    const user = localStorage.getItem("USER");

    if (user) {
      return JSON.parse(user)
    }

    return undefined
  } catch (e) {
    return undefined
  }
}

function removeLocalUser(): void {
  localStorage.removeItem("USER");
}
