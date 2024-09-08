
import { createInfiniteQuery } from '@tanstack/svelte-query';
import type { ListPostsResponse } from '$lib/api/posts';
import { QK_POSTS } from './consts';
import { feedsApi, type ListFeedParams } from '$lib/api';

export function useListFeed(params: ListFeedParams = {}) {
  return createInfiniteQuery<ListPostsResponse, Error>({
    queryKey: [QK_POSTS, 'feed'],
    queryFn: ({ pageParam = undefined }) => {
      return feedsApi.listFeedPosts({
        ...params,
        cursor: pageParam as string | undefined
      })
    },
    getNextPageParam: (lastPage) => lastPage.next_cursor || undefined,
    initialPageParam: undefined,
  });
}
