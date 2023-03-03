import type { NextRequest } from "next/server";
import { checkEnvInit, commonResponse, jsonResponse } from "@/../core/utils/common";

export const config = {
  runtime: 'edge',
}

export default async function handler(req: NextRequest) {
  if (!checkEnvInit) { return await commonResponse(500) } // Check if GOCQ_HTTP_URL is set
  if (req.method != 'POST') { return await commonResponse(405) } // Check if method is POST

  try {
    const json = await req.json()
    console.log(json)
    return await jsonResponse(200, json)
  } catch (error) {
    console.log(error)
    return await commonResponse(400)
  }
}