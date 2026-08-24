import {
  login,
  logout,
  signup,
  startGoogleLogin,
} from "../api/auth.api";

import type {
  AuthUser,
  LoginRequest,
  SignupRequest,
} from "../types/auth.types";

export async function loginUser(
  data: LoginRequest,
): Promise<AuthUser> {
  return login(data);
}

export async function signupUser(
  data: SignupRequest,
): Promise<AuthUser> {
  return signup(data);
}

export async function logoutUser(): Promise<void> {
  await logout();
}

export function loginWithGoogle(): void {
  startGoogleLogin();
}