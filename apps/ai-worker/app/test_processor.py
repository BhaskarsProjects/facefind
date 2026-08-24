from pathlib import Path

from app.embedder import FaceEmbedder
from app.processor import PhotoProcessor


def main():
    image_path = Path("test.jpg")

    image_bytes = image_path.read_bytes()

    embedder = FaceEmbedder()
    processor = PhotoProcessor(embedder)

    faces = processor.process(image_bytes)

    print(f"Detected {len(faces)} face(s)")

    for index, face in enumerate(faces, start=1):
        print(f"\nFace {index}")
        print(f"Bounding box: {face.x1}, {face.y1}, {face.x2}, {face.y2}")
        print(f"Confidence: {face.confidence}")
        print(f"Landmarks: {face.landmarks}")
        print(f"Embedding shape: {face.embedding.shape}")
        print(f"Embedding dtype: {face.embedding.dtype}")


if __name__ == "__main__":
    main()