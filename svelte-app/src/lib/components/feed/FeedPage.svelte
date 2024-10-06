<script>
  import { useListFeed, usePublicListFeed } from "$lib/queries";
  import PageContent from "../layout/PageContent.svelte";
  import PostFeed from "../post/PostFeed.svelte";
  import FeedHeader from "./FeedHeader.svelte";
  import { auth } from "$lib/stores";

  $: isAuthenticated = !!$auth?.accessToken;
  $: feedData = isAuthenticated ? useListFeed() : usePublicListFeed();
</script>

<PageContent>
  <FeedHeader />
  <PostFeed postData={feedData} />
</PageContent>
