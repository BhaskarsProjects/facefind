import { z } from "zod";
import { createFolderSchema } from "../schemas/folder.schema";

export type CreateFolderInput = z.infer<typeof createFolderSchema>;