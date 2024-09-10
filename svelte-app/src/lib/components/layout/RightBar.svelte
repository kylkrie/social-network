<script lang="ts">
  const features = [
    "OpenID Connect authentication",
    "Create and share posts with your network",
    "Like and bookmark posts",
    "Reply to posts and engage in conversations",
    "Quote and repost interesting content",
    "Customizable user profiles",
    "Light/Dark mode themes",
    "Responsive design for mobile and desktop",
  ];
</script>

<div class="p-4">
  <h2 class="text-xl font-bold mb-4">Yabro Features</h2>
  <ul class="space-y-2">
    {#each features as feature}
      <li class="flex items-start">
        <span class="text-primary mr-2">•</span>
        <span>{feature}</span>
      </li>
    {/each}
  </ul>

  <div class="mt-6">
    <h3 class="text-lg font-semibold mb-2">Discover More</h3>
    <p class="text-sm text-text-secondary">
      Explore these features and more by using Yabro Social. Connect with
      others, share your thoughts, and join the conversation!
    </p>
  </div>
</div>
