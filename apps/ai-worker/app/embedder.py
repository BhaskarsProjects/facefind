# app/embedder.py

import numpy as np

from insightface.app import FaceAnalysis

from models.face import DetectedFace


class FaceEmbedder:

    def __init__(self):
        self.app = FaceAnalysis(
            name="buffalo_l",
            providers=["CPUExecutionProvider"],
        )

        self.app.prepare(
            ctx_id=0,
            det_size=(640, 640),
        )

    def embed(
        self,
        image: np.ndarray,
    ) -> list[DetectedFace]:

        faces = self.app.get(image)

        results: list[DetectedFace] = []

        for face in faces:
            bbox = face.bbox

            results.append(
                DetectedFace(
                    x1=float(bbox[0]),
                    y1=float(bbox[1]),
                    x2=float(bbox[2]),
                    y2=float(bbox[3]),
                    confidence=float(face.det_score),
                    landmarks=[
                        [
                            float(point[0]),
                            float(point[1]),
                        ]
                        for point in face.kps
                    ],
                    embedding=face.embedding,
                )
            )

        return results