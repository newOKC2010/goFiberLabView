import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';
import { AuthToken } from '@/global/globalAuth';
import { LabResultResponse } from '@/app/labView/result/utils/types';

export async function fetchLabResults(
  cid: string,
  startDate: string,
  endDate: string
): Promise<{ success: boolean; data?: LabResultResponse; message?: string }> {
  try {
    const token = AuthToken.getToken();
    const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.LAB.VIEW_RESULTS}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({ cid, start_date: startDate, end_date: endDate }),
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
