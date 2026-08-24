import { useCallback, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";

import { searchFolderPhotos } from "../services/photo.service";
import type { FaceSearchResult } from "../types/photo.types";

export function useFolderSearch(folderId: string) {
  const { getAccessTokenSilently } = useAuth0();

  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [results, setResults] = useState<FaceSearchResult[]>([]);
  const [isSearching, setIsSearching] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const selectFile = useCallback((file: File | null) => {
    setSelectedFile(file);
    setResults([]);
    setError(null);
  }, []);

  const search = useCallback(async () => {
    if (!selectedFile || isSearching) {
      return;
    }

    try {
      setIsSearching(true);
      setError(null);

      const accessToken = await getAccessTokenSilently();

      const searchResults = await searchFolderPhotos(
        folderId,
        selectedFile,
        accessToken,
      );

      setResults(searchResults);
    } catch (err) {
      console.error("Failed to search folder:", err);

      setError(
        "Failed to search this folder. Please try again.",
      );
    } finally {
      setIsSearching(false);
    }
  }, [
    folderId,
    selectedFile,
    isSearching,
    getAccessTokenSilently,
  ]);

  return {
    selectedFile,
    results,
    isSearching,
    error,
    selectFile,
    search,
  };
}