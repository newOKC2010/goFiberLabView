'use client';

import { useEffect } from 'react';
import { createPortal } from 'react-dom';
import { LabGroupedResult, LabCategory } from '@/app/labView/result/utils/types';
import { formatThaiDate, THAI_MONTHS } from '@/components/dataPicker/handler/datePickerHandlers';

interface PrintReportProps {
  ptName: string;
  cid: string;
  startDate: string;
  endDate: string;
  results: LabGroupedResult[];
  onClose: () => void;
}

function groupByMonth(results: LabGroupedResult[]): [string, LabGroupedResult[]][] {
  const map = new Map<string, LabGroupedResult[]>();
  for (const group of results) {
    const [year, month] = group.order_date.split('-');
    const key = `${year}-${month}`;
    if (!map.has(key)) map.set(key, []);
    map.get(key)!.push(group);
  }
  return Array.from(map.entries());
}

function formatMonthYear(key: string): string {
  const [year, month] = key.split('-').map(Number);
  return `${THAI_MONTHS[month - 1]} พ.ศ. ${year + 543}`;
}

function maskCID(cid: string): string {
  if (cid.length !== 13) return cid;
  return `${cid[0]}-${cid.slice(1, 5)}-${cid.slice(5, 10)}-${cid.slice(10, 12)}-${cid[12]}`;
}

interface ReportContentProps {
  ptName: string;
  cid: string;
  startDate: string;
  endDate: string;
  results: LabGroupedResult[];
}

function ReportContent({ ptName, cid, startDate, endDate, results }: ReportContentProps) {
  const monthGroups = groupByMonth(results);
  const todayISO = new Date().toISOString().split('T')[0];
  const printedDate = formatThaiDate(todayISO);
  const totalItems = results.reduce((acc, g) => acc + g.groups.reduce((a, c) => a + c.items.length, 0), 0);

  return (
    <div className="bg-white w-full">

      {/* Header */}
      <div style={{ padding: '2.5rem 3rem 1.25rem', borderBottom: '2px solid #374151' }}>
        <div style={{ display: 'flex', alignItems: 'center', gap: '1.25rem', marginBottom: '0.75rem' }}>
          <span className="material-symbols-outlined" style={{ fontSize: '3.5rem', color: '#1d4ed8', fontVariationSettings: "'wght' 200" }}>
            lab_research
          </span>
          <div style={{ flex: 1 }}>
            <h1 style={{ fontSize: '1.2rem', fontWeight: 700, color: '#1f2937', margin: 0 }}>
              รายงานผลการตรวจทางห้องปฏิบัติการ โรงพยาบาลบางเลน
            </h1>
            <p style={{ fontSize: '0.85rem', color: '#6b7280', fontWeight: 700, margin: '0.2rem 0 0' }}>
              Laboratory Investigation Report
            </p>
          </div>
          <div style={{ textAlign: 'right' }}>
            <p style={{ fontSize: '0.7rem', color: '#9ca3af', fontWeight: 700, margin: 0 }}>วันที่พิมพ์</p>
            <p style={{ fontSize: '0.85rem', fontWeight: 700, color: '#374151', margin: 0 }}>{printedDate}</p>
          </div>
        </div>
      </div>

      {/* Patient Info */}
      <div style={{ padding: '1.25rem 3rem', backgroundColor: '#eff6ff', borderBottom: '1px solid #e5e7eb' }}>
        <p style={{ fontSize: '0.65rem', fontWeight: 700, color: '#2563eb', textTransform: 'uppercase', letterSpacing: '0.1em', margin: '0 0 0.75rem' }}>
          ข้อมูลผู้รับบริการ
        </p>
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '0.5rem 2.5rem' }}>
          <InfoRowInline label="ชื่อ-นามสกุล" value={ptName || '-'} />
          <InfoRowInline label="เลขบัตรประชาชน" value={maskCID(cid)} />
          <InfoRowInline label="วันที่เริ่มต้น" value={formatThaiDate(startDate)} />
          <InfoRowInline label="วันที่สิ้นสุด" value={formatThaiDate(endDate)} />
        </div>
      </div>

      {/* Results */}
      <div style={{ padding: '1.5rem 3rem', display: 'flex', flexDirection: 'column', gap: '1.75rem' }}>
        {monthGroups.map(([monthKey, groups]) => (
          <section key={monthKey}>
            {/* Month Header */}
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.75rem', marginBottom: '1rem' }}>
              <div style={{ width: '4px', height: '20px', backgroundColor: '#2563eb', borderRadius: '2px', flexShrink: 0 }} />
              <span style={{ fontSize: '0.8rem', fontWeight: 700, color: '#1d4ed8', textTransform: 'uppercase', letterSpacing: '0.05em' }}>
                {formatMonthYear(monthKey)}
              </span>
              <div style={{ flex: 1, height: '1px', backgroundColor: '#bfdbfe' }} />
              <span style={{ fontSize: '0.7rem', fontWeight: 700, color: '#9ca3af' }}>
                {groups.reduce((a, g) => a + g.groups.reduce((b, c) => b + c.items.length, 0), 0)} รายการ
              </span>
            </div>

            {groups.map((group, gi) => (
              <div key={group.order_date} style={{ marginTop: gi > 0 ? '1.25rem' : 0 }}>
                <p style={{ fontSize: '0.7rem', fontWeight: 700, color: '#6b7280', marginBottom: '0.4rem' }}>
                  วันที่ตรวจ : {formatThaiDate(group.order_date)}
                </p>

                {group.groups.map((cat: LabCategory) => (
                  <div key={cat.group_name} style={{ marginBottom: '0.75rem' }}>
                    {/* Category Label */}
                    <div style={{ backgroundColor: '#f0f9ff', padding: '0.3rem 0.75rem', marginBottom: '0', borderLeft: '3px solid #38bdf8' }}>
                      <span style={{ fontSize: '0.68rem', fontWeight: 700, color: '#0369a1', textTransform: 'uppercase', letterSpacing: '0.06em' }}>
                        {cat.group_name}
                      </span>
                    </div>
                    <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '0.82rem' }}>
                      <thead>
                        <tr style={{ backgroundColor: '#f3f4f6' }}>
                          <th style={thStyle()}>
                            <span style={{ display: 'block', width: '2rem', textAlign: 'center' }}>#</span>
                          </th>
                          <th style={thStyleWide()}>รายการตรวจ</th>
                          <th style={thStyleResult()}>ผลการตรวจ</th>
                        </tr>
                      </thead>
                      <tbody>
                        {cat.items.map((item, idx) => (
                          <tr key={idx} style={{ backgroundColor: idx % 2 === 0 ? '#ffffff' : '#f9fafb' }}>
                            <td style={tdCenter()}>{idx + 1}</td>
                            <td style={tdLeft()}>{item.lab_items_name}</td>
                            <td style={tdResult()}>{item.lab_order_result ?? '-'}</td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                ))}
              </div>
            ))}
          </section>
        ))}
      </div>

      {/* Summary */}
      <div style={{ padding: '0.6rem 3rem', backgroundColor: '#f9fafb', borderTop: '1px solid #e5e7eb' }}>
        <p style={{ fontSize: '0.7rem', fontWeight: 700, color: '#6b7280', textAlign: 'right', margin: 0 }}>
          รายการทั้งหมด {totalItems} รายการ | ข้อมูล {monthGroups.length} เดือน
        </p>
      </div>

      {/* Footer */}
      <div style={{ padding: '1.5rem 3rem 2.5rem', borderTop: '1px solid #e5e7eb' }}>
        <p style={{ textAlign: 'center', fontSize: '0.65rem', color: '#d1d5db', fontWeight: 700, margin: 0 }}>
          เอกสารนี้ออกโดยระบบ Lab View — พิมพ์เมื่อ {printedDate}
        </p>
      </div>
    </div>
  );
}

function InfoRowInline({ label, value }: { label: string; value: string }) {
  return (
    <div style={{ display: 'flex', gap: '0.5rem', fontSize: '0.82rem' }}>
      <span style={{ fontWeight: 700, color: '#6b7280', width: '9rem', flexShrink: 0 }}>{label}</span>
      <span style={{ fontWeight: 700, color: '#1f2937' }}>: {value}</span>
    </div>
  );
}

const border = '1px solid #d1d5db';
function thStyle() {
  return { border, padding: '0.5rem 0.75rem', fontWeight: 700, color: '#374151', textAlign: 'center' as const };
}
function thStyleWide() {
  return { border, padding: '0.5rem 1rem', fontWeight: 700, color: '#374151', textAlign: 'left' as const };
}
function thStyleResult() {
  return { border, padding: '0.5rem 1rem', fontWeight: 700, color: '#374151', textAlign: 'center' as const, width: '11rem' };
}
function tdCenter() {
  return { border, padding: '0.5rem 0.75rem', textAlign: 'center' as const, fontWeight: 700, color: '#9ca3af' };
}
function tdLeft() {
  return { border, padding: '0.6rem 1rem', fontWeight: 700, color: '#1f2937' };
}
function tdResult() {
  return { border, padding: '0.6rem 1rem', textAlign: 'center' as const, fontWeight: 700, color: '#1d4ed8' };
}

export default function PrintReport({ ptName, cid, startDate, endDate, results, onClose }: PrintReportProps) {
  useEffect(() => {
    const style = document.createElement('style');
    style.id = 'lab-print-styles';
    style.textContent = `
      @page { margin: 0; }
      @media print {
        body > *:not(#lab-print-portal) { display: none !important; }
        #lab-print-portal { display: block !important; }
      }
    `;
    document.head.appendChild(style);
    return () => { document.getElementById('lab-print-styles')?.remove(); };
  }, []);

  const props = { ptName, cid, startDate, endDate, results };

  return (
    <>
      {/* Screen overlay */}
      <div className="fixed inset-0 z-50 bg-black/50 flex items-start justify-center overflow-y-auto py-8 px-4">

        {/* Action Bar */}
        <div className="fixed top-4 right-4 z-[60] flex gap-2">
          <button
            onClick={() => window.print()}
            className="flex items-center gap-2 px-5 py-2.5 bg-blue-600 hover:bg-blue-700 text-white text-sm font-bold rounded-xl shadow-lg transition cursor-pointer"
          >
            <span className="material-symbols-outlined text-base" style={{ fontVariationSettings: "'wght' 700" }}>print</span>
            พิมพ์รายงาน
          </button>
          <button
            onClick={onClose}
            className="flex items-center gap-2 px-4 py-2.5 bg-white hover:bg-gray-100 text-gray-700 text-sm font-bold rounded-xl shadow-lg transition cursor-pointer"
          >
            <span className="material-symbols-outlined text-base">close</span>
            ปิด
          </button>
        </div>

        {/* Preview Paper */}
        <div className="bg-white w-full max-w-3xl rounded-xl shadow-2xl overflow-hidden">
          <ReportContent {...props} />
        </div>
      </div>

      {/* Print Portal */}
      {createPortal(
        <div id="lab-print-portal" style={{ display: 'none' }}>
          <ReportContent {...props} />
        </div>,
        document.body
      )}
    </>
  );
}