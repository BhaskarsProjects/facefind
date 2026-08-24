import { Trash2 } from "lucide-react";

import { useFolderPhotos } from "../hooks/useFolderPhotos";

interface FolderPhotoGridProps {
  folderId: string;
}

export function FolderPhotoGrid({
  folderId,
}: FolderPhotoGridProps) {
  const {
    photos,
    loading,
    error,
    deletingPhotoId,
    deletePhoto,
  } = useFolderPhotos(folderId);

  if (loading) {
    return (
      <div className="flex min-h-48 items-center justify-center rounded-xl border border-dashed">
        <p className="text-sm text-muted-foreground">
          Loading photos...
        </p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-48 items-center justify-center rounded-xl border border-dashed">
        <p className="text-sm text-destructive">
          {error}
        </p>
      </div>
    );
  }

  if (photos.length === 0) {
    return (
      <div className="flex min-h-48 items-center justify-center rounded-xl border border-dashed">
        <p className="text-sm text-muted-foreground">
          No photos in this folder yet.
        </p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
      {photos.map((photo) => {
        const deleting = deletingPhotoId === photo.id;

        return (
          <div
            key={photo.id}
            className="group relative overflow-hidden rounded-xl border"
          >
            <img
              src={photo.url}
              alt=""
              className="aspect-square w-full object-cover"
            />

            <button
              type="button"
              onClick={() => deletePhoto(photo.id)}
              disabled={deleting}
              className="absolute right-2 top-2 rounded-md bg-black/60 p-2 text-white opacity-0 transition-opacity group-hover:opacity-100 disabled:cursor-not-allowed disabled:opacity-50"
              aria-label="Delete photo"
            >
              <Trash2 className="size-4" />
            </button>
          </div>
        );
      })}
    </div>
  );
}