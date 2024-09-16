<script lang="ts">
  import { page } from "$app/stores";
  import PageContent from "$lib/components/layout/PageContent.svelte";
  import PostCard from "$lib/components/post/PostCard.svelte";
  import PostFeed from "$lib/components/post/PostFeed.svelte";
  import { useGetPost, useListPosts } from "$lib/queries";
  import { useListReplies } from "$lib/queries/posts/listReplies";
  import { buildPostData, getQuoteForPost, getReplyForPost } from "$lib/util";
  import { ArrowLeft } from "lucide-svelte";

  $: postId = $page.params.id;
  $: getPost = useGetPost(postId);
  $: post = $getPost.data.post;
  $: includes = $getPost.data.includes;
  $: postData = post ? buildPostData(post, includes) : undefined;
  $: replyPost = post ? getReplyForPost(post, includes) : undefined;
  $: quotePost = post ? getQuoteForPost(post, includes) : undefined;
  $: feed = useListReplies(postId);

  function goBack() {
    history.back();
  }
</script>

<PageContent>
  <div class="text-text p-4 font-bold text-lg flex items-center">
    <button on:click={goBack} class="mr-4 hover:text-primary-light">
      <ArrowLeft size={24} />
    </button>
    <h1>Post</h1>
  </div>
  {#if postData}
    {#if replyPost}
      <PostCard data={replyPost} variant="reply_start" />
      <PostCard
        data={postData}
        variant="reply_end"
        clickable={false}
        showReplyInput={true}
      />
    {:else}
      <PostCard
        data={postData}
        quote={quotePost}
        clickable={false}
        showReplyInput={true}
      />
    {/if}
  {/if}
  {#if feed && $feed.data.posts?.length > 0}
    <PostFeed postData={feed} showReplySource={false} />
  {/if}
</PageContent>
