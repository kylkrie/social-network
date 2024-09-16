import {
  createInfiniteQuery,
  type InfiniteData,
  type InfiniteQueryObserverResult,
} from "@tanstack/svelte-query";
import type {
  ParsedIncludesData,
  ParsedListPostsResponse,
} from "$lib/api/posts";
import { usersApi, type ListPostsParams, type Post } from "$lib/api";
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
  username: string,
  params: ListPostsParams = {},
): Readable<ListPostsQueryResult> {
  const query = createInfiniteQuery<ParsedListPostsResponse, Error>({
    queryKey: [
      QK_POSTS,
      "post",
      username,
      params.conversation_id,
      params.replies,
    ],
    queryFn: ({ pageParam = undefined }) => {
      return usersApi.listPosts(username, {
        ...params,
        cursor: pageParam as string | undefined,
      });
    },
    getNextPageParam: (lastPage) => lastPage.nextCursor || undefined,
    initialPageParam: undefined,
  });

  return derived(query, processListPostsQuery);
}

export function processListPostsQuery(
  query: InfiniteQueryObserverResult<
    InfiniteData<ParsedListPostsResponse, unknown>,
    Error
  >,
): ListPostsQueryResult {
  const processedData: ProcessedListPostsData = {
    posts: query.data?.pages.flatMap((page) => page.posts) ?? [],

    includes: {
      posts:
        query.data?.pages.reduce(
          (acc, page) => ({ ...acc, ...page.includes.posts }),
          {},
        ) ?? {},
      users:
        query.data?.pages.reduce(
          (acc, page) => ({ ...acc, ...page.includes.users }),
          {},
        ) ?? {},
      likedPosts:
        query.data?.pages.reduce(
          (acc, page) => ({ ...acc, ...page.includes.likedPosts }),
          {},
        ) ?? {},
      bookmarkedPosts:
        query.data?.pages.reduce(
          (acc, page) => ({ ...acc, ...page.includes.bookmarkedPosts }),
          {},
        ) ?? {},
      media:
        query.data?.pages.reduce(
          (acc, page) => ({ ...acc, ...page.includes.media }),
          {},
        ) ?? {},
    },
  };

  return {
    data: processedData,
    query: query,
  };
}
