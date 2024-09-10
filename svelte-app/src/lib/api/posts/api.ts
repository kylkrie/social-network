import { api, cleanUrlParams } from "$lib/api";
import type {
  CreatePostRequest,
  UpdatePostRequest,
  GetPostParams,
  ListPostsParams,
  GetPostResponse,
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
    const response = await api.post(API_PATH, postData);
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

  updatePost: async (
    id: string,
    postData: UpdatePostRequest,
  ): Promise<void> => {
    await api.put(`${API_PATH}/${id}`, postData);
  },

  deletePost: async (id: string): Promise<void> => {
    await api.delete(`${API_PATH}/${id}`);
  },

  listPosts: async (
    params: ListPostsParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}?${queryString}`);
    return parseListPostsResponse(response);
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
};
