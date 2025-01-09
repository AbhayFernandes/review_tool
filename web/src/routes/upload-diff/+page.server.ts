import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';

export const actions: Actions = {
  default: async ({ request }) => {
    const data = await request.formData();
    const user = data.get('user') as string;
    const diff = data.get('diff') as string;

    if (!user || !diff) {
      return fail(400, { error: 'Missing user or diff' });
    }

    // For now, do a fetch to /diffs
    // This is a placeholder. We'll do real logic once gRPC is in place
    const res = await fetch('http://localhost:3000/diffs', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ user, diff })
    });

    if (!res.ok) {
      return fail(res.status, { error: 'Failed to upload diff' });
    }

    return { success: true };
  }
};
