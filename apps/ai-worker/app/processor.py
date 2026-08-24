import cv2
import numpy as np

from app.embedder import FaceEmbedder
from app.storage import R2Storage
from models.face import DetectedFace

class PhotoProcessor:

    def __init__(
        self,
        storage: R2Storage,
        embedder: FaceEmbedder,
    ):
        self.storage = storage
        self.embedder = embedder

    def process(
        self,
        storage_key: str,
    ) -> list[DetectedFace]:

        image_bytes = self.storage.download(
            storage_key,
        )

        image = np.frombuffer(
            image_bytes,
            dtype=np.uint8,
        )

        decoded = cv2.imdecode(
            image,
            cv2.IMREAD_COLOR,
        )

        if decoded is None:
            raise ValueError(
                f"failed to decode image: {storage_key}"
            )

        return self.embedder.embed(decoded)

    def process_bytes(
        self,
        image_bytes: bytes,
    ) -> list[DetectedFace]:
        image = np.frombuffer(
            image_bytes,
            dtype=np.uint8,
        )

        decoded = cv2.imdecode(
            image,
            cv2.IMREAD_COLOR,
        )

        if decoded is None:
            raise ValueError(
                "failed to decode image"
            )

        return self.embedder.embed(decoded)