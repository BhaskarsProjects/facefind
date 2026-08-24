import { apiClient } from "@/shared/services/api-client";

import type {
  CreateFolderRequest,
  Folder,
  UpdateFolderRequest,
} from "../types/folder.types";

export async function getFolders(
  accessToken: string,
): Promise<Folder[]> {
  const response = await apiClient.get<Folder[]>(
    "/folders",
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data;
}

export async function createFolder(
  payload: CreateFolderRequest,
  accessToken: string,
): Promise<Folder> {
  const response = await apiClient.post<Folder>(
    "/folders",
    payload,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data;
}

export async function updateFolder(
  id: string,
  payload: UpdateFolderRequest,
  accessToken: string,
): Promise<Folder> {
  const response = await apiClient.patch<Folder>(
    `/folders/${id}`,
    payload,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data;
}