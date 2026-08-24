import { useAuth0 } from "@auth0/auth0-react";

export function useAuth() {
  const {
    user,
    isAuthenticated,
    isLoading,
    error,
  } = useAuth0();

  return {
    user: user ?? null,
    isAuthenticated,
    isLoading,
    error,
  };
}