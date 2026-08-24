import { useCallback, useEffect, useState } from "react";

import { useAuth0 } from "@auth0/auth0-react";

import {
  deleteFolderPhoto,
  getFolderPhotos,
} from "../services/photo.service";

import type { Photo } from "../types/photo.types";

export function useFolderPhotos(folderId: string) {
  const { getAccessTokenSilently } = useAuth0();

  const [photos, setPhotos] = useState<Photo[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [deletingPhotoId, setDeletingPhotoId] = useState<string | null>(
    null,
  );

  useEffect(() => {
    if (!folderId) {
      return;
    }

    let cancelled = false;

    async function loadPhotos() {
      try {
        setLoading(true);
        setError(null);

        const accessToken = await getAccessTokenSilently();

        const response = await getFolderPhotos(
          folderId,
          accessToken,
        );

        if (!cancelled) {
          setPhotos(response.photos);
        }
      } catch (error) {
        if (!cancelled) {
          console.error(
            "Failed to fetch folder photos:",
            error,
          );

          setError("Failed to load photos.");
        }
      } finally {
        if (!cancelled) {
          setLoading(false);
        }
      }
    }

    loadPhotos();

    return () => {
      cancelled = true;
    };
  }, [folderId, getAccessTokenSilently]);

  const deletePhoto = useCallback(
    async (photoId: string) => {
      try {
        setDeletingPhotoId(photoId);

        const accessToken = await getAccessTokenSilently();

        await deleteFolderPhoto(
          folderId,
          photoId,
          accessToken,
        );

        setPhotos((currentPhotos) =>
          currentPhotos.filter(
            (photo) => photo.id !== photoId,
          ),
        );
      } catch (error) {
        console.error(
          "Failed to delete photo:",
          error,
        );

        throw error;
      } finally {
        setDeletingPhotoId(null);
      }
    },
    [folderId, getAccessTokenSilently],
  );

  return {
    photos,
    loading,
    error,
    deletingPhotoId,
    deletePhoto,
  };
}