<script lang="ts">
  export let title = "";
</script>

<div class="rounded-lg border bg-surface text-text shadow-sm">
  {#if title}
    <div class="flex flex-col space-y-1.5 p-6">
      <h3
        class="text-2xl font-semibold leading-none tracking-tight text-primary"
      >
        {title}
      </h3>
    </div>
  {/if}
  <div class="p-6 pt-0">
    <slot />
  </div>
</div>
