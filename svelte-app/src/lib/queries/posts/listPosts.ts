import {
  createInfiniteQuery,
  type InfiniteData,
  type InfiniteQueryObserverResult,
} from "@tanstack/svelte-query";
import { postsApi } from "$lib/api/posts";
import type {
  ListPostsParams,
  ParsedIncludesData,
  ParsedListPostsResponse,
} from "$lib/api/posts";
import type { Post } from "$lib/api";
import { QK_POSTS } from "./consts";
import { derived, type Readable } from "svelte/store";

export interface ProcessedListPostsData {
  posts: Post[];
  includes: ParsedIncludesData;
}

export interface ListPostsQueryResult {
  data: ProcessedListPostsData;
  query: InfiniteQueryObserverResult<
    InfiniteData<ParsedListPostsResponse, unknown>,
    Error
  >;
}

export function useListPosts(
  params: ListPostsParams = {},
): Readable<ListPostsQueryResult> {
  const query = createInfiniteQuery<ParsedListPostsResponse, Error>({
    queryKey: [QK_POSTS, "post", params.conversation_id, params.replies],
    queryFn: ({ pageParam = undefined }) => {
      return postsApi.listPosts({
        ...params,
        cursor: pageParam as string | undefined,
      });
    },
    getNextPageParam: (lastPage) => lastPage.nextCursor || undefined,
    initialPageParam: undefined,
  });

  return derived(query, ($query): ListPostsQueryResult => {
    const processedData: ProcessedListPostsData = {
      posts: $query.data?.pages.flatMap((page) => page.posts) ?? [],
      includes: {
        posts:
          $query.data?.pages.reduce(
            (acc, page) => ({ ...acc, ...page.includes.posts }),
            {},
          ) ?? {},
        users:
          $query.data?.pages.reduce(
            (acc, page) => ({ ...acc, ...page.includes.users }),
            {},
          ) ?? {},
      },
    };

    return {
      data: processedData,
      query: $query,
    };
  });
}
