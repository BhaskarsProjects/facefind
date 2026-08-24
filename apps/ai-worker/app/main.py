import os

from dotenv import load_dotenv

from app.embedder import FaceEmbedder
from app.processor import PhotoProcessor
from app.storage import R2Storage
from app.backend_client import BackendClient


def main():
    load_dotenv()

    storage = R2Storage(
        account_id=os.environ["R2_ACCOUNT_ID"],
        access_key_id=os.environ["R2_ACCESS_KEY_ID"],
        secret_access_key=os.environ["R2_SECRET_ACCESS_KEY"],
        bucket=os.environ["R2_BUCKET"],
    )

    embedder = FaceEmbedder()

    processor = PhotoProcessor(
        storage=storage,
        embedder=embedder,
    )

    backend = BackendClient()

    photo_id = "1cb90967-54e2-4e3c-88fc-e70351686e82"

    storage_key = (
        "folders/"
        "a4a41d62-e6b1-43f0-9f58-fb5d3b9f3c06/"
        "photos/"
        "1c5ebab5-43bf-4c11-b32c-61c6da4e850a"
    )

    faces = processor.process(storage_key)

    print(f"Detected {len(faces)} face(s)")

    for index, face in enumerate(faces, start=1):
        print(f"\nFace {index}")
        print("Confidence:", face.confidence)
        print("Landmarks:", face.landmarks)
        print("Embedding:", face.embedding.shape)

    backend.save_faces(
        photo_id=photo_id,
        faces=[
            serialize_face(face)
            for face in faces
        ],
    )

    print("Faces saved successfully.")

def serialize_face(face):
    return {
        "x1": face.x1,
        "y1": face.y1,
        "x2": face.x2,
        "y2": face.y2,
        "confidence": face.confidence,
        "landmarks": face.landmarks,
        "embedding": face.embedding.tolist(),
    }

if __name__ == "__main__":
    main()