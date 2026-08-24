import { useState } from "react";
import { Link } from "react-router-dom";
import { Eye, EyeOff } from "lucide-react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { toast } from "sonner";

import { useLogin } from "@/modules/auth/hooks/useLogin";
import {
  loginSchema,
  type LoginFormValues,
} from "@/modules/auth/schemas/login.schema";

import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { Label } from "@/shared/components/ui/label";
import { Separator } from "@/shared/components/ui/separator";

import googleIcon from "@/assets/google.svg";

export default function LoginForm() {
  const [showPassword, setShowPassword] = useState(false);

  const { mutateAsync: login, isPending } = useLogin();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: LoginFormValues) {
    try {
      await login(values);

      toast.success("Signed in successfully");
    } catch {
      toast.error("Unable to sign in", {
        description:
          "Please check your email and password and try again.",
      });
    }
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Welcome back
        </h1>

        <p className="text-sm text-muted-foreground">
          Sign in to your account to continue
        </p>
      </div>

      <div className="space-y-4">
        {/* Google Login */}
        <Button
          type="button"
          variant="outline"
          className="
            w-full h-10
            transition-all duration-200 ease-out
            hover:scale-[1.02]
            hover:bg-transparent
            hover:text-foreground
            hover:border-input
            hover:shadow-md
            active:scale-[0.98]
          "
        >
          <img src={googleIcon} alt="" className="h-4 w-4" />
          Continue with Google
        </Button>

        {/* Divider */}
        <div className="flex items-center gap-4">
          <Separator className="flex-1" />

          <span className="text-xs uppercase text-muted-foreground">
            Or
          </span>

          <Separator className="flex-1" />
        </div>

        {/* Login Form */}
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Email */}
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>

            <Input
              id="email"
              type="email"
              placeholder="you@example.com"
              autoComplete="email"
              disabled={isPending}
              {...register("email")}
            />

            {errors.email && (
              <p className="text-sm text-destructive">
                {errors.email.message}
              </p>
            )}
          </div>

          {/* Password */}
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label htmlFor="password">Password</Label>

              <Link
                to="/auth/forgot-password"
                className="
                  text-sm
                  text-muted-foreground
                  hover:text-foreground
                "
              >
                Forgot password?
              </Link>
            </div>

            <div className="relative">
              <Input
                id="password"
                type={showPassword ? "text" : "password"}
                placeholder="••••••••"
                autoComplete="current-password"
                disabled={isPending}
                {...register("password")}
              />

              <button
                type="button"
                onClick={() =>
                  setShowPassword((visible) => !visible)
                }
                disabled={isPending}
                aria-label={
                  showPassword
                    ? "Hide password"
                    : "Show password"
                }
                className="
                  absolute right-0 top-0
                  flex h-full w-10
                  items-center justify-center
                  text-muted-foreground
                  transition-colors
                  hover:text-foreground
                  disabled:pointer-events-none
                "
              >
                {showPassword ? (
                  <EyeOff className="h-4 w-4" />
                ) : (
                  <Eye className="h-4 w-4" />
                )}
              </button>
            </div>

            {errors.password && (
              <p className="text-sm text-destructive">
                {errors.password.message}
              </p>
            )}
          </div>

          {/* Submit */}
          <Button
            type="submit"
            disabled={isPending}
            className="
              w-full h-10
              bg-black text-white
              transition-all duration-200 ease-out
              hover:bg-black hover:text-white
              hover:scale-[1.02]
              hover:shadow-md
              active:scale-[0.98]
            "
          >
            {isPending ? "Signing in..." : "Sign in"}
          </Button>
        </form>
      </div>

      {/* Signup */}
      <p className="text-center text-sm text-muted-foreground">
        Don't have an account?{" "}
        <Link
          to="/auth/signup"
          className="
            font-medium
            text-foreground
            underline
            underline-offset-4
          "
        >
          Sign up
        </Link>
      </p>
    </div>
  );
}