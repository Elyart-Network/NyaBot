import { extendFunctions } from "@/loader";
import type { NextRequest } from "next/server";
import { checkEnvInit, commonResponse, jsonResponse } from "../../../core/nextkit/common";

export const config = {
  runtime: 'edge',
}

export default async function handler(req: NextRequest) {
  if (!checkEnvInit) { return await commonResponse(500) } // Check if GOCQ_HTTP_URL is set
  if (req.method != 'POST') { return await commonResponse(405) } // Check if method is POST

  try {
    const json = await req.json()
    await extendFunctions["callbackHandler"](json)
    return await jsonResponse(200, { status: 'ok' })
  } catch (error) {
    console.log(error)
    return await commonResponse(400)
  }
}