CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE faces (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    photo_id UUID NOT NULL
        REFERENCES photos(id)
        ON DELETE CASCADE,

    x1 REAL NOT NULL,
    y1 REAL NOT NULL,
    x2 REAL NOT NULL,
    y2 REAL NOT NULL,

    confidence REAL NOT NULL,

    landmarks JSONB NOT NULL,

    embedding VECTOR(512) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_faces_photo_id
    ON faces(photo_id);