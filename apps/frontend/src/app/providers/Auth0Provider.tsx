import { Auth0Provider } from "@auth0/auth0-react";
import { env } from "@/app/config/env";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <Auth0Provider
      domain={env.auth0Domain}
      clientId={env.auth0ClientId}
      cacheLocation="localstorage"
      authorizationParams={{
        redirect_uri: window.location.origin,
        audience: env.auth0Audience,
      }}
    >
      {children}
    </Auth0Provider>
  );
}
