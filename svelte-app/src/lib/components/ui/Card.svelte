<script lang="ts">
  export let title = "";
</script>

<div class="rounded-lg border bg-background text-text shadow-sm">
  {#if title}
    <div class="flex flex-col space-y-1.5 p-4">
      <h3
        class="text-2xl font-semibold leading-none tracking-tight text-primary"
      >
        {title}
      </h3>
    </div>
  {/if}
  <div class="p-4">
    <slot />
  </div>
</div>
