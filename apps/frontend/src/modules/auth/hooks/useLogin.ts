import {
  useMutation,
  useQueryClient,
} from "@tanstack/react-query";

import { loginUser } from "../services/auth.service";
import type { LoginRequest } from "../types/auth.types";

export function useLogin() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: LoginRequest) => loginUser(data),

    onSuccess: async () => {
      await queryClient.invalidateQueries({
        queryKey: ["auth", "me"],
      });
    },
  });
}