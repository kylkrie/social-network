<script lang="ts">
  import TabView from "$lib/components/ui/TabView.svelte";
  import ProfileInfo from "$lib/components/profile/ProfileInfo.svelte";
  import UserMedia from "$lib/components/profile/UserMedia.svelte";
  import UserLikes from "$lib/components/profile/UserLikes.svelte";
  import { page } from "$app/stores";
  import { useListPosts } from "$lib/queries";
  import PostFeed from "$lib/components/post/PostFeed.svelte";
  import type { Readable } from "svelte/store";
  import type { ListPostsQueryResult } from "$lib/queries";
  import PageContent from "$lib/components/layout/PageContent.svelte";

  $: username = $page.params.id;
  const tabs = ["Posts", "Replies", "Media", "Likes"];
  let activeTab = "Posts";

  let postsQuery: Readable<ListPostsQueryResult>;
  let repliesQuery: Readable<ListPostsQueryResult>;

  $: {
    if (activeTab === "Posts" && !postsQuery) {
      postsQuery = useListPosts({ username });
    } else if (activeTab === "Replies" && !repliesQuery) {
      repliesQuery = useListPosts({ username, replies: true });
    }
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
  {:else if activeTab === "Likes"}
    <UserLikes />
  {/if}
</PageContent>
