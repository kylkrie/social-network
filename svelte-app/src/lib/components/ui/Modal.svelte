<script lang="ts">
  import { fade, scale } from "svelte/transition";

  export let isOpen = false;
  export let onClose: () => void;
  export let width = "max-w-md";
</script>

<svelte:window on:keydown={(event) => event.key === "Escape" && onClose()} />

{#if isOpen}
  <div
    class="fixed inset-0 bg-gray-700 bg-opacity-50 z-50 flex items-center justify-center p-4"
    on:click={onClose}
    transition:fade={{ duration: 200 }}
  >
    <div
      class="bg-background rounded-lg shadow-xl {width} w-full"
      on:click|stopPropagation
      transition:scale={{ duration: 200, start: 0.95 }}
    >
      <slot />
    </div>
  </div>
{/if}
