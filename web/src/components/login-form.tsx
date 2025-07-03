import useLogin from "@/api/mutations/login";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm, useStore } from "@tanstack/react-form";
import { NavLink } from "react-router-dom";
import { z } from "zod";
import GoogleIcon from "./icons/Google";
import GithubIcon from "./icons/Github";
import { Loader2Icon } from "lucide-react";

export default function LoginForm() {
  const { mutate, isPending } = useLogin();

  const form = useForm({
    defaultValues: {
      email: "",
      password: "",
    },
    validators: {
      onSubmit: z.object({
        email: z.string().email("Please provide valid email."),
        password: z
          .string()
          .min(3, "Password can't be shorter than 3 letters.")
          .max(72, "Password can't be bigger than 72 letters."),
      }),
    },
    onSubmit: ({ value }) => {
      mutate(value);
    },
  });

  const formErrorMap = useStore(form.store, (state) => state.errorMap);

  return (
    <div className="flex flex-col gap-6">
      <Card className="overflow-hidden p-0">
        <CardContent className="grid p-0 md:grid-cols-2">
          <form
            className="p-6 md:p-8"
            onSubmit={(e) => {
              e.preventDefault();
              e.stopPropagation();
              form.handleSubmit();
            }}
          >
            <div className="flex flex-col gap-6">
              <div className="flex flex-col items-center text-center">
                <h1 className="text-2xl font-bold">Welcome back</h1>
                <p className="text-muted-foreground text-balance">
                  Login to your Echo account
                </p>
              </div>

              <form.Field name="email">
                {(field) => (
                  <div className="grid gap-3">
                    <Label htmlFor="email">Email</Label>
                    <Input
                      disabled={isPending}
                      id="email"
                      type="email"
                      value={field.state.value}
                      placeholder="me@example.com"
                      onChange={(e) => field.handleChange(e.target.value)}
                      required
                    />
                  </div>
                )}
              </form.Field>

              <form.Field name="password">
                {(field) => (
                  <div className="grid gap-3">
                    <Label htmlFor="password">Password</Label>
                    <Input
                      disabled={isPending}
                      id="password"
                      type="password"
                      value={field.state.value}
                      onChange={(e) => field.handleChange(e.target.value)}
                      required
                    />
                  </div>
                )}
              </form.Field>

              {formErrorMap.onSubmit && (
                <span className="text-muted-foreground text-sm">
                  {Object.values(formErrorMap.onSubmit)[0][0].message}
                </span>
              )}

              {isPending ? (
                <Button className="w-full" disabled>
                  <Loader2Icon className="animate-spin" />
                  Please wait
                </Button>
              ) : (
                <Button type="submit" className="w-full">
                  Login
                </Button>
              )}

              <div className="after:border-border relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t">
                <span className="bg-card text-muted-foreground relative z-10 px-2">
                  Or continue with
                </span>
              </div>
              <div className="grid gap-4 md:grid-cols-2">
                <Button variant="outline" type="button" className="w-full">
                  <GoogleIcon />
                  Google
                </Button>
                <Button variant="outline" type="button" className="w-full">
                  <GithubIcon />
                  Github
                </Button>
              </div>
              <div className="text-center text-sm">
                Don&apos;t have an account?{" "}
                <NavLink
                  to={"/signup"}
                  className="underline underline-offset-4"
                >
                  Sign up
                </NavLink>
              </div>
            </div>
          </form>
          <div className="bg-muted relative hidden md:block">
            <img
              src="/Blurred Data Center View.jpeg"
              alt="Image"
              className="absolute inset-0 h-full w-full scale-x-[-1] object-cover grayscale dark:brightness-[0.8]"
            />
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
