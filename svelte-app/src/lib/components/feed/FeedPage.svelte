<script lang="ts">
  import FeedHeader from "./FeedHeader.svelte";
  import FeedPosts from "./FeedPosts.svelte";
</script>

<div class="profile-page border-x border-border">
  <FeedHeader />
  <FeedPosts />
</div>

<style>
  .profile-page {
    max-width: 800px;
    margin: 0 auto;
  }
</style>
