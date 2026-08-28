'use client';

import { useState, useEffect } from 'react';
import { createPortal } from 'react-dom';
import { LabGroupedResult } from '@/app/labView/result/utils/types';
import { formatThaiDate } from '@/components/dataPicker/handler/datePickerHandlers';

interface FlatItem {
  id: string;
  order_date: string;
  group_name: string;
  lab_items_name: string;
  lab_order_result: string | null;
  lab_items_normal_value_ref: string | null;
}

interface PrintStickerProps {
  ptName: string;
  results: LabGroupedResult[];
  onClose: () => void;
}

function flattenItems(results: LabGroupedResult[]): FlatItem[] {
  const items: FlatItem[] = [];
  results.forEach(group => {
    group.groups.forEach(cat => {
      cat.items.forEach((item, idx) => {
        items.push({
          id: `${group.order_date}-${cat.group_name}-${idx}`,
          order_date: group.order_date,
          group_name: cat.group_name,
          lab_items_name: item.lab_items_name,
          lab_order_result: item.lab_order_result,
          lab_items_normal_value_ref: item.lab_items_normal_value_ref,
        });
      });
    });
  });
  return items;
}

function chunkArray<T>(arr: T[], size: number): T[][] {
  const chunks: T[][] = [];
  for (let i = 0; i < arr.length; i += size) chunks.push(arr.slice(i, i + size));
  return chunks;
}

const font = 'Mitr, sans-serif';

function StickerPage({ ptName, items, isLast }: { ptName: string; items: FlatItem[]; isLast: boolean }) {
  const date = items[0]?.order_date ?? '';
  return (
    <div style={{
      width: '3.5in', height: '2.5in',
      padding: '0.12in 0.15in',
      boxSizing: 'border-box',
      display: 'flex', flexDirection: 'column',
      fontFamily: font,
      pageBreakAfter: isLast ? 'auto' : 'always',
      overflow: 'hidden',
    }}>

      {/* Header */}
      <div style={{ borderBottom: '1.5px solid #374151', paddingBottom: '4pt', marginBottom: '4pt' }}>
        <div style={{ fontSize: '7pt', fontWeight: 600, color: '#374151' }}>
          โรงพยาบาลบางเลน · ผลตรวจ Lab
        </div>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginTop: '2pt' }}>
          <span style={{ fontSize: '8.5pt', fontWeight: 700, color: '#1f2937' }}>{ptName}</span>
          <span style={{ fontSize: '6.5pt', fontWeight: 400, color: '#6b7280' }}>{formatThaiDate(date)}</span>
        </div>
      </div>

      {/* Table */}
      <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '8pt', fontFamily: font }}>
        <thead>
          <tr style={{ backgroundColor: '#f3f4f6' }}>
            <th style={{ textAlign: 'left', padding: '2pt 4pt', fontWeight: 600, color: '#374151', fontSize: '7pt', borderBottom: '1px solid #d1d5db' }}>
              รายการตรวจ
            </th>
            <th style={{ textAlign: 'center', padding: '2pt 4pt', fontWeight: 600, color: '#374151', fontSize: '7pt', borderBottom: '1px solid #d1d5db', width: '22%' }}>
              ผล
            </th>
            <th style={{ textAlign: 'center', padding: '2pt 4pt', fontWeight: 600, color: '#374151', fontSize: '7pt', borderBottom: '1px solid #d1d5db', width: '30%' }}>
              ค่าปกติ
            </th>
          </tr>
        </thead>
        <tbody>
          {items.map((item, idx) => (
            <tr key={item.id} style={{ backgroundColor: idx % 2 === 0 ? '#ffffff' : '#f9fafb' }}>
              <td style={{ padding: '2.5pt 4pt', fontWeight: 500, color: '#1f2937', fontSize: '7.5pt' }}>
                {item.lab_items_name}
              </td>
              <td style={{ padding: '2.5pt 4pt', textAlign: 'center', fontWeight: 700, color: '#1d4ed8', fontSize: '7.5pt' }}>
                {item.lab_order_result ?? '-'}
              </td>
              <td style={{ padding: '2.5pt 4pt', textAlign: 'center', fontWeight: 400, color: '#6b7280', fontSize: '6.5pt' }}>
                {item.lab_items_normal_value_ref ?? '-'}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default function PrintSticker({ ptName, results, onClose }: PrintStickerProps) {
  const allItems = flattenItems(results);
  const [selected, setSelected] = useState<Set<string>>(new Set(allItems.map(i => i.id)));

  useEffect(() => {
    const style = document.createElement('style');
    style.id = 'lab-sticker-print-styles';
    style.textContent = `
      @page { size: 3.5in 2.5in; margin: 0mm; }
      @media print {
        body > *:not(#lab-sticker-portal) { display: none !important; }
        #lab-sticker-portal { display: block !important; }
        html, body { margin: 0 !important; padding: 0 !important; }
      }
    `;
    document.head.appendChild(style);
    return () => { document.getElementById('lab-sticker-print-styles')?.remove(); };
  }, []);

  const toggleItem = (id: string) => {
    setSelected(prev => {
      const next = new Set(prev);
      next.has(id) ? next.delete(id) : next.add(id);
      return next;
    });
  };

  const toggleAll = () =>
    setSelected(prev => prev.size === allItems.length ? new Set() : new Set(allItems.map(i => i.id)));

  const selectedItems = allItems.filter(i => selected.has(i.id));
  const pages = chunkArray(selectedItems, 5);
  const allChecked = selected.size === allItems.length;

  const groupedByDate = allItems.reduce<Record<string, FlatItem[]>>((acc, item) => {
    (acc[item.order_date] ??= []).push(item);
    return acc;
  }, {});

  return (
    <>
      <div className="fixed inset-0 z-50 bg-black/50 flex items-center justify-center p-4">
        <div className="bg-white rounded-2xl shadow-2xl w-full max-w-lg max-h-[85vh] flex flex-col">

          {/* Header */}
          <div className="flex items-center justify-between px-5 py-4 border-b border-gray-100">
            <div className="flex items-center gap-2">
              <span className="material-symbols-outlined text-purple-500" style={{ fontVariationSettings: "'wght' 700" }}>
                label
              </span>
              <h2 className="font-bold text-gray-800">พิมพ์ Sticker Lab</h2>
              <span className="text-xs bg-purple-100 text-purple-600 font-bold px-2 py-0.5 rounded-full">
                {selectedItems.length} รายการ · {pages.length} แผ่น
              </span>
            </div>
            <button onClick={onClose} className="p-1 hover:bg-gray-100 rounded-lg transition cursor-pointer">
              <span className="material-symbols-outlined text-gray-500">close</span>
            </button>
          </div>

          {/* Select All */}
          <div className="px-5 py-3 border-b border-gray-50 flex items-center justify-between">
            <label className="flex items-center gap-2 cursor-pointer select-none">
              <input type="checkbox" checked={allChecked} onChange={toggleAll} className="w-4 h-4 accent-purple-500" />
              <span className="text-sm font-bold text-gray-600">เลือกทั้งหมด</span>
            </label>
            <p className="text-xs text-gray-400 font-bold">5 รายการ / แผ่น · 3.5&quot; × 2.5&quot;</p>
          </div>

          {/* Item List */}
          <div className="flex-1 overflow-y-auto px-5 py-3 space-y-4">
            {Object.entries(groupedByDate).map(([date, items]) => (
              <div key={date}>
                <p className="text-xs font-bold text-blue-600 mb-2 flex items-center gap-1">
                  <span className="material-symbols-outlined text-sm" style={{ fontVariationSettings: "'wght' 700" }}>calendar_today</span>
                  {formatThaiDate(date)}
                </p>
                <div className="space-y-1">
                  {items.map(item => (
                    <label key={item.id} className="flex items-center gap-3 p-2.5 rounded-xl hover:bg-gray-50 cursor-pointer select-none">
                      <input
                        type="checkbox"
                        checked={selected.has(item.id)}
                        onChange={() => toggleItem(item.id)}
                        className="w-4 h-4 accent-purple-500 flex-shrink-0"
                      />
                      <div className="flex-1 min-w-0">
                        <p className="text-sm font-bold text-gray-700 truncate">{item.lab_items_name}</p>
                        <p className="text-xs text-gray-400">{item.group_name}{item.lab_items_normal_value_ref ? ` · ค่าปกติ: ${item.lab_items_normal_value_ref}` : ''}</p>
                      </div>
                      <span className="text-sm font-bold text-blue-700 flex-shrink-0">{item.lab_order_result ?? '-'}</span>
                    </label>
                  ))}
                </div>
              </div>
            ))}
          </div>

          {/* Footer */}
          <div className="px-5 py-4 border-t border-gray-100 flex gap-2 justify-end">
            <button
              onClick={onClose}
              className="px-4 py-2 text-sm font-bold text-gray-600 hover:bg-gray-100 rounded-xl transition cursor-pointer"
            >
              ยกเลิก
            </button>
            <button
              onClick={() => window.print()}
              disabled={selectedItems.length === 0}
              className="flex items-center gap-2 px-5 py-2 bg-purple-500 hover:bg-purple-600 disabled:opacity-50 disabled:cursor-not-allowed text-white text-sm font-bold rounded-xl transition cursor-pointer"
            >
              <span className="material-symbols-outlined text-base" style={{ fontVariationSettings: "'wght' 700" }}>print</span>
              พิมพ์ {pages.length} แผ่น
            </button>
          </div>
        </div>
      </div>

      {/* Print Portal */}
      {createPortal(
        <div id="lab-sticker-portal" style={{ display: 'none' }}>
          {pages.map((pageItems, pageIdx) => (
            <StickerPage
              key={pageIdx}
              ptName={ptName}
              items={pageItems}
              isLast={pageIdx === pages.length - 1}
            />
          ))}
        </div>,
        document.body
      )}
    </>
  );
}
