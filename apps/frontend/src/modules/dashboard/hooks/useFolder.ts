import {
  useCallback,
  useEffect,
  useState,
} from "react";

import { useAuth0 } from "@auth0/auth0-react";

import {
  createFolder,
  getFolders,
  updateFolder,
} from "../services/folder.service";

import type {
  CreateFolderRequest,
  Folder,
  UpdateFolderRequest,
} from "../types/folder.types";

export function useFolders() {
  const {
    getAccessTokenSilently,
  } = useAuth0();

  const [folders, setFolders] = useState<Folder[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const loadFolders = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);

      const accessToken =
        await getAccessTokenSilently();

      const data = await getFolders(accessToken);

      setFolders(data);
    } catch (error) {
      console.error(error);
      setError("Failed to load folders");
    } finally {
      setLoading(false);
    }
  }, [getAccessTokenSilently]);

  const addFolder = async (
    payload: CreateFolderRequest,
  ) => {
    const accessToken =
      await getAccessTokenSilently();

    const folder = await createFolder(
      payload,
      accessToken,
    );

    setFolders((current) => [
      ...current,
      folder,
    ]);

    return folder;
  };

  const renameFolder = async (
    id: string,
    payload: UpdateFolderRequest,
  ) => {
    const accessToken =
      await getAccessTokenSilently();

    const updatedFolder = await updateFolder(
      id,
      payload,
      accessToken,
    );

    setFolders((current) =>
      current.map((folder) =>
        folder.id === id
          ? updatedFolder
          : folder,
      ),
    );

    return updatedFolder;
  };

  useEffect(() => {
    const timer = setTimeout(() => {
      void loadFolders();
    }, 0);

    return () => clearTimeout(timer);
  }, [loadFolders]);

  return {
    folders,
    loading,
    error,
    loadFolders,
    addFolder,
    renameFolder,
  };
}