import useUser from "@/api/query/user.ts";

export default function DashboardPage() {
  const { user, logout } = useUser()

  return <div>
    {JSON.stringify(user)}
    <button onClick={logout}>
      Log out
    </button>
  </div>
}
