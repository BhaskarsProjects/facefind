import {
  Folder,
  Home,
  Clock3,
  Star,
  Users,
  Settings,
  Plus,
} from "lucide-react";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/shared/components/ui/sidebar";

import { Button } from "@/shared/components/ui/button";

import CreateFolderDialog from "@/modules/folders/components/create-folder-dialog";

const mainItems = [
  {
    title: "My Library",
    icon: Home,
    path: "/dashboard",
  },
  {
    title: "Recent",
    icon: Clock3,
    path: "/dashboard/recent",
  },
  {
    title: "Starred",
    icon: Star,
    path: "/dashboard/starred",
  },
];

const folders = ["Wedding", "Family", "Travel", "Work"];

export function AppSidebar() {
  const navigate = useNavigate();

  const [folderCreation, setFolderCreation] = useState(false);

  return (
    <Sidebar>
      <SidebarHeader className="p-4">
        <div className="flex items-center gap-2">
          <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-primary-foreground">
            <span className="text-sm font-bold">F</span>
          </div>

          <span className="text-lg font-semibold">FaceFind</span>
        </div>
      </SidebarHeader>

      <SidebarContent className="px-2">
        {/* Main navigation */}
        <div className="mb-6">
          <SidebarMenu>
            {mainItems.map((item) => (
              <SidebarMenuItem key={item.title}>
                <SidebarMenuButton onClick={() => navigate(item.path)}>
                  <item.icon />
                  <span>{item.title}</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </div>

        {/* Folders */}
        <div className="mb-6">
          <div className="mb-2 flex items-center justify-between px-2">
            <span className="text-xs font-medium uppercase tracking-wider text-muted-foreground">
              Folders
            </span>

            <Button
              type="button"
              variant="ghost"
              className="text-muted-foreground transition-colors hover:text-foreground"
              onClick={() => setFolderCreation(true)}
            >
              <Plus className="h-4 w-4" />
            </Button>
          </div>

          <SidebarMenu>
            {folders.map((folder) => (
              <SidebarMenuItem key={folder}>
                <SidebarMenuButton
                  onClick={() =>
                    navigate(`/dashboard/folders/${encodeURIComponent(folder)}`)
                  }
                >
                  <Folder />
                  <span>{folder}</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </div>

        {/* Shared */}
        <div>
          <div className="mb-2 px-2">
            <span className="text-xs font-medium uppercase tracking-wider text-muted-foreground">
              Shared
            </span>
          </div>

          <SidebarMenu>
            <SidebarMenuItem>
              <SidebarMenuButton onClick={() => navigate("/dashboard/shared")}>
                <Users />
                <span>Shared with me</span>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </div>
      </SidebarContent>

      {/* Settings */}
      <SidebarFooter className="px-2 pb-4">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton onClick={() => navigate("/dashboard/settings")}>
              <Settings />
              <span>Settings</span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>

      <CreateFolderDialog
        open={folderCreation}
        onOpenChange={setFolderCreation}
      />
    </Sidebar>
  );
}
