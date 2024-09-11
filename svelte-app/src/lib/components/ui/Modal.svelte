<script lang="ts">
  import { fade, fly } from "svelte/transition";
  export let isOpen = false;
  export let onClose: () => void;
  export let width = "max-w-md";
</script>

<svelte:window on:keydown={(event) => event.key === "Escape" && onClose()} />

{#if isOpen}
  <div
    class="fixed inset-0 bg-gray-700 bg-opacity-50 z-50 flex items-start sm:items-center justify-center p-0 sm:p-4 overflow-y-auto"
    on:click={onClose}
    transition:fade={{ duration: 200 }}
  >
    <div
      class="bg-background w-full min-h-screen sm:min-h-0 sm:rounded-lg sm:shadow-xl {width}"
      on:click|stopPropagation
      transition:fly={{ y: 100, duration: 200 }}
    >
      <div class="max-h-screen overflow-y-auto">
        <slot />
      </div>
    </div>
  </div>
{/if}
