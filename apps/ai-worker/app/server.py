import cv2
import numpy as np

from pydantic import BaseModel
from fastapi import (
    FastAPI,
    File,
    HTTPException,
    UploadFile,
)

from app.embedder import FaceEmbedder
from app.processor import PhotoProcessor
from app.storage import R2Storage
from app.backend_client import BackendClient

import os
from dotenv import load_dotenv


load_dotenv()


app = FastAPI(
    title="FaceFind ML Service",
)


class ProcessPhotoRequest(BaseModel):
    photo_id: str
    storage_key: str


# Load dependencies once when the service starts.
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


@app.get("/health")
def health():
    return {
        "status": "ok",
    }


@app.post("/embed")
async def embed(
    image: UploadFile = File(...),
):
    image_bytes = await image.read()

    buffer = np.frombuffer(
        image_bytes,
        dtype=np.uint8,
    )

    decoded = cv2.imdecode(
        buffer,
        cv2.IMREAD_COLOR,
    )

    if decoded is None:
        raise HTTPException(
            status_code=400,
            detail="failed to decode image",
        )

    faces = embedder.embed(decoded)

    if len(faces) == 0:
        raise HTTPException(
            status_code=422,
            detail="no face detected",
        )

    if len(faces) > 1:
        raise HTTPException(
            status_code=422,
            detail="exactly one face is required",
        )

    face = faces[0]

    return {
        "confidence": face.confidence,
        "embedding_dimension": len(face.embedding),
        "embedding": face.embedding.tolist(),
    }


@app.post("/process")
async def process_photo(
    request: ProcessPhotoRequest,
):
    faces = processor.process(
        request.storage_key,
    )

    backend.save_faces(
        photo_id=request.photo_id,
        faces=[
            serialize_face(face)
            for face in faces
        ],
    )

    return {
        "photo_id": request.photo_id,
        "face_count": len(faces),
    }


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