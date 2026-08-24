import { X } from "lucide-react";
import { useState } from "react";

import { useCreateFolder } from "../hooks/useCreateFolder";

interface CreateFolderDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export default function CreateFolderDialog({
  open,
  onOpenChange,
}: CreateFolderDialogProps) {
  const [name, setName] = useState("");
  const { mutate, isPending } = useCreateFolder();

  if (!open) {
    return null;
  }

  const handleClose = () => {
    setName("");
    onOpenChange(false);
  };

  const handleCreate = () => {
    const trimmedName = name.trim();

    if (!trimmedName) return;

    mutate(
      { name: trimmedName },
      {
        onSuccess: () => {
          handleClose();
        },
      },
    );
  };

  return (
    <div className="fixed inset-0 z-[100] flex min-h-screen w-screen items-center justify-center bg-black/50 p-4">
      <div className="w-full max-w-md rounded-lg bg-background shadow-lg">
        <div className="flex items-center justify-between border-b p-5">
          <div>
            <h2 className="text-lg font-semibold">Create folder</h2>
            <p className="mt-1 text-sm text-muted-foreground">
              Enter a name for your new folder.
            </p>
          </div>

          <button
            type="button"
            onClick={handleClose}
            className="rounded-md p-2 hover:bg-accent"
          >
            <X className="h-5 w-5" />
          </button>
        </div>

        <div className="p-5">
          <label htmlFor="folder-name" className="text-sm font-medium">
            Folder name
          </label>

          <input
            id="folder-name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="e.g. Wedding"
            className="mt-2 h-10 w-full rounded-md border bg-background px-3 text-sm outline-none focus:ring-2 focus:ring-ring"
            autoFocus
          />

          <div className="mt-5 flex justify-end gap-2">
            <button
              type="button"
              onClick={handleClose}
              className="h-10 rounded-md border px-4 text-sm font-medium hover:bg-accent"
            >
              Cancel
            </button>

            <button
              type="button"
              onClick={handleCreate}
              disabled={!name.trim() || isPending}
              className="h-10 rounded-md bg-primary px-4 text-sm font-medium text-primary-foreground hover:bg-primary/90 disabled:cursor-not-allowed disabled:opacity-50"
            >
              {isPending ? "Creating..." : "Create"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
