<script>
  import FeedHeader from "$lib/components/feed/FeedHeader.svelte";
  import PageContent from "$lib/components/layout/PageContent.svelte";
  import PostFeed from "$lib/components/post/PostFeed.svelte";
  import { useGetCurrentUser } from "$lib/queries";
  import { useUserBookmarks } from "$lib/queries/users/listBookmarks";

  $: user = useGetCurrentUser();
  $: feedData = $user?.data ? useUserBookmarks($user.data.username) : undefined;
</script>

<PageContent>
  <div class="text-text border border-x-0 border-t-0 border-border p-2">
    <div class="flex justify-between items-center">
      <div class="flex flex-col">
        <h1 class="text-2xl font-bold">Bookmarks</h1>
        <p class="text-text-tertiary">@{$user?.data?.username}</p>
      </div>
    </div>
  </div>
  {#if feedData}
    <PostFeed postData={feedData} />
  {/if}
</PageContent>
