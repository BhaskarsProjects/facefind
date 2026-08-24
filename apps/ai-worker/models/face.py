# models/face.py

from dataclasses import dataclass
import numpy as np

@dataclass
class DetectedFace:
    x1: float
    y1: float
    x2: float
    y2: float
    confidence: float
    landmarks: list[list[float]]
    embedding: np.ndarray