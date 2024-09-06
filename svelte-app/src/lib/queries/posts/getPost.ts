import { createQuery } from '@tanstack/svelte-query';
import { api } from '$lib/api';

export function useGetPost(id: number) {
  return createQuery({
    queryKey: ['post', id],
    queryFn: () => api.get(`/posts/v1/${id}`),
  });
}
