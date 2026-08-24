export interface Folder {
  id: string;
  owner_sub: string;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface CreateFolderRequest {
  name: string;
}

export interface UpdateFolderRequest {
  name: string;
}