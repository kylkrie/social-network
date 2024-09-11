<script lang="ts">
  export let type: "text" | "password" | "email" | "number" = "text";
  export let placeholder = "";
  export let value = "";
  export let id: string | undefined = undefined;
  export let disabled: boolean | undefined = undefined;
</script>

<input
  {type}
  {placeholder}
  {id}
  {disabled}
  bind:value
  class="flex h-10 w-full rounded-md border border-input bg-surface px-3 py-2 text-sm text-text ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-text-secondary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
/>
