<script lang="ts" module>
  export interface Tab {
    name: string;
    component: any;
  }
</script>
<script lang="ts">
  export let tabs: Tab[] = [];
  export let activeTab: string = tabs[0]?.name || "";
</script>

<div class="tab-view">
  <div class="tab-header">
    {#each tabs as tab}
      <button
        class="tab-button"
        class:active={activeTab === tab.name}
        on:click={() => (activeTab = tab.name)}
      >
        {tab.name}
      </button>
    {/each}
  </div>
  <div class="tab-content">
    {#each tabs as tab}
      {#if activeTab === tab.name}
        <svelte:component this={tab.component} />
      {/if}
    {/each}
  </div>
</div>

<style>
  .tab-view {
    width: 100%;
  }

  .tab-header {
    display: flex;
    border-bottom: 1px solid var(--color-border);
  }

  .tab-button {
    padding: 0.5rem 1rem;
    border: none;
    background: none;
    cursor: pointer;
    font-size: 1rem;
    color: var(--color-text-secondary);
    transition: color 0.3s ease;
  }

  .tab-button:hover {
    color: var(--color-text);
  }

  .tab-button.active {
    color: var(--color-primary);
    border-bottom: 2px solid var(--color-primary);
  }
</style>
