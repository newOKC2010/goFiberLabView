import { showAlert } from '@/global/globalSwal';
import { registerService } from '@/app/auth/component/register/service/serviceRegister';
import { RegisterErrors } from '@/app/auth/component/register/utils/types';

const THAI_NAME_REGEX = /^[ก-๙.\s]+$/;
const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export const onCidChange = (
  value: string,
  setCid: (v: string) => void,
  setErrors: React.Dispatch<React.SetStateAction<RegisterErrors>>
) => {
  const val = value.replace(/\D/g, '').slice(0, 13);
  setCid(val);
  setErrors((prev) => ({ ...prev, cid: '' }));
};

export const onNameChange = (
  value: string,
  setFullName: (v: string) => void,
  setErrors: React.Dispatch<React.SetStateAction<RegisterErrors>>
) => {
  const val = value.replace(/[^ก-๙.\s]/g, '');
  setFullName(val);
  setErrors((prev) => ({ ...prev, fullName: '' }));
};

export const onEmailChange = (
  value: string,
  setEmail: (v: string) => void,
  setErrors: React.Dispatch<React.SetStateAction<RegisterErrors>>
) => {
  setEmail(value);
  setErrors((prev) => ({ ...prev, email: '' }));
};

const emptyErrors = (): RegisterErrors => ({
  cid: '', fullName: '', email: '',
  facilityType: '', facilityCode: '', facilityName: '',
});

export const validate = (
  cid: string,
  fullName: string,
  email: string,
  facilityType: string,
  facilityCode: string,
  facilityName: string,
  setErrors: React.Dispatch<React.SetStateAction<RegisterErrors>>
): boolean => {
  const err = emptyErrors();

  if (!cid) err.cid = 'กรุณากรอกเลขบัตรประชาชน';
  else if (cid.length !== 13) err.cid = 'เลขบัตรประชาชนต้องมี 13 หลัก';

  if (!fullName.trim()) err.fullName = 'กรุณากรอกชื่อ-นามสกุล';
  else if (!THAI_NAME_REGEX.test(fullName.trim())) err.fullName = 'ใส่ได้เฉพาะภาษาไทยและ .';

  if (!email.trim()) err.email = 'กรุณากรอกอีเมล';
  else if (!EMAIL_REGEX.test(email.trim())) err.email = 'รูปแบบอีเมลไม่ถูกต้อง';

  if (!facilityType.trim()) err.facilityType = 'กรุณาเลือกประเภทสถานพยาบาล';
  if (!facilityCode.trim()) err.facilityCode = 'กรุณากรอกรหัสสถานพยาบาล';
  if (!facilityName.trim()) err.facilityName = 'กรุณากรอกชื่อสถานพยาบาล';

  setErrors(err);
  return !err.cid && !err.fullName && !err.email && !err.facilityType && !err.facilityCode && !err.facilityName;
};

export const handleRegisterSubmit = async (
  cid: string,
  fullName: string,
  email: string,
  facilityType: string,
  facilityCode: string,
  facilityName: string,
  setErrors: React.Dispatch<React.SetStateAction<RegisterErrors>>,
  setLoading: (v: boolean) => void,
  onSuccess: () => void
) => {
  if (!validate(cid, fullName, email, facilityType, facilityCode, facilityName, setErrors)) return;

  setLoading(true);
  try {
    const data = await registerService({
      cid,
      full_name: fullName.trim(),
      email: email.trim(),
      facility_type: facilityType.trim(),
      facility_code: facilityCode.trim(),
      facility_name: facilityName.trim(),
    });
    showAlert('สำเร็จ', data.message, 'success');
    onSuccess();
  } catch (error) {
    const message = error instanceof Error ? error.message : 'ไม่สามารถลงทะเบียนได้';
    showAlert('เกิดข้อผิดพลาด', message, 'error');
  } finally {
    setLoading(false);
  }
};
