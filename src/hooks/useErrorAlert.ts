'use client';

import { useEffect } from 'react';
import { useSearchParams, useRouter } from 'next/navigation';
import { showAlert } from '@/global/globalSwal';

export function useErrorAlert() {
  const searchParams = useSearchParams();
  const router = useRouter();

  useEffect(() => {
    const error = searchParams.get('error');
    
    if (error) {
      switch (error) {
        case 'no_token':
          showAlert('กรุณาเข้าสู่ระบบ', 'ยังไม่มี Token หรือ Token หมดอายุ กรุณาเข้าสู่ระบบใหม่อีกครั้ง', 'warning');
          break;
        case 'auth_failed':
          showAlert('หมดเวลาเซสชัน', 'กรุณาเข้าสู่ระบบใหม่อีกครั้ง', 'warning');
          break;
        case 'network_error':
          showAlert('ข้อผิดพลาดเครือข่าย', 'ไม่สามารถเชื่อมต่อเซิร์ฟเวอร์ได้', 'error');
          break;
        default:
          showAlert('ไม่มีสิทธิ์เข้าถึง', 'คุณไม่มีสิทธิ์เข้าถึงหน้านี้', 'error');
      }
      
      // ลบ error parameter จาก URL
      const newUrl = new URL(window.location.href);
      newUrl.searchParams.delete('error');
      router.replace(newUrl.pathname + newUrl.search, { scroll: false });
    }
  }, [searchParams, router]);
}