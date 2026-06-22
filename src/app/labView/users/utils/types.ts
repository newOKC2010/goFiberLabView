export interface UserItem {
  id: number;
  full_name: string;
  status: boolean;
  created_at: string;
}

export interface GetUsersResponse {
  success: boolean;
  total: number;
  users: UserItem[];
}
