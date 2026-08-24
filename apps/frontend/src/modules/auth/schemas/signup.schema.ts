import { z } from "zod";

export const signupSchema = z
  .object({
    email: z.email("Please enter a valid email address"),

    password: z
      .string()
      .min(1, "Password is required")
      .min(15, "Password must be at least 15 characters")
      .max(72, "Password must not exceed 72 characters"),

    confirmPassword: z
      .string()
      .min(1, "Please confirm your password"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

export type SignupFormValues = z.infer<typeof signupSchema>;