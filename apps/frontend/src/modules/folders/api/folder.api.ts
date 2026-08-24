import type { CreateFolderInput } from "../types/folder.types";
import { apiClient } from "@/shared/services/api-client";

export async function createFolderApi(
  data: CreateFolderInput,
  accessToken: string,
) {
  const response = await apiClient.post("/folders", data, {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });

  return response.data;
}