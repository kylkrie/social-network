<script lang="ts">
  import TabView from "$lib/components/ui/TabView.svelte";
  import ProfileInfo from "$lib/components/profile/ProfileInfo.svelte";
  import UserPosts from "$lib/components/profile/UserPosts.svelte";
  import UserMedia from "$lib/components/profile/UserMedia.svelte";
  import UserLikes from "$lib/components/profile/UserLikes.svelte";
  import UserReplies from "$lib/components/profile/UserReplies.svelte";

  const tabs = ["Posts", "Replies", "Media", "Likes"];

  let activeTab = "Posts";
</script>

<div class="profile-page border-x border-border">
  <ProfileInfo />
  <TabView {tabs} bind:activeTab />
  {#if activeTab === "Posts"}
    <UserPosts />
  {:else if activeTab === "Replies"}
    <UserReplies />
  {:else if activeTab === "Media"}
    <UserMedia />
  {:else if activeTab === "Likes"}
    <UserLikes />
  {/if}
</div>

<style>
  .profile-page {
    max-width: 800px;
    margin: 0 auto;
  }
</style>
