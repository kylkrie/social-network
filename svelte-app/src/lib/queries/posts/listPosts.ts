import { createInfiniteQuery } from '@tanstack/svelte-query';
import { api } from '$lib/api';

interface ListPostsParams {
  limit?: number;
}

export function useListPosts({ limit = 20 }: ListPostsParams = {}) {
  //@ts-ignore
  return createInfiniteQuery({
    queryKey: ['posts'],
    queryFn: ({ pageParam = undefined }) =>
      api.get(`/posts/v1?limit=${limit}${pageParam ? `&cursor=${pageParam}` : ''}`),
    getNextPageParam: (lastPage: any) => lastPage.next_cursor || undefined,
  });
}
