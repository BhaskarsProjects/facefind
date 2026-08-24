import cv2

from insightface.app import FaceAnalysis


image = cv2.imread("test5.jpg")

if image is None:
    raise ValueError("Failed to read test image.")


app = FaceAnalysis(
    name="buffalo_l",
    providers=["CPUExecutionProvider"],
)

app.prepare(
    ctx_id=0,
    det_size=(640, 640),
)

faces = app.get(image)

print(f"Detected {len(faces)} face(s)")

for index, face in enumerate(faces, start=1):
    print(f"\nFace {index}")

    print(f"Bounding box: {face.bbox}")
    print(f"Confidence: {face.det_score}")
    print(f"Landmarks: {face.kps}")

    embedding = face.embedding

    print(f"Embedding type: {type(embedding)}")
    print(f"Embedding shape: {embedding.shape}")
    print(f"Embedding dtype: {embedding.dtype}")
    print(f"Embedding norm: {float((embedding ** 2).sum() ** 0.5)}")
    print(f"First 10 values: {embedding[:10]}")