import { fetchLabResults } from '@/app/labView/result/service/serviceGetLabResult';
import { LabResultResponse } from '@/app/labView/result/utils/types';

export async function handleGetLabResult(
  cid: string,
  startDate: string,
  endDate: string
): Promise<{ success: boolean; data?: LabResultResponse; message?: string }> {
  if (!cid || cid.length !== 13) {
    return { success: false, message: 'กรุณากรอกเลขบัตรประชาชน 13 หลักให้ครบถ้วน' };
  }

  if (!startDate || !endDate) {
    return { success: false, message: 'กรุณาเลือกช่วงวันที่' };
  }

  if (startDate > endDate) {
    return { success: false, message: 'วันที่เริ่มต้นต้องไม่เกินวันที่สิ้นสุด' };
  }

  return await fetchLabResults(cid, startDate, endDate);
}
