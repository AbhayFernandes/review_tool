<script lang="ts">
    import type { ActionData } from './$types';
    let { form }: { form: ActionData } = $props();
  </script>
  
  <h1>Upload Diff</h1>
  <form method="POST" action="?/default">
    <label for="user">User</label>
    <input type="text" id="user" name="user" required />
  
    <label for="diff">Diff</label>
    <textarea id="diff" name="diff" rows="8" required></textarea>
  
    <button>Submit</button>
  </form>
  
  {#if form?.success}
    <p style="color: green;">Diff uploaded successfully!</p>
  {/if}
  {#if form?.error}
    <p style="color: red;">Error: {form.error}</p>
  {/if}
  