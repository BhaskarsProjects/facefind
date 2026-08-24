// import axios from "axios";

// import { env } from "@/app/config/env";

// export const apiClient = axios.create({
//   baseURL: env.apiUrl + "/api/v1",
//   withCredentials: true,
//   headers: {
//     "Content-Type": "application/json",
//   },
// });

import axios from "axios";
import { env } from "@/app/config/env";

export const apiClient = axios.create({
  baseURL: env.apiUrl + "/api/v1",
  withCredentials: true,
});