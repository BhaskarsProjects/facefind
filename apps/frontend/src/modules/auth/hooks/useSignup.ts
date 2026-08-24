import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

import { signupUser } from "../services/auth.service";
import type { SignupRequest } from "../types/auth.types";

export function useSignup() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: (data: SignupRequest) => signupUser(data),

    onSuccess: () => {
      toast.success("Account created successfully");

      navigate("/auth/login");
    },
  });
}