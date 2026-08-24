import { apiClient } from "@/shared/services/api-client";

import type {
  CreateFolderRequest,
  Folder,
  UpdateFolderRequest,
} from "../types/folder.types";

export async function getFolders(): Promise<Folder[]> {
  const response = await apiClient.get<Folder[]>("/folders");

  return response.data;
}

export async function createFolder(
  payload: CreateFolderRequest,
): Promise<Folder> {
  const response = await apiClient.post<Folder>(
    "/folders",
    payload,
  );

  return response.data;
}

export async function updateFolder(
  id: string,
  payload: UpdateFolderRequest,
): Promise<Folder> {
  const response = await apiClient.patch<Folder>(
    `/folders/${id}`,
    payload,
  );

  return response.data;
}