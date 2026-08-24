export interface Photo {
  id: string;
  url: string;
}

export interface UploadPhotoResponse {
  photos: Photo[];
}

export interface PhotoUploadError {
  fileName: string;
  message: string;
}

export interface GetFolderPhotosResponse {
  photos: Photo[];
}

export interface FaceSearchResult {
  face_id: string;
  photo_id: string;
  confidence: number;
  distance: number;
  url: string;
}

export interface FaceSearchResponse {
  results: FaceSearchResult[];
}