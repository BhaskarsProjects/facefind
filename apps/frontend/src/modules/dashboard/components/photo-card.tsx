interface PhotoCardProps {
  src: string;
  alt?: string;
}

export function PhotoCard({
  src,
  alt = "Photo",
}: PhotoCardProps) {
  return (
    <button
      type="button"
      className="group overflow-hidden rounded-xl border bg-card"
    >
      <div className="aspect-square overflow-hidden">
        <img
          src={src}
          alt={alt}
          className="h-full w-full object-cover transition duration-300 group-hover:scale-105"
        />
      </div>
    </button>
  );
}