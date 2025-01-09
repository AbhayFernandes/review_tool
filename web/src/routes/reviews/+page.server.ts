import type { PageServerLoad } from './$types';

// Define the Review interface to model the data structure
interface Review {
  id: string;
  title: string;
  state: string;
}

// Export an interface for PageData to ensure SvelteKit knows the structure of `data`
export interface PageData {
  reviews: Review[];
}

// Load function fetches data from the API and returns it
export const load: PageServerLoad = async ({ fetch }) => {
  // TODO: Replace this with a gRPC API call once the backend is implemented
  const res = await fetch('/api/reviews');
  const reviews = (await res.json()) as Review[];

  return { reviews }; // Return the reviews array to the client
};
