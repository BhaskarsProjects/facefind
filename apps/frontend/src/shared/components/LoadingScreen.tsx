import { Spinner } from "@/shared/components/ui/spinner";

export default function LoadingScreen() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="flex items-center gap-3 text-sm text-muted-foreground">
        <Spinner />
        <span>Loading...</span>
      </div>
    </div>
  );
}