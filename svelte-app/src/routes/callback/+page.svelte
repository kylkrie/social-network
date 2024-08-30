<script lang="ts">
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import { handleCallback } from "$lib/auth";

  let message = "";

  onMount(async () => {
    if (browser) {
      const urlParams = new URLSearchParams(window.location.search);
      const code = urlParams.get("code");
      const state = urlParams.get("state");

      if (code && state) {
        const success = await handleCallback(code, state);
        if (!success) {
          message = "Authentication failed. Please try again.";
        }
      } else {
        message = "Invalid OAuth callback. Missing code or state.";
      }
    }
  });
</script>

<h1>{message}</h1>
