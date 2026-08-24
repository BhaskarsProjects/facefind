import { Navigate, Outlet, useLocation } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import LoadingScreen from "@/shared/components/LoadingScreen";


export default function GuestGuard() {
  const {
    isLoading,
    isAuthenticated,
    error,
  } = useAuth0();

  const location = useLocation();

  if (isLoading) {
    return <LoadingScreen />
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

  if (isAuthenticated) {
    return (
      <Navigate
        to="/dashboard"
        replace
        state={{ from: location }}
      />
    );
  }

  return <Outlet />;
}