import { initContract } from "@ts-rest/core";
import { z } from "zod";
import { ZNameResponse } from "@go-backend-boilerplate/zod";

const c = initContract();

export const nameContract = c.router({
  getName: {
    summary: "Get greeting",
    path: "/name/:name",
    method: "GET",
    description: "Returns a greeting for the provided name",
    pathParams: z.object({
      name: z.string(),
    }),
    responses: {
      200: ZNameResponse,
    },
  },
});