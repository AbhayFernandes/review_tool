<script lang="ts">
  // 1) define interface right here
  interface Review {
    id: string;
    title: string;
    state: string;
  }
  interface MyPageData {
    reviews: Review[];
  }
  
  export let data: MyPageData;
  </script>
  
  <h1>All Reviews</h1>
  {#if data.reviews?.length > 0}
    <ul>
      {#each data.reviews as r}
        <li>
          <a href="/reviews/{r.id}">{r.title} ({r.state})</a>
        </li>
      {/each}
    </ul>
  {:else}
    <p>No reviews found.</p>
  {/if}
  