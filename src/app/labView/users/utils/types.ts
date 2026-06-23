export interface UserItem {
  id: number;
  full_name: string;
  status: boolean;
  created_at: string;
}

export interface GetUsersResponse {
  success: boolean;
  message: string;
  data: UserItem[];
  total_count: number;
  total_pages: number;
  current_page: number;
}
