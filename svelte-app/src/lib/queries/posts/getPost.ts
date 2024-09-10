import { createQuery, type QueryObserverResult } from "@tanstack/svelte-query";
import { postsApi } from "$lib/api/posts";
import type {
  GetPostParams,
  ParsedGetPostResponse,
  ParsedIncludesData,
  ParsedUserInteractions,
} from "$lib/api/posts";
import type { Post } from "$lib/api";
import { QK_POST } from "./consts";
import { derived, type Readable } from "svelte/store";

export interface ProcessedGetPostData {
  post: Post;
  includes: ParsedIncludesData;
}

export interface GetPostQueryResult {
  data: ProcessedGetPostData;
  query: QueryObserverResult<ParsedGetPostResponse, Error>;
}

export function useGetPost(
  id: string,
  params: GetPostParams = {},
): Readable<GetPostQueryResult> {
  const query = createQuery<ParsedGetPostResponse, Error>({
    queryKey: [QK_POST, id, params],
    queryFn: () => postsApi.getPost(id, params),
    enabled: !!id,
  });

  return derived(query, ($query): GetPostQueryResult => {
    const processedData: ProcessedGetPostData = {
      post: $query.data?.post ?? null,
      includes: $query.data?.includes,
    };

    return {
      data: processedData,
      query: $query,
    };
  });
}
