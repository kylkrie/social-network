<script lang="ts">
  import TabView from "$lib/components/ui/TabView.svelte";
  import ProfileInfo from "$lib/components/profile/ProfileInfo.svelte";
  import UserPosts from "$lib/components/profile/UserPosts.svelte";
  import UserMedia from "$lib/components/profile/UserMedia.svelte";
  import UserLikes from "$lib/components/profile/UserLikes.svelte";
  import UserReplies from "$lib/components/profile/UserReplies.svelte";

  const tabs = [
    { name: "Posts", component: UserPosts },
    { name: "Replies", component: UserReplies },
    { name: "Media", component: UserMedia },
    { name: "Likes", component: UserLikes },
  ];

  let activeTab = "Posts";
</script>

<div class="profile-page border-x border-border">
  <ProfileInfo />
  <TabView {tabs} bind:activeTab />
</div>

<style>
  .profile-page {
    max-width: 800px;
    margin: 0 auto;
  }
</style>
