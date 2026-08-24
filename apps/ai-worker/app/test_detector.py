import cv2

from app.detector import FaceDetector


image = cv2.imread("test5.jpg")

if image is None:
    raise ValueError("Failed to read test image.")

detector = FaceDetector()

faces = detector.detect(image)

print(f"Detected {len(faces)} face(s)")

for index, face in enumerate(faces, start=1):
    print(f"\nFace {index}")
    print(
        f"Bounding box: "
        f"{face.x1}, {face.y1}, "
        f"{face.x2}, {face.y2}"
    )
    print(f"Confidence: {face.confidence}")
    print(f"Landmarks: {face.landmarks}")