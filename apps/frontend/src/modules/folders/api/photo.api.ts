import { apiClient } from "@/shared/services/api-client";

import type {
  GetFolderPhotosResponse,
  UploadPhotoResponse,
  FaceSearchResponse,
} from "../types/photo.types";

export async function uploadPhotos(
  folderId: string,
  formData: FormData,
  accessToken: string,
): Promise<UploadPhotoResponse> {
  const response = await apiClient.post<UploadPhotoResponse>(
    `/folders/${folderId}/photos`,
    formData,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "multipart/form-data",
      },
    },
  );

  return response.data;
}

export async function getFolderPhotos(
  folderId: string,
  accessToken: string,
): Promise<GetFolderPhotosResponse> {
  const response = await apiClient.get<GetFolderPhotosResponse>(
    `/${folderId}/photos`,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data;
}

export async function deletePhoto(
  folderId: string,
  photoId: string,
  accessToken: string,
): Promise<void> {
  await apiClient.delete(
    `/${folderId}/photos/${photoId}`,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
}

export async function searchFolder(
  folderId: string,
  file: File,
  accessToken: string,
): Promise<FaceSearchResponse["results"]> {
  const formData = new FormData();

  formData.append("image", file);

  const response = await apiClient.post<FaceSearchResponse>(
    `/folders/${folderId}/faces/search`,
    formData,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "multipart/form-data",
      },
    },
  );

  return response.data.results;
}