<script lang="ts">
  import { theme, toggleTheme } from "$lib/theme";
  import { Sun, Moon } from "lucide-svelte";
  import { onMount } from "svelte";

  let currentTheme: "light" | "dark";

  onMount(() => {
    // Ensure the initial theme is applied
    theme.subscribe((value) => {
      currentTheme = value;
    });
  });

  function handleToggle() {
    toggleTheme();
  }
</script>

<button
  on:click={handleToggle}
  class="p-2 rounded-full bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-200 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
  aria-label="Toggle theme"
>
  {#if currentTheme === "light"}
    <Moon size={20} />
  {:else}
    <Sun size={20} />
  {/if}
</button>

<style>
  button {
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
