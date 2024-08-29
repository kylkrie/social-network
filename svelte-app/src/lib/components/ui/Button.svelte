<script lang="ts">
  export let variant: 'default' | 'outline' | 'ghost' = 'default';
  export let size: 'sm' | 'md' | 'lg' = 'md';
  export let type: 'button' | 'submit' | 'reset' = 'button';
</script>


<button
  {type}

  class="inline-flex items-center justify-center rounded-md font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none ring-offset-background"
  class:bg-primary={variant === 'default'}
  class:text-primary-foreground={variant === 'default'}
  class:hover:bg-primary={variant === 'default'}

  class:border={variant === 'outline'}
  class:border-input={variant === 'outline'}
  class:bg-background={variant === 'outline'}
  class:hover:bg-accent={variant === 'outline'}
  class:hover:text-accent-foreground={variant === 'outline'}
  class:hover:bg-accent={variant === 'ghost'}
  class:hover:text-accent-foreground={variant === 'ghost'}
  class:h-9={size === 'sm'}

  class:px-3={size === 'sm'}
  class:h-10={size === 'md'}
  class:px-4={size === 'md'}
  class:h-11={size === 'lg'}

  class:px-8={size === 'lg'}
  on:click

>
  <slot />
</button>
