import { api, cleanUrlParams, publicApi } from "$lib/api";
import type {
  CreatePostRequest,
  UpdatePostRequest,
  GetPostParams,
  GetPostResponse,
  ListRepliesParams,
} from "./dtos";
import {
  parseGetPostResponse,
  parseListPostsResponse,
  type ParsedGetPostResponse,
  type ParsedListPostsResponse,
} from "./parsed";

const API_PATH = "/posts";

export const postsApi = {
  createPost: async (postData: CreatePostRequest): Promise<GetPostResponse> => {
    const formData = new FormData();
    formData.append("content", postData.content);
    if (postData.reply_to_post_id) {
      formData.append("reply_to_post_id", postData.reply_to_post_id);
    }
    if (postData.quote_post_id) {
      formData.append("quote_post_id", postData.quote_post_id);
    }
    if (postData.media) {
      for (let i = 0; i < postData.media.length; i++) {
        formData.append("media", postData.media[i]);
      }
    }
    const response = await api.post(API_PATH, formData);
    return response;
  },

  getPost: async (
    id: string,
    params: GetPostParams = {},
  ): Promise<ParsedGetPostResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}/${id}?${queryString}`);
    return parseGetPostResponse(response);
  },

  getPublicPost: async (
    id: string,
    params: GetPostParams = {},
  ): Promise<ParsedGetPostResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await publicApi.get(`/posts/${id}?${queryString}`);
    return parseGetPostResponse(response);
  },

  updatePost: async (
    id: string,
    postData: UpdatePostRequest,
  ): Promise<void> => {
    await api.put(`${API_PATH}/${id}`, postData);
  },

  deletePost: async (id: string): Promise<void> => {
    await api.delete(`${API_PATH}/${id}`);
  },

  likePost: async (id: string): Promise<void> => {
    await api.post(`${API_PATH}/${id}/likes`);
  },

  unlikePost: async (id: string): Promise<void> => {
    await api.delete(`${API_PATH}/${id}/likes`);
  },

  bookmarkPost: async (id: string): Promise<void> => {
    await api.post(`${API_PATH}/${id}/bookmarks`);
  },

  unbookmarkPost: async (id: string): Promise<void> => {
    await api.delete(`${API_PATH}/${id}/bookmarks`);
  },

  listReplies: async (
    postId: string,
    params: ListRepliesParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(
      `${API_PATH}/${postId}/replies?${queryString}`,
    );
    return parseListPostsResponse(response);
  },
};
