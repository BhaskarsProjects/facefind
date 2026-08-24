import { useState } from "react";
import { ImagePlus, X } from "lucide-react";

import { Button } from "@/shared/components/ui/button";

import { usePhotoUpload } from "../hooks/usePhotoUpload";

interface PhotoUploadDialogProps {
  folderId: string;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function PhotoUploadDialog({
  folderId,
  open,
  onOpenChange,
}: PhotoUploadDialogProps) {
  const {
    upload,
    uploading,
    progress,
    errors,
  } = usePhotoUpload(folderId);

  const [files, setFiles] = useState<File[]>([]);

  if (!open) {
    return null;
  }

  const handleFileChange = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const selectedFiles = Array.from(
      event.target.files ?? [],
    );

    setFiles((current) => {
      const remaining = 100 - current.length;

      return [
        ...current,
        ...selectedFiles.slice(0, remaining),
      ];
    });

    event.target.value = "";
  };

  const removeFile = (index: number) => {
    setFiles((current) =>
      current.filter(
        (_, fileIndex) => fileIndex !== index,
      ),
    );
  };

  const handleUpload = async () => {
    if (files.length === 0) {
      return;
    }

    await upload(files);
  };

  const handleClose = () => {
    if (uploading) {
      return;
    }

    setFiles([]);
    onOpenChange(false);
  };

  return (
    <div className="fixed inset-0 z-[100] min-h-screen w-screen overflow-y-auto bg-black/50 p-4">
      <div className="flex min-h-full items-center justify-center">
        <div className="w-full max-w-2xl rounded-xl border bg-background shadow-xl">
          {/* Header */}
          <div className="flex items-center justify-between border-b p-5">
            <div>
              <h2 className="text-lg font-semibold">
                Upload photos
              </h2>

              <p className="mt-1 text-sm text-muted-foreground">
                Add photos to this folder.
              </p>
            </div>

            <button
              type="button"
              onClick={handleClose}
              disabled={uploading}
              className="rounded-md p-2 hover:bg-accent disabled:cursor-not-allowed disabled:opacity-50"
            >
              <X className="h-5 w-5" />
            </button>
          </div>

          {/* Content */}
          <div className="space-y-5 p-5">
            {/* File picker */}
            <label
              className={[
                "flex h-32 flex-col items-center justify-center gap-2",
                "rounded-lg border border-dashed transition",
                uploading || files.length >= 100
                  ? "cursor-not-allowed opacity-50"
                  : "cursor-pointer hover:bg-accent",
              ].join(" ")}
            >
              <ImagePlus className="h-6 w-6 text-muted-foreground" />

              <span className="text-sm font-medium">
                Choose photos
              </span>

              <span className="text-center text-xs text-muted-foreground">
                JPEG, PNG, or WebP · Maximum 25 MB each ·
                Up to 100 photos
              </span>

              <input
                type="file"
                accept="image/jpeg,image/png,image/webp"
                multiple
                disabled={
                  uploading ||
                  files.length >= 100
                }
                className="hidden"
                onChange={handleFileChange}
              />
            </label>

            {/* Selected photos */}
            {files.length > 0 && (
              <div>
                <div className="mb-3 flex items-center justify-between">
                  <span className="text-sm font-medium">
                    Selected photos
                  </span>

                  <span className="text-sm text-muted-foreground">
                    {files.length} / 100
                  </span>
                </div>

                <div className="grid max-h-64 grid-cols-4 gap-3 overflow-y-auto sm:grid-cols-5">
                  {files.map((file, index) => (
                    <PhotoPreview
                      key={`${file.name}-${file.lastModified}-${index}`}
                      file={file}
                      onRemove={() => removeFile(index)}
                      disabled={uploading}
                    />
                  ))}
                </div>
              </div>
            )}

            {/* Errors */}
            {errors.length > 0 && (
              <div className="rounded-lg border border-destructive/30 bg-destructive/5 p-3">
                <p className="text-sm font-medium text-destructive">
                  Some files could not be uploaded.
                </p>

                <ul className="mt-2 space-y-1">
                  {errors.map((error, index) => (
                    <li
                      key={`${error.fileName}-${index}`}
                      className="text-xs text-muted-foreground"
                    >
                      <span className="font-medium">
                        {error.fileName}:
                      </span>{" "}
                      {error.message}
                    </li>
                  ))}
                </ul>
              </div>
            )}

            {/* Progress */}
            {uploading && (
              <div className="space-y-2">
                <div className="flex justify-between text-xs text-muted-foreground">
                  <span>Uploading photos...</span>

                  <span>{progress}%</span>
                </div>

                <div className="h-2 overflow-hidden rounded-full bg-muted">
                  <div
                    className="h-full bg-primary transition-all"
                    style={{
                      width: `${progress}%`,
                    }}
                  />
                </div>
              </div>
            )}
          </div>

          {/* Footer */}
          <div className="flex justify-end gap-2 border-t p-5">
            <Button
              type="button"
              variant="outline"
              disabled={uploading}
              onClick={handleClose}
            >
              Cancel
            </Button>

            <Button
              type="button"
              disabled={
                uploading ||
                files.length === 0
              }
              onClick={handleUpload}
            >
              {uploading
                ? "Uploading..."
                : `Upload ${files.length || ""} photos`}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

interface PhotoPreviewProps {
  file: File;
  onRemove: () => void;
  disabled: boolean;
}

function PhotoPreview({
  file,
  onRemove,
  disabled,
}: PhotoPreviewProps) {
  const [preview] = useState(() =>
    URL.createObjectURL(file),
  );

  return (
    <div className="group relative aspect-square overflow-hidden rounded-lg border bg-muted">
      <img
        src={preview}
        alt={file.name}
        className="h-full w-full object-cover"
      />

      <button
        type="button"
        onClick={onRemove}
        disabled={disabled}
        className="absolute right-1 top-1 rounded-full bg-black/60 p-1 text-white opacity-0 transition group-hover:opacity-100 disabled:cursor-not-allowed"
      >
        <X className="h-3 w-3" />
      </button>
    </div>
  );
}