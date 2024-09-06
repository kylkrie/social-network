import { api } from '$lib/api';
import type {
  User,
  GetUserParams,
  GetCurrentUserParams,
  UpdateUserRequest
} from './dtos';

const API_PATH = '/users/v1';

export const usersApi = {
  /**
   * Get a user by username
   */
  getUser: async (username: string, params: GetUserParams = {}): Promise<User> => {
    const queryString = new URLSearchParams(params as any).toString();
    const response = await api.get(`${API_PATH}/${username}?${queryString}`);
    return response.data;
  },

  /**
   * Get the current user's profile
   */
  getCurrentUser: async (params: GetCurrentUserParams = {}): Promise<User> => {
    const queryString = new URLSearchParams(params as any).toString();
    const response = await api.get(`${API_PATH}?${queryString}`);
    return response.data;
  },

  /**
   * Update the current user's profile
   */
  updateCurrentUser: async (userData: UpdateUserRequest): Promise<void> => {
    await api.put(API_PATH, userData);
  }
};
