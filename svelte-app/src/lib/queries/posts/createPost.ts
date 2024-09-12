import { createMutation, useQueryClient } from "@tanstack/svelte-query";
import { postsApi } from "$lib/api/posts";
import type { CreatePostRequest, GetPostResponse, Post } from "$lib/api/posts";
import { QK_POSTS } from "./consts";

export function useCreatePost() {
  const queryClient = useQueryClient();

  return createMutation<GetPostResponse, Error, CreatePostRequest>({
    mutationFn: (postData: CreatePostRequest) => postsApi.createPost(postData),
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({ queryKey: [QK_POSTS] });
    },
  });
}
