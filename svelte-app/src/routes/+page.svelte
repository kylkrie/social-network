<script lang="ts">
  import { auth } from "$lib/stores";
  import FeedPage from "$lib/components/feed/FeedPage.svelte";
  import LandingPage from "$lib/components/LandingPage.svelte";
</script>

{#if !$auth}
  <LandingPage />
{:else}
  <FeedPage />
{/if}
