<script lang="ts">
  export let data;

  // data.reviewId is from the +page.server.ts
  // We might fetch the single review or show it directly. 
  // For now, let's just display the ID.
</script>

<h1>Single Review</h1>
<p>Review ID: {data.reviewId}</p>
