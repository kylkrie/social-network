import { createQuery } from '@tanstack/svelte-query';
import { postsApi } from '$lib/api/posts';
import type { Post, GetPostParams } from '$lib/api/posts';
import { QK_POST } from './consts';

export function useGetPost(id: number, params: GetPostParams = {}) {
  return createQuery<Post, Error>({
    queryKey: [QK_POST, id, params],
    queryFn: () => postsApi.getPost(id, params),
  });
}
