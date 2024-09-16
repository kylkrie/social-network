<script>
  import { useListFeed } from "$lib/queries";
  import PageContent from "../layout/PageContent.svelte";
  import PostFeed from "../post/PostFeed.svelte";
  import FeedHeader from "./FeedHeader.svelte";

  const feedData = useListFeed();
</script>

<PageContent>
  <FeedHeader />
  <PostFeed postData={feedData} />
</PageContent>
