import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';
import { AuthToken } from '@/global/globalAuth';

export async function updateUserStatus(
  id: number,
  status: boolean
): Promise<{ success: boolean; message?: string }> {
  try {
    const token = AuthToken.getToken();
    const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.USER.UPDATE_USER_STATUS}`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ id, status }),
    });

    const data = await res.json();

    if (!res.ok) {
      return { success: false, message: data.message || 'อัพเดทสถานะไม่สำเร็จ' };
    }

    return { success: true, message: data.message };
  } catch {
    return { success: false, message: 'เกิดข้อผิดพลาดในการเชื่อมต่อ' };
  }
}
