import { useParams } from "react-router-dom";

import { FolderHeader } from "../components/folder-header";
import { FolderSearch } from "../components/folder-search";
import { FolderPhotoGrid } from "../components/folder-photo-grid";

export default function FolderPage() {
  const { folderId } = useParams<{ folderId: string }>();

  if (!folderId) {
    return (
      <div className="p-6">
        <p className="text-sm text-destructive">
          Folder not found.
        </p>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-7xl px-6 py-8">
      <div className="space-y-8">
        <FolderHeader folderId={folderId} />

        <FolderSearch folderId={folderId} />

        <section>
          <div className="mb-4">
            <h2 className="text-lg font-semibold">
              Photos
            </h2>

            <p className="text-sm text-muted-foreground">
              Photos in this folder.
            </p>
          </div>

          <FolderPhotoGrid folderId={folderId} />
        </section>
      </div>
    </div>
  );
}