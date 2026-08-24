import os

import requests

class BackendClient:
    def __init__(self):
        self.base_url = os.environ["BACKEND_URL"].rstrip("/")
        self.internal_token = os.environ["AI_WORKER_TOKEN"]

    def save_faces(
        self,
        photo_id: str,
        faces: list[dict],
    ) -> None:
        response = requests.post(
            f"{self.base_url}/internal/photos/{photo_id}/faces",
            json={
                "faces": faces,
            },
            headers={
                "Authorization": f"Bearer {self.internal_token}",
                "Content-Type": "application/json",
            },
            timeout=60,
        )

        response.raise_for_status()