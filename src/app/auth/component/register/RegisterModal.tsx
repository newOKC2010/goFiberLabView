'use client';

import { useState } from 'react';
import { Modal } from '@/components/modal/mainModal';
import { Button } from '@/components/buttonClick/mainButton';
import { InputText } from '@/components/input/text/mainInputText';
import { RegisterModalProps, RegisterErrors } from '@/app/auth/component/register/utils/types';
import Dropdown from '@/components/dropdown/mainDropdown';
import { onCidChange, onNameChange, onEmailChange, handleRegisterSubmit } from '@/app/auth/component/register/handler/handlerRegister';

const FACILITY_TYPES = [
  { value: '01', label: '01 - โรงพยาบาลศูนย์/ทั่วไป' },
  { value: '02', label: '02 - โรงพยาบาลชุมชน' },
  { value: '03', label: '03 - โรงพยาบาลส่งเสริมสุขภาพตำบล' },
  { value: '04', label: '04 - สถานีอนามัย' },
  { value: '05', label: '05 - คลินิก' },
];

const emptyErrors = (): RegisterErrors => ({
  cid: '', fullName: '', email: '',
  facilityType: '', facilityCode: '', facilityName: '',
});

export const RegisterModal = ({ isOpen, onClose }: RegisterModalProps) => {
  const [cid, setCid] = useState('');
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [facilityType, setFacilityType] = useState('');
  const [facilityCode, setFacilityCode] = useState('');
  const [facilityName, setFacilityName] = useState('');
  const [errors, setErrors] = useState<RegisterErrors>(emptyErrors());
  const [loading, setLoading] = useState(false);

  const resetForm = () => {
    setCid(''); setFullName(''); setEmail('');
    setFacilityType(''); setFacilityCode(''); setFacilityName('');
    setErrors(emptyErrors());
  };

  const handleClose = () => { resetForm(); onClose(); };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await handleRegisterSubmit(
      cid, fullName, email,
      facilityType, facilityCode, facilityName,
      setErrors, setLoading, handleClose
    );
  };

  return (
    <Modal
      isOpen={isOpen}
      onClose={handleClose}
      title="สมัครใช้งาน"
      contentClassName="!bg-white"
      icon={
        <span
          className="material-symbols-outlined text-blue-500"
          style={{ fontVariationSettings: "'wght' 700", fontSize: 'clamp(2rem, 5vw, 3rem)' }}
        >
          person_add
        </span>
      }
    >
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">เลขบัตรประชาชน</label>
          <InputText
            type="text"
            value={cid}
            onChange={(e) => onCidChange(e.target.value, setCid, setErrors)}
            placeholder="เลขบัตรประชาชน 13 หลัก"
            inputMode="numeric"
            maxLength={13}
            disabled={loading}
            icon="badge"
            error={errors.cid}
          />
        </div>

        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">ชื่อ-นามสกุล</label>
          <InputText
            type="text"
            value={fullName}
            onChange={(e) => onNameChange(e.target.value, setFullName, setErrors)}
            placeholder="ชื่อ นามสกุล (ภาษาไทยเท่านั้น)"
            disabled={loading}
            icon="person"
            error={errors.fullName}
          />
        </div>

        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">อีเมล</label>
          <InputText
            type="text"
            value={email}
            onChange={(e) => onEmailChange(e.target.value, setEmail, setErrors)}
            placeholder="name@example.com"
            disabled={loading}
            icon="mail"
            error={errors.email}
          />
        </div>

        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">ประเภทสถานพยาบาล</label>
          <Dropdown
            options={FACILITY_TYPES}
            value={facilityType}
            onChange={(val) => { setFacilityType(val); setErrors((p) => ({ ...p, facilityType: '' })); }}
            placeholder="-- เลือกประเภทสถานพยาบาล --"
            inModal={true}
          />
          {errors.facilityType && <p className="text-xs text-red-500 ml-1">{errors.facilityType}</p>}
        </div>

        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">รหัสสถานพยาบาล</label>
          <InputText
            type="text"
            value={facilityCode}
            onChange={(e) => {
              const val = e.target.value.replace(/\D/g, '').slice(0, 10);
              setFacilityCode(val);
              setErrors((p) => ({ ...p, facilityCode: '' }));
            }}
            placeholder="รหัสสถานพยาบาล (ตัวเลขไม่เกิน 10 หลัก)"
            disabled={loading}
            icon="tag"
            inputMode="numeric"
            error={errors.facilityCode}
            className="font-bold"
          />
        </div>

        <div className="space-y-1">
          <label className="text-xs font-bold text-gray-700 ml-1">ชื่อสถานพยาบาล</label>
          <InputText
            type="text"
            value={facilityName}
            onChange={(e) => { setFacilityName(e.target.value); setErrors((p) => ({ ...p, facilityName: '' })); }}
            placeholder="ชื่อโรงพยาบาล / คลินิก"
            disabled={loading}
            icon="local_hospital"
            error={errors.facilityName}
          />
        </div>

        <div className="flex justify-center pt-2">
          <Button
            type="submit"
            variant="pastel"
            customWidth={200}
            loading={loading}
            disabled={loading}
            className="rounded-2xl bg-gradient-to-r from-blue-400 to-cyan-500 hover:from-blue-500 hover:to-cyan-600 text-white"
            icon="how_to_reg"
          >
            ลงทะเบียน
          </Button>
        </div>
      </form>
    </Modal>
  );
};
