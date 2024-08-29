<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import { browser } from "$app/environment";
  import { handleCallback } from "$lib/auth";

  let message = "Processing OAuth callback...";

  onMount(async () => {
    if (browser) {
      const urlParams = new URLSearchParams(window.location.search);
      const code = urlParams.get("code");
      const state = urlParams.get("state");

      if (code && state) {
        const success = await handleCallback(code, state);
        if (success) {
          message = "Authentication successful!";
          setTimeout(() => goto("/"), 2000);
        } else {
          message = "Authentication failed. Please try again.";
        }
      } else {
        message = "Invalid OAuth callback. Missing code or state.";
      }
    }
  });
</script>

<h1>{message}</h1>
