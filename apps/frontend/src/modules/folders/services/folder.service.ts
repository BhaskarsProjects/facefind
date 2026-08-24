import { createFolderApi } from "../api/folder.api";
import type { CreateFolderInput } from "../types/folder.types";

export async function createFolder(
  data: CreateFolderInput,
  accessToken: string,
) {
  return createFolderApi(data, accessToken);
}
