<script lang="ts">
  import { onMount } from "svelte";
  import { handleCallback } from "$lib/auth";
  import { goto } from "$app/navigation";

  let message = "";

  onMount(async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get("code");
    const state = urlParams.get("state");

    if (code && state) {
      const success = await handleCallback(code, state);
      if (success) {
        goto("/");
      } else {
        message = "Authentication failed. Please try again.";
      }
    } else {
      message = "Invalid OAuth callback. Missing code or state.";
    }
  });
</script>

<h1>{message}</h1>
