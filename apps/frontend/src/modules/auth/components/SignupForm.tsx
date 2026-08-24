import { useState } from "react";
import { Link } from "react-router-dom";
import { Eye, EyeOff } from "lucide-react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { toast } from "sonner";

import { useSignup } from "@/modules/auth/hooks/useSignup";
import {
  signupSchema,
  type SignupFormValues,
} from "@/modules/auth/schemas/signup.schema";

import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { Label } from "@/shared/components/ui/label";
import { Separator } from "@/shared/components/ui/separator";

import googleIcon from "@/assets/google.svg";

export default function SignupForm() {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const { mutateAsync: signup, isPending } = useSignup();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupFormValues>({
    resolver: zodResolver(signupSchema),
    defaultValues: {
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  async function onSubmit(values: SignupFormValues) {
    try {
      await signup(values);
    } catch {
      toast.error("Unable to create your account", {
        description: "Please check your details and try again.",
      });
    }
  }

  return (
    <div className="w-full space-y-6">
      {/* Header */}
      <div className="space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Create your account
        </h1>

        <p className="text-sm text-muted-foreground">
          Sign up to get started
        </p>
      </div>

      <div className="space-y-4">
        {/* Google */}
        <Button
          type="button"
          variant="outline"
          disabled={isPending}
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

        {/* Signup Form */}
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
            <Label htmlFor="password">Password</Label>

            <div className="relative">
              <Input
                id="password"
                type={showPassword ? "text" : "password"}
                placeholder="••••••••"
                autoComplete="new-password"
                disabled={isPending}
                className="pr-10"
                {...register("password")}
              />

              <button
                type="button"
                onClick={() =>
                  setShowPassword((visible) => !visible)
                }
                disabled={isPending}
                aria-label={
                  showPassword ? "Hide password" : "Show password"
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

          {/* Confirm Password */}
          <div className="space-y-2">
            <Label htmlFor="confirm-password">
              Confirm password
            </Label>

            <div className="relative">
              <Input
                id="confirm-password"
                type={showConfirmPassword ? "text" : "password"}
                placeholder="••••••••"
                autoComplete="new-password"
                disabled={isPending}
                className="pr-10"
                {...register("confirmPassword")}
              />

              <button
                type="button"
                onClick={() =>
                  setShowConfirmPassword((visible) => !visible)
                }
                disabled={isPending}
                aria-label={
                  showConfirmPassword
                    ? "Hide confirm password"
                    : "Show confirm password"
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
                {showConfirmPassword ? (
                  <EyeOff className="h-4 w-4" />
                ) : (
                  <Eye className="h-4 w-4" />
                )}
              </button>
            </div>

            {errors.confirmPassword && (
              <p className="text-sm text-destructive">
                {errors.confirmPassword.message}
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
              hover:bg-black
              hover:text-white
              hover:scale-[1.02]
              hover:shadow-md
              active:scale-[0.98]
            "
          >
            {isPending ? "Creating account..." : "Create account"}
          </Button>
        </form>
      </div>

      {/* Login Link */}
      <p className="text-center text-sm text-muted-foreground">
        Already have an account?{" "}
        <Link
          to="/auth/login"
          className="
            font-medium
            text-foreground
            underline
            underline-offset-4
            hover:text-foreground/80
          "
        >
          Sign in
        </Link>
      </p>
    </div>
  );
}