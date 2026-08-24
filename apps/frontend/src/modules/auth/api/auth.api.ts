import { apiClient } from "@/shared/services/api-client";

import type {
  AuthUser,
  LoginRequest,
  SignupRequest,
} from "../types/auth.types";

export async function login(
  data: LoginRequest,
): Promise<AuthUser> {
  const response = await apiClient.post<AuthUser>(
    "/auth/login",
    data,
  );

  return response.data;
}

export async function signup(
  data: SignupRequest,
): Promise<AuthUser> {
  const response = await apiClient.post<AuthUser>(
    "/auth/signup",
    data,
  );

  return response.data;
}

export async function logout(): Promise<void> {
  await apiClient.post("/auth/logout");
}

export function startGoogleLogin(): void {
  window.location.href = `${apiClient.defaults.baseURL}/auth/google`;
}