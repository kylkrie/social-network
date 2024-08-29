<script lang="ts">
  export let variant: "default" | "destructive" = "default";
  export let title = "";
</script>

<div
  role="alert"
  class="relative w-full rounded-lg border p-4"
  class:bg-background={variant === "default"}
  class:text-foreground={variant === "default"}
  class:bg-destructive={variant === "destructive"}
  class:text-destructive-foreground={variant === "destructive"}
>
  {#if title}
    <h5 class="mb-1 font-medium leading-none tracking-tight">{title}</h5>
  {/if}
  <div class="text-sm [&_p]:leading-relaxed">
    <slot />
  </div>
</div>
