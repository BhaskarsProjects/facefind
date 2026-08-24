const apiUrl = import.meta.env.VITE_API_URL;
const auth0Domain = import.meta.env.VITE_AUTH0_DOMAIN;
const auth0ClientId = import.meta.env.VITE_AUTH0_CLIENT_ID;
const auth0Audience = import.meta.env.VITE_AUTH0_AUDIENCE;

if (!apiUrl) {
  throw new Error("VITE_API_URL is not defined");
}

if (!auth0Domain) {
  throw new Error("VITE_AUTH0_DOMAIN is not defined");
}

if (!auth0ClientId) {
  throw new Error("VITE_AUTH0_CLIENT_ID is not defined");
}

if (!auth0Audience) {
  throw new Error("VITE_AUTH0_AUDIENCE is not defined");
}

export const env = {
  apiUrl,
  auth0Domain,
  auth0ClientId,
  auth0Audience,
} as const;