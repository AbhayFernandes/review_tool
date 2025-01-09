import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

// When the gRPC API is fully implemented, this function will do actual work
async function uploadDiffToGrpc(diff: string, user: string) {
  // TODO: real gRPC logic goes here
  return { success: true }; 
}

export const POST: RequestHandler = async ({ request }) => {
  const body = await request.json();
  const diff = body.diff;
  const user = body.user;

  if (!diff || !user) {
    return new Response('Missing diff or user', { status: 400 });
  }

  // call our placeholder function
  const result = await uploadDiffToGrpc(diff, user);

  return json(result);
};
