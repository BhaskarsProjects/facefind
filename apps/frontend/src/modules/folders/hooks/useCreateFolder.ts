import { useMutation } from "@tanstack/react-query";
import { createFolder } from "../services/folder.service";
import type { CreateFolderInput } from "../types/folder.types";

import { useAuth0 } from "@auth0/auth0-react";

export function useCreateFolder() {
  const { getAccessTokenSilently } = useAuth0();

  return useMutation({
    mutationFn: async (data: CreateFolderInput) => {
      const accessToken = await getAccessTokenSilently();

      createFolder(data, accessToken);
    },
  });
}
