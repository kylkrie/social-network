import { createInfiniteQuery } from '@tanstack/svelte-query';
import { postsApi } from '$lib/api/posts';
import type { ListPostsParams, ListPostsResponse } from '$lib/api/posts';
import { QK_POSTS } from './consts';

export function useListPosts(params: ListPostsParams = {}) {
  return createInfiniteQuery<ListPostsResponse, Error>({
    queryKey: [QK_POSTS, 'post', params.conversation_id, params.replies],
    queryFn: ({ pageParam = undefined }) => {
      return postsApi.listPosts({
        ...params,
        cursor: pageParam as string | undefined
      })
    },
    getNextPageParam: (lastPage) => lastPage.next_cursor || undefined,
    initialPageParam: undefined,
  });
}
