import { NavLink } from "react-router-dom";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { useForm, useStore } from "@tanstack/react-form";
import { useSignup } from "@/api/signup";
import { z } from "zod";
import type { HTMLInputTypeAttribute } from "react";
import { Loader2Icon } from "lucide-react";
import type { AxiosError } from "axios";
import GoogleIcon from "./icons/Google";
import GithubIcon from "./icons/Github";

type Field = {
  name: "first_name" | "last_name" | "email" | "phone_number" | "password";
  type: HTMLInputTypeAttribute;
  placeholder: string;
  display: string;
};

export default function SignupForm() {
  const { mutate, isPending, isError, error
  } = useSignup()

  const form = useForm({
    defaultValues: {
      first_name: "",
      last_name: "",
      email: "",
      password: "",
      phone_number: "",
    },
    validators: {
      onChange: z.object({
        first_name: z
          .string()
          .min(2, "First name must have atleast 2 letters.")
          .max(10, "First name should have 10 letters at max."),
        last_name: z
          .string()
          .min(2, "Last name must have atleast 2 letters.")
          .max(10, "First name should have 10 letters at max."),
        email: z.string().email("Please provide valid email."),
        phone_number: z
          .string()
          .length(10, "Please provide valid phone number."),
        password: z
          .string()
          .min(3, "Password must have atleast 3 letters.")
          .max(72, "Password should have 72 letters at max."),
      }),
    },
    onSubmit: ({ value }) => {
      const user = {
        ...value,
        avatar:
          "https://vercel.com/api/www/avatar?s=44&teamId=team_uwb1qNu0MOIaDS7cMTDVUn9M",
      };

      mutate(user);
    },
  });

  const formErrorMap = useStore(form.store, (state) => state.errorMap);

  const fields: Field[] = [
    {
      name: "first_name",
      type: "text",
      placeholder: "John",
      display: "First Name",
    },
    {
      name: "last_name",
      type: "text",
      placeholder: "Doe",
      display: "Last Name",
    },
    {
      name: "email",
      type: "email",
      placeholder: "me@example.com",
      display: "Email",
    },
    {
      name: "phone_number",
      type: "tel",
      placeholder: "",
      display: "Phone Number",
    },
    {
      name: "password",
      type: "password",
      placeholder: "",
      display: "Password",
    },
  ];

  return (
    <form
      className="flex flex-col gap-6"
      onSubmit={(e) => {
        e.preventDefault();
        e.stopPropagation();
        form.handleSubmit();
      }}>
      <div className="flex flex-col items-center gap-2 text-center">
        <h1 className="text-2xl font-bold">Welcome to Echo</h1>
        <p className="text-muted-foreground text-sm text-balance">
          Create your new Echo account
        </p>
      </div>

      <div className="grid gap-6">
        {fields.map((f) => (
          <form.Field name={f.name} key={f.name}>
            {(field) => (
              <div className="grid gap-3">
                <Label htmlFor={field.name}>{f.display}</Label>
                <Input
                  disabled={isPending}
                  id={field.name}
                  type={f.type}
                  value={field.state.value}
                  placeholder={f.placeholder}
                  onChange={(e) => field.handleChange(e.target.value)}
                  required
                />
              </div>
            )}
          </form.Field>
        ))}

        {formErrorMap.onChange ? (
          <span className="text-muted-foreground text-sm">
            {Object.values(formErrorMap.onChange)[0][0].message}
          </span>
        ) : (
          isError && (
            <span className="text-sm text-red-400">
              {((error as AxiosError).response?.data as { error: string }).error || "An unexpected error occurred. Please try again."}
            </span>
          )
        )}

        {isPending ? (
          <Button className="w-full" disabled>
            <Loader2Icon className="animate-spin" />
            Please wait
          </Button>
        ) : (
          <Button type="submit" className="w-full">
            Sign Up
          </Button>
        )}
      </div>

      <div className="after:border-border relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t">
        <span className="bg-card text-muted-foreground relative z-10 px-2">
          Or continue with
        </span>
      </div>
      <div className="grid gap-4 lg:grid-cols-2">
        <Button variant="outline" className="w-full">
          <GoogleIcon />
          Google
        </Button>
        <Button variant="outline" className="w-full">
          <GithubIcon />
          Github
        </Button>
      </div>
      <div className="text-center text-sm">
        Already have an account?{" "}
        <NavLink to="/login" className="underline underline-offset-4">
          Log in
        </NavLink>
      </div>
    </form>
  );
}
