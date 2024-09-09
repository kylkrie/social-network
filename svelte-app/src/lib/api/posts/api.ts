import { api, cleanUrlParams } from "$lib/api";
import type {
  CreatePostRequest,
  UpdatePostRequest,
  GetPostParams,
  ListPostsParams,
  ListPostsResponse,
  GetPostResponse,
} from "./dtos";

const API_PATH = "/posts";

export const postsApi = {
  /**
   * Create a new post
   */
  createPost: async (postData: CreatePostRequest): Promise<GetPostResponse> => {
    const response = await api.post(API_PATH, postData);
    return response;
  },

  /**
   * Get a post by ID
   */
  getPost: async (
    id: string,
    params: GetPostParams = {},
  ): Promise<GetPostResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}/${id}?${queryString}`);
    return response;
  },

  /**
   * Update an existing post
   */
  updatePost: async (
    id: string,
    postData: UpdatePostRequest,
  ): Promise<void> => {
    await api.put(`${API_PATH}/${id}`, postData);
  },

  /**
   * Delete a post
   */
  deletePost: async (id: string): Promise<void> => {
    await api.delete(`${API_PATH}/${id}`);
  },

  /**
   * List posts
   */
  listPosts: async (
    params: ListPostsParams = {},
  ): Promise<ListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}?${queryString}`);
    return response;
  },
};
