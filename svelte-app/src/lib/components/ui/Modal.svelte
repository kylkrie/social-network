<!-- src/lib/components/Modal.svelte -->
<script lang="ts">
  import { fade, scale } from "svelte/transition";

  export let isOpen = false;
  export let title = "";
  export let width = "max-w-md";

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === "Escape" && isOpen) {
      isOpen = false;
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

{#if isOpen}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
    on:click={() => (isOpen = false)}
    transition:fade={{ duration: 200 }}
  >
    <div
      class="bg-white dark:bg-gray-800 rounded-lg shadow-xl {width} w-full"
      on:click|stopPropagation
      transition:scale={{ duration: 200, start: 0.95 }}
    >
      {#if title}
        <div class="border-b border-gray-200 dark:border-gray-700 px-6 py-4">
          <h2 class="text-xl font-semibold text-gray-800 dark:text-white">
            {title}
          </h2>
        </div>
      {/if}
      <div class="p-6">
        <slot />
      </div>
    </div>
  </div>
{/if}
