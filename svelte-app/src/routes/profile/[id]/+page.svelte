<script lang="ts">
  import TabView from "$lib/components/ui/TabView.svelte";
  import ProfileInfo from "$lib/components/profile/ProfileInfo.svelte";
  import UserMedia from "$lib/components/profile/UserMedia.svelte";
  import { page } from "$app/stores";
  import { useListPosts } from "$lib/queries";
  import PostFeed from "$lib/components/post/PostFeed.svelte";
  import PageContent from "$lib/components/layout/PageContent.svelte";
  import { useUserLikes } from "$lib/queries/users/listLikes";

  const tabs = ["Posts", "Replies", "Media", "Likes"];
  let activeTab = "Posts";

  $: username = $page.params.id;

  $: postsQuery = activeTab === "Posts" ? useListPosts(username) : undefined;
  $: repliesQuery =
    activeTab === "Replies"
      ? useListPosts(username, { replies: true })
      : undefined;
  $: likesQuery = activeTab === "Likes" ? useUserLikes(username) : undefined;

  $: if (username) {
    // Reset to Posts tab when username changes
    activeTab = "Posts";
  }
</script>

<PageContent>
  <ProfileInfo profile={username} />
  <TabView {tabs} bind:activeTab />
  {#if activeTab === "Posts" && postsQuery}
    <PostFeed postData={postsQuery} />
  {:else if activeTab === "Replies" && repliesQuery}
    <PostFeed postData={repliesQuery} />
  {:else if activeTab === "Media"}
    <UserMedia />
  {:else if activeTab === "Likes" && likesQuery}
    <PostFeed postData={likesQuery} />
  {/if}
</PageContent>
