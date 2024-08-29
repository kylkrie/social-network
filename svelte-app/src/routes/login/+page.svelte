<script lang="ts">
  import { startOAuth } from "$lib/auth";
</script>

<button on:click={startOAuth}>Login with OAuth</button>
