import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ params }) => {
  // We'll receive an `id` from /reviews/[id].
  // Eventually, we might fetch gRPC data about this specific review.
  const { id } = params;
  
  // For now, just return that ID so +page.svelte can display it
  return { reviewId: id };
};

export const actions: Actions = {
  default: async ({ request }) => {
    // TODO: Handle form submission logic
    // For now, do nothing
    return {};
  }
};
