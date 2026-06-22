import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';
import { AuthToken } from '@/global/globalAuth';
import { GetUsersResponse } from '@/app/labView/users/utils/types';

export async function fetchUsers(): Promise<{ success: boolean; data?: GetUsersResponse; message?: string }> {
  try {
    const token = AuthToken.getToken();
    const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.USER.VIEW_USER}`, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    const data = await res.json();

    if (!res.ok) {
      return { success: false, message: data.message || 'ไม่สามารถดึงข้อมูลได้' };
    }

    return { success: true, data };
  } catch {
    return { success: false, message: 'เกิดข้อผิดพลาดในการเชื่อมต่อ' };
  }
}
