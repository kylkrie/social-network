<script lang="ts">
  export let variant: "default" | "outline" = "default";
  export let size: "sm" | "md" | "lg" = "md";
  export let type: "button" | "submit" | "reset" = "button";
  export let disabled = false;
</script>

<button
  {type}
  {disabled}
  class="inline-flex items-center justify-center rounded-md font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary focus-visible:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none"
  class:bg-primary={variant === "default"}
  class:text-white={variant === "default"}
  class:hover:bg-primary-dark={variant === "default"}
  class:border={variant === "outline"}
  class:border-primary={variant === "outline"}
  class:text-primary={variant === "outline"}
  class:hover:bg-primary-light={variant === "outline"}
  class:hover:text-primary-dark={variant === "outline"}
  class:text-sm={size === "sm"}
  class:px-3={size === "sm"}
  class:py-1={size === "sm"}
  class:text-base={size === "md"}
  class:px-4={size === "md"}
  class:py-2={size === "md"}
  class:text-lg={size === "lg"}
  class:px-6={size === "lg"}
  class:py-3={size === "lg"}
  on:click
>
  <slot />
</button>
