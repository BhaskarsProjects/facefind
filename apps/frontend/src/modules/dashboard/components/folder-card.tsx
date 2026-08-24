import { Folder } from "lucide-react";
import { useNavigate } from "react-router-dom";


interface FolderCardProps {
  id: string;
  name: string;
}

export function FolderCard({
  id,
  name,
}: FolderCardProps) {

  const navigate = useNavigate();

  return (
    <button
      type="button"
      onClick={() => navigate(`/folders/${id}`)}
      className="group rounded-xl border bg-card p-4 text-left transition hover:bg-accent"
    >
      <div className="flex items-start justify-between">
        <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-muted">
          <Folder className="h-5 w-5 text-muted-foreground" />
        </div>
      </div>

      <div className="mt-4">
        <p className="font-medium group-hover:underline">
          {name}
        </p>

      </div>
    </button>
  );
}