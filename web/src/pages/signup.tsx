import SignupForm from "@/components/signup-form";
import { NavLink } from "react-router-dom";

export default function SignupPage() {
  return <div className="grid min-h-svh lg:grid-cols-2">
    <div className="flex flex-col gap-4 p-6 md:p-10">
      <NavLink to={"/"} className="font-medium hover:underline w-fit">
        Echo
      </NavLink>
      <div className="flex flex-1 items-center justify-center">
        <div className="w-full max-w-sm">
          <SignupForm />
        </div>
      </div>
    </div>
    <div className="bg-muted relative hidden lg:block">
      <img
        src="/Server Rack Close-Up.jpeg"
        alt="Server Image"
        className="absolute inset-0 h-full w-full object-cover dark:brightness-70"
      />
    </div>
  </div >
}
