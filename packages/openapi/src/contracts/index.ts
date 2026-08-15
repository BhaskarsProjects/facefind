import { initContract } from "@ts-rest/core";
import { healthContract } from "./health.js";
import { nameContract } from "./name.js";

const c = initContract();

export const apiContract = c.router({
  Health: healthContract,
  Name: nameContract,
});