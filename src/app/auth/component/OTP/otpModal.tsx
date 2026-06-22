'use client';

import { useState, useEffect } from 'react';
import { Modal } from '@/components/modal/mainModal';
import { Button } from '@/components/buttonClick/mainButton';

import CidInput from '@/components/input/cid/mainInputCID';

interface OtpModalProps {
  isOpen: boolean;
  onClose: () => void;
  email: string;
  onVerify: (otp: string) => void;
  loading?: boolean;
}

export const OtpModal = ({ isOpen, onClose, email, onVerify, loading = false }: OtpModalProps) => {
  const [otp, setOtp] = useState('');

  useEffect(() => {
    if (isOpen) {
      setOtp('');
    }
  }, [isOpen]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (otp.length === 6) {
      onVerify(otp);
    }
  };

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      title="ยืนยัน OTP"
      contentClassName="!bg-white/80 !backdrop-blur-xl !border !border-white/50"
      icon={
        <span className="material-symbols-outlined text-5xl text-blue-500"
          style={{
            fontVariationSettings: "'wght' 700",
            fontSize: 'clamp(2rem, 5vw, 3rem)',
            transition: 'all 0.3s ease'
          }}>
          shield_lock
        </span>
      }
    >
      <div className="space-y-6">
        <p className="text-center text-gray-600 text-sm font-bold">
          กรอกรหัส OTP 6 หลัก ที่ส่งไปยัง<br />
          <span className="font-bold text-gray-800">{email}</span>
          <br />
          หรือ line หมอพร้อมของคุณ
        </p>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="flex justify-center py-6 sm:py-0 scale-130 sm:scale-100">
            <CidInput
              value={otp}
              onChange={setOtp}
              pattern={[1, 1, 1, 1, 1, 1]}
              size="lg"
              disabled={loading}
            />
          </div>

          <div className="flex justify-center">
            <Button
              type="submit"
              variant="pastel"
              customWidth={200}
              loading={loading}
              disabled={loading || otp.length !== 6}
              className="rounded-2xl bg-gradient-to-r from-blue-400 to-cyan-500 hover:from-blue-500 hover:to-cyan-600 text-white"
              icon="check_circle"
            >
              ยืนยัน
            </Button>
          </div>

          <p className="text-center text-blue-500 text-xs font-bold flex items-center justify-center gap-1">
            <span className="material-symbols-outlined text-sm"
              style={{
                fontVariationSettings: "'wght' 700",
                fontSize: 'clamp(1rem, 2vw, 1.5rem)',
                transition: 'all 0.3s ease'
              }}>
              schedule
            </span>
            รหัส OTP มีอายุ 2 นาที
          </p>
        </form>
      </div>
    </Modal>
  );
};
