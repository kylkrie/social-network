<script lang="ts">
  export let tabs: string[] = [];
  export let activeTab: string = tabs[0] || "";
</script>

<div class="tab-header">
  {#each tabs as tab}
    <button
      class="tab-button mx-4"
      class:active={activeTab === tab}
      on:click={() => (activeTab = tab)}
    >
      {tab}
    </button>
  {/each}
</div>

<style>
  .tab-header {
    display: flex;
    border-bottom: 1px solid var(--border);
  }

  .tab-button {
    padding: 0.5rem 1rem;
    border: none;
    background: none;
    cursor: pointer;
    font-size: 1rem;
    color: var(--text-secondary);
    transition: color 0.3s ease;
  }

  .tab-button:hover {
    color: var(--text-secondary);
  }

  .tab-button.active {
    color: var(--text);
    border-bottom: 2px solid var(--primary);
  }
</style>
