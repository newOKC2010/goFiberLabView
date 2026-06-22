import { loginRequest } from '@/app/auth/request/service/serviceReq';
import { showAlert } from '@/global/globalSwal';

export const handleLoginRequest = async (email: string) => {
  try {
    const data = await loginRequest(email);
    showAlert('สำเร็จ', data.message, 'success');
    return { success: true, data };
  } catch (error) {
    const message = error instanceof Error ? error.message : 'ไม่สามารถส่งคำขอเข้าสู่ระบบได้';
    showAlert('เกิดข้อผิดพลาด', message, 'error');
    return { success: false, error: message };
  }
};
