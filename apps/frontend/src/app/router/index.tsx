import { createBrowserRouter } from "react-router-dom";

import GuestGuard from "./guards/GuestGuard";
import ProtectedGuard from "./guards/ProtectedGuard";

import LandingPage from "@/modules/landing/pages/LandingPage";
import DashboardLayout from "@/app/router/layouts/DashboardLayout";
import DashboardPage from "@/modules/dashboard/pages/DashboardPage";
import FolderPage from "@/modules/folders/pages/FolderPage";

const router = createBrowserRouter([
  // Guest-only
  {
    element: <GuestGuard />,
    children: [
      {
        path: "/",
        element: <LandingPage />,
      },
    ],
  },

  // Protected
  {
    element: <ProtectedGuard />,
    children: [
      {
        element: <DashboardLayout />,
        children: [
          {
            path: "/dashboard",
            element: <DashboardPage />,
          },
          {
            path: "/folders/:folderId",
            element: <FolderPage />,
          }
        ],
      },
    ],
  },
]);

export default router;