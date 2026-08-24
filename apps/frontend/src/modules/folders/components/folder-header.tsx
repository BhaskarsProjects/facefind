import { useState } from "react";

import { Button } from "@/shared/components/ui/button";

import { PhotoUploadDialog } from "./photo-upload-dialog";

interface FolderHeaderProps {
  folderId: string;
}

export function FolderHeader({
  folderId,
}: FolderHeaderProps) {
  const [uploadOpen, setUploadOpen] = useState(false);

  return (
    <>
      <section className="flex items-center justify-between gap-4">
        <div>
          <h1 className="text-2xl font-semibold tracking-tight">
            Folder
          </h1>

          <p className="mt-1 text-sm text-muted-foreground">
            Folder ID: {folderId}
          </p>
        </div>

        <div className="flex items-center gap-2">
          <Button variant="outline">
            Access
          </Button>

          <Button
            type="button"
            onClick={() => setUploadOpen(true)}
          >
            Upload photos
          </Button>
        </div>
      </section>

      <PhotoUploadDialog
        folderId={folderId}
        open={uploadOpen}
        onOpenChange={setUploadOpen}
      />
    </>
  );
}