import { useAuth0 } from "@auth0/auth0-react";

export default function HomePage() {
  const {
    isLoading,
    isAuthenticated,
    user,
    error,
  } = useAuth0();

  if (isLoading) {
    return <div>Auth0 loading...</div>;
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
    return <div>Not authenticated</div>;
  }

  return (
    <div>
      <h1>Welcome {user?.name}</h1>

      <p>Email: {user?.email}</p>
      <p>Auth0 ID: {user?.sub}</p>

      {user?.picture && (
        <img
          src={user.picture}
          alt={user.name ?? "Profile"}
          className="h-12 w-12 rounded-full"
        />
      )}
    </div>
  );
}