import { RouterProvider } from "react-router-dom";

import Auth0Provider from "./providers/Auth0Provider";
import QueryProvider from "./providers/QueryProvider";
import router from "./router";

import { Toaster } from "@/shared/components/ui/sonner";

export default function App() {
  return (
    <QueryProvider>
      <Auth0Provider>
        <RouterProvider router={router} />
      </Auth0Provider>
      <Toaster />
    </QueryProvider>

    
  );
}