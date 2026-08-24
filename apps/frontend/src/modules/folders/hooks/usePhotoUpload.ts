import { useCallback, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";

import { uploadFolderPhotos } from "../services/photo.service";
import type { PhotoUploadError } from "../types/photo.types";

const MAX_FILES = 100;
const MAX_FILE_SIZE = 25 * 1024 * 1024;

const ACCEPTED_TYPES = new Set([
  "image/jpeg",
  "image/png",
  "image/webp",
]);

export function usePhotoUpload(folderId: string) {
  const { getAccessTokenSilently } = useAuth0();

  const [uploading, setUploading] = useState(false);
  const [progress, setProgress] = useState(0);
  const [errors, setErrors] = useState<PhotoUploadError[]>([]);

  const validateFiles = useCallback((files: File[]) => {
    const validationErrors: PhotoUploadError[] = [];

    if (files.length > MAX_FILES) {
      validationErrors.push({
        fileName: "Upload",
        message: `You can upload a maximum of ${MAX_FILES} photos at once.`,
      });

      return validationErrors;
    }

    for (const file of files) {
      if (!ACCEPTED_TYPES.has(file.type)) {
        validationErrors.push({
          fileName: file.name,
          message:
            "Only JPEG, PNG, and WebP images are supported.",
        });

        continue;
      }

      if (file.size > MAX_FILE_SIZE) {
        validationErrors.push({
          fileName: file.name,
          message: "Each photo must be 25 MB or smaller.",
        });
      }
    }

    return validationErrors;
  }, []);

  const upload = useCallback(
    async (files: File[]) => {
      setErrors([]);
      setProgress(0);

      const validationErrors = validateFiles(files);

      if (validationErrors.length > 0) {
        setErrors(validationErrors);
        return;
      }

      if (files.length === 0) {
        return;
      }

      try {
        setUploading(true);

        const accessToken =
          await getAccessTokenSilently();

        await uploadFolderPhotos(
          folderId,
          files,
          accessToken,
        );

        setProgress(100);
      } catch (error) {
        console.error(
          "Failed to upload photos:",
          error,
        );

        setErrors([
          {
            fileName: "Upload",
            message: "Failed to upload photos.",
          },
        ]);
      } finally {
        setUploading(false);
      }
    },
    [
      folderId,
      getAccessTokenSilently,
      validateFiles,
    ],
  );

  return {
    upload,
    uploading,
    progress,
    errors,
  };
}