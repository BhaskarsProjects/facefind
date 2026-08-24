import { useRef } from "react";
import { ImagePlus, Search } from "lucide-react";

import { Button } from "@/shared/components/ui/button";

import { useFolderSearch } from "../hooks/useFolderSearch";

interface FolderSearchProps {
  folderId: string;
}

export function FolderSearch({
  folderId,
}: FolderSearchProps) {
  const inputRef = useRef<HTMLInputElement>(null);

  const {
    selectedFile,
    results,
    isSearching,
    error,
    selectFile,
    search,
  } = useFolderSearch(folderId);

  const handleFileChange = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const file = event.target.files?.[0];

    if (!file) {
      return;
    }

    selectFile(file);
  };

  /*
   * The backend searches faces, so multiple faces can belong
   * to the same photo.
   *
   * Only display each photo once.
   */
  const uniqueResults = Array.from(
    new Map(
      results.map((result) => [
        result.photo_id,
        result,
      ]),
    ).values(),
  );

  return (
    <section>
      {/* Search */}
      <div className="max-w-2xl">
        <input
          ref={inputRef}
          type="file"
          accept="image/*"
          className="hidden"
          onChange={handleFileChange}
        />

        <button
          type="button"
          onClick={() => inputRef.current?.click()}
          className="flex h-32 w-full flex-col items-center justify-center gap-2 rounded-xl border border-dashed bg-background transition hover:bg-accent"
        >
          <ImagePlus className="h-6 w-6 text-muted-foreground" />

          <span className="text-sm font-medium">
            {selectedFile
              ? selectedFile.name
              : "Upload a photo to find yourself"}
          </span>

          <span className="text-xs text-muted-foreground">
            Choose a clear photo containing the person
            you're looking for
          </span>
        </button>

        {selectedFile && (
          <div className="mt-3 flex justify-end">
            <Button
              type="button"
              onClick={search}
              disabled={isSearching}
            >
              <Search />

              {isSearching
                ? "Searching..."
                : "Find photos"}
            </Button>
          </div>
        )}

        {error && (
          <p className="mt-3 text-sm text-destructive">
            {error}
          </p>
        )}

        <p className="mt-2 text-xs text-muted-foreground">
          Upload a photo and we'll find matching faces
          in this folder.
        </p>
      </div>

      {/* Results */}
      {uniqueResults.length > 0 && (
        <div className="mt-8">
          <div className="mb-4">
            <h2 className="text-lg font-semibold">
              Matching photos
            </h2>

            <p className="text-sm text-muted-foreground">
              Found {uniqueResults.length} matching photos.
            </p>
          </div>

          <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-4">
            {uniqueResults.map((result) => (
              <div
                key={result.photo_id}
                className="overflow-hidden rounded-xl border bg-background"
              >
                <img
                  src={result.url}
                  alt="Matching photo"
                  className="aspect-square w-full object-cover"
                  loading="lazy"
                />
              </div>
            ))}
          </div>
        </div>
      )}

      {/* No results */}
      {!isSearching &&
        selectedFile &&
        uniqueResults.length === 0 &&
        !error && (
          <div className="mt-8 rounded-xl border p-8 text-center">
            <p className="text-sm font-medium">
              No matching photos found
            </p>

            <p className="mt-1 text-xs text-muted-foreground">
              Try uploading a clearer photo of the person.
            </p>
          </div>
        )}
    </section>
  );
}