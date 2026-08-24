import { uploadPhotos as uploadPhotosApi } from "../api/photo.api";
import { getFolderPhotos as getFolderPhotosApi } from "../api/photo.api";
import { deletePhoto as deletePhotoApi } from "../api/photo.api";
import { searchFolder as searchFolderApi } from "../api/photo.api";

import type {
  UploadPhotoResponse,
  GetFolderPhotosResponse,
  FaceSearchResponse,
} from "../types/photo.types";

export async function uploadFolderPhotos(
  folderId: string,
  files: File[],
  accessToken: string,
): Promise<UploadPhotoResponse> {
  const formData = new FormData();

  for (const file of files) {
    formData.append("photos", file);
  }

  return uploadPhotosApi(
    folderId,
    formData,
    accessToken,
  );
}

export async function getFolderPhotos(
  folderId: string,
  accessToken: string,
): Promise<GetFolderPhotosResponse> {
  return getFolderPhotosApi(
    folderId,
    accessToken,
  );
}

export async function deleteFolderPhoto(
  folderId: string,
  photoId: string,
  accessToken: string,
): Promise<void> {
  return deletePhotoApi(
    folderId,
    photoId,
    accessToken,
  );
}

export async function searchFolderPhotos(
  folderId: string,
  file: File,
  accessToken: string,
): Promise<FaceSearchResponse["results"]> {
  return searchFolderApi(
    folderId,
    file,
    accessToken,
  );
}