import { Search, Upload } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { UserMenu } from "@/shared/components/UserMenu";

export function DashboardHeader() {
  return (
    <header className="flex h-16 items-center justify-between border-b px-6">
      <div className="flex w-full max-w-md items-center">
        <div className="relative w-full">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />

          <Input
            placeholder="Search your photos..."
            className="pl-9"
          />
        </div>
      </div>

      <div className="ml-6 flex items-center gap-3">
        <Button>
          <Upload />
          Upload
        </Button>

        <UserMenu />
      </div>
    </header>
  );
}