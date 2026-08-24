// import SignupForm from "../components/SignupForm"

// export default function SignupPage() {
//   return <SignupForm />
// }

import { useAuth0 } from "@auth0/auth0-react";

export default function SignupPage() {
  const { loginWithRedirect } = useAuth0();

  return (
    <button
      type="button"
      onClick={() =>
        loginWithRedirect({
          authorizationParams: {
            screen_hint: "signup",
          },
        })
      }
    >
      Create account
    </button>
  );
}
