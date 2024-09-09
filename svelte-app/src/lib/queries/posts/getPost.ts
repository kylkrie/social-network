import { createQuery } from "@tanstack/svelte-query";
import { postsApi } from "$lib/api/posts";
import type { GetPostParams, GetPostResponse } from "$lib/api/posts";
import { QK_POST } from "./consts";

export function useGetPost(id: string, params: GetPostParams = {}) {
  return createQuery<GetPostResponse, Error>({
    queryKey: [QK_POST, id, params],
    queryFn: () => postsApi.getPost(id, params),
    enabled: !!id,
  });
}
