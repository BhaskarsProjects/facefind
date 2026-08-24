import { Navigate, Outlet } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";

export default function ProtectedGuard() {
  const {
    isLoading,
    isAuthenticated,
    error,
  } = useAuth0();

  if (isLoading) {
    return <div>Checking authentication...</div>;
  }

  if (error) {
    return (
      <pre>
        {JSON.stringify(
          {
            message: error.message,
            name: error.name,
          },
          null,
          2
        )}
      </pre>
    );
  }

  if (!isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
}