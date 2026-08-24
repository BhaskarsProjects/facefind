import { useFolders } from "@/modules/dashboard/hooks/useFolder";

import { FolderCard } from "../components/folder-card";
import { PhotoCard } from "../components/photo-card";

const recentPhotos = [
  "https://images.unsplash.com/photo-1519741497674-611481863552",
  "https://images.unsplash.com/photo-1511285560929-80b456fea0bc",
  "https://images.unsplash.com/photo-1507504031003-b417219a0fde",
  "https://images.unsplash.com/photo-1490750967868-88aa4486c946",
  "https://images.unsplash.com/photo-1520854221256-17451cc331bf",
  "https://images.unsplash.com/photo-1504150558240-0b4fd8946624",
];

export default function DashboardPage() {
  const {
    folders,
    loading,
    error,
  } = useFolders();

  return (
    <div className="mx-auto max-w-7xl px-6 py-8">
      {/* Header */}
      <section>
        <div>
          <h1 className="text-2xl font-semibold tracking-tight">
            My Library
          </h1>

          <p className="mt-1 text-sm text-muted-foreground">
            Organize and find your photos.
          </p>
        </div>
      </section>

      {/* Folders */}
      <section className="mt-10">
        <div className="mb-4 flex items-center justify-between">
          <h2 className="text-lg font-semibold">
            Folders
          </h2>

          <button
            type="button"
            className="text-sm font-medium text-muted-foreground hover:text-foreground"
          >
            View all
          </button>
        </div>

        {/* Loading */}
        {loading && (
          <p className="text-sm text-muted-foreground">
            Loading folders...
          </p>
        )}

        {/* Error */}
        {!loading && error && (
          <p className="text-sm text-destructive">
            {error}
          </p>
        )}

        {/* Empty */}
        {!loading && !error && folders.length === 0 && (
          <p className="text-sm text-muted-foreground">
            No folders yet.
          </p>
        )}

        {/* Folders */}
        {!loading && !error && folders.length > 0 && (
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
            {folders.map((folder) => (
              <FolderCard
                key={folder.id}
                id={folder.id}
                name={folder.name}
              />
            ))}
          </div>
        )}
      </section>

      {/* Recent photos */}
      <section className="mt-10">
        <div className="mb-4 flex items-center justify-between">
          <h2 className="text-lg font-semibold">
            Recent photos
          </h2>

          <button
            type="button"
            className="text-sm font-medium text-muted-foreground hover:text-foreground"
          >
            View all
          </button>
        </div>

        <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-6">
          {recentPhotos.map((photo) => (
            <PhotoCard
              key={photo}
              src={`${photo}?auto=format&fit=crop&w=500&q=80`}
            />
          ))}
        </div>
      </section>
    </div>
  );
}