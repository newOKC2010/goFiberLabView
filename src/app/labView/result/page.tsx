'use client';

import { useState } from 'react';
import CidInput from '@/components/input/cid/mainInputCID';
import HeaderFiltersDateRange from '@/components/dataPicker/headerFiltersDateRange';
import { Button } from '@/components/buttonClick/mainButton';
import { handleGetLabResult } from '@/app/labView/result/handler/handlerLabResult';
import { LabResultResponse, LabGroupedResult, LabCategory } from '@/app/labView/result/utils/types';
import { showToast } from '@/global/globalSwal';
import { formatThaiDate } from '@/components/dataPicker/handler/datePickerHandlers';
import PrintReport from '@/app/labView/result/components/PrintReport';

export default function ResultPage() {
  const today = new Date().toISOString().split('T')[0];
  const [cid, setCid] = useState('');
  const [startDate, setStartDate] = useState(today);
  const [endDate, setEndDate] = useState(today);
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<LabResultResponse | null>(null);
  const [searched, setSearched] = useState(false);
  const [openDates, setOpenDates] = useState<string[]>([]);
  const [isPrintOpen, setIsPrintOpen] = useState(false);

  const toggleDate = (date: string) => {
    setOpenDates(prev =>
      prev.includes(date) ? prev.filter(d => d !== date) : [...prev, date]
    );
  };

  const handleReset = () => {
    setCid('');
    setStartDate(today);
    setEndDate(today);
    setResult(null);
    setSearched(false);
    setOpenDates([]);
  };

  const handleSearch = async () => {
    setLoading(true);
    const res = await handleGetLabResult(cid, startDate, endDate);
    setLoading(false);
    setSearched(true);

    if (!res.success) {
      showToast(res.message || 'เกิดข้อผิดพลาด', 'error');
      setResult(null);
      return;
    }

    setResult(res.data || null);
    if (res.data && res.data.results.length > 0) {
      setOpenDates([res.data.results[0].order_date]);
    }
  };

  const totalItems = (group: LabGroupedResult) =>
    group.groups.reduce((acc, g) => acc + g.items.length, 0);

  return (
    <div className="p-4 sm:p-6 max-w-4xl mx-auto">

      {isPrintOpen && result && (
        <PrintReport
          ptName={result.pt_name}
          cid={cid}
          startDate={startDate}
          endDate={endDate}
          results={result.results}
          onClose={() => setIsPrintOpen(false)}
        />
      )}

      {/* Header */}
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-gray-800 flex items-center gap-2">
          <span className="material-symbols-outlined text-blue-500" style={{ fontVariationSettings: "'wght' 700" }}>
            lab_research
          </span>
          ค้นหาผลตรวจ Lab
        </h1>
        <p className="text-sm text-gray-500 font-bold mt-1 ml-8">กรอกเลขบัตรประชาชนและช่วงวันที่เพื่อค้นหาผล</p>
      </div>

      {/* Filter Card */}
      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 mb-6">
        <div className="mb-5">
          <label className="flex items-center gap-2 text-sm font-bold text-gray-700 mb-3">
            <span className="material-symbols-outlined text-blue-500 text-base" style={{ fontVariationSettings: "'wght' 700" }}>
              badge
            </span>
            เลขบัตรประชาชน 13 หลัก
          </label>
          <div className="flex justify-center sm:justify-start overflow-x-auto pb-1">
            <CidInput value={cid} onChange={setCid} pattern={[1, 4, 5, 2, 1]} size="md" />
          </div>
        </div>

        <div className="mb-5">
          <HeaderFiltersDateRange
            startDate={startDate}
            endDate={endDate}
            onStartDateChange={setStartDate}
            onEndDateChange={setEndDate}
          />
        </div>

        <div className="flex justify-end gap-2">
          <Button
            variant="pastel"
            size="md"
            disabled={loading}
            onClick={handleReset}
            className="bg-gray-100 hover:bg-gray-200 text-gray-600 rounded-xl"
            icon="restart_alt"
          >
            รีเซ็ต
          </Button>
          <Button
            variant="pastel"
            size="md"
            loading={loading}
            disabled={loading || cid.length !== 13 || !startDate || !endDate}
            onClick={handleSearch}
            className="bg-gradient-to-r from-blue-400 to-cyan-500 hover:from-blue-500 hover:to-cyan-600 text-white rounded-xl"
            icon="search"
          >
            ค้นหา
          </Button>
        </div>
      </div>

      {/* Results */}
      {searched && result && (
        <>
          {/* Patient Info */}
          <div className="bg-blue-50 border border-blue-100 rounded-2xl px-5 py-4 mb-4 flex items-center justify-between">
            <div className="flex items-center gap-3">
              <span className="material-symbols-outlined text-blue-500 text-2xl" style={{ fontVariationSettings: "'wght' 700" }}>
                person
              </span>
              <div>
                <p className="text-xs text-gray-500 font-bold">ชื่อผู้ป่วย</p>
                <p className="text-lg font-bold text-gray-800">{result.pt_name || '-'}</p>
              </div>
            </div>
            <div className="flex items-center gap-3">
              <div className="text-right">
                <p className="text-xs text-gray-500 font-bold">รายการทั้งหมด</p>
                <p className="text-xl font-bold text-blue-600">{result.total}</p>
              </div>
              {result.results.length > 0 && (
                <button
                  onClick={() => setIsPrintOpen(true)}
                  className="flex items-center gap-1.5 px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white text-sm font-bold rounded-xl transition cursor-pointer"
                >
                  <span className="material-symbols-outlined text-base" style={{ fontVariationSettings: "'wght' 700" }}>print</span>
                  พิมพ์รายงาน
                </button>
              )}
            </div>
          </div>

          {/* No Data */}
          {result.results.length === 0 && (
            <div className="bg-white rounded-2xl border border-gray-100 p-12 text-center">
              <span className="material-symbols-outlined text-gray-300 text-5xl" style={{ fontVariationSettings: "'wght' 300" }}>
                inbox
              </span>
              <p className="text-gray-500 font-bold mt-3">ไม่พบข้อมูลในช่วงวันที่ที่เลือก</p>
            </div>
          )}

          {/* Grouped by date */}
          <div className="space-y-3">
            {result.results.map((group: LabGroupedResult) => (
              <div key={group.order_date} className="bg-white rounded-2xl border border-gray-100 overflow-hidden shadow-sm">

                {/* Date Header */}
                <button
                  onClick={() => toggleDate(group.order_date)}
                  className="w-full flex items-center justify-between px-5 py-4 hover:bg-gray-50 transition-colors cursor-pointer"
                >
                  <div className="flex items-center gap-3">
                    <span className="material-symbols-outlined text-blue-500 text-lg" style={{ fontVariationSettings: "'wght' 700" }}>
                      calendar_today
                    </span>
                    <span className="font-bold text-gray-800">{formatThaiDate(group.order_date)}</span>
                    <span className="text-xs bg-blue-100 text-blue-600 font-bold px-2 py-0.5 rounded-full">
                      {totalItems(group)} รายการ
                    </span>
                  </div>
                  <span className="material-symbols-outlined text-gray-400">
                    {openDates.includes(group.order_date) ? 'expand_less' : 'expand_more'}
                  </span>
                </button>

                {/* Categories */}
                {openDates.includes(group.order_date) && (
                  <div className="border-t border-gray-100 divide-y divide-gray-50">
                    {group.groups.map((cat: LabCategory) => (
                      <div key={cat.group_name}>

                        {/* Category Label */}
                        <div className="flex items-center gap-2 px-5 py-2.5 bg-gray-50">
                          <span className="material-symbols-outlined text-blue-400 text-sm" style={{ fontVariationSettings: "'wght' 500" }}>
                            category
                          </span>
                          <span className="text-xs font-bold text-blue-700 uppercase tracking-wide">
                            {cat.group_name}
                          </span>
                          <span className="text-xs text-gray-400 font-bold">({cat.items.length})</span>
                        </div>

                        {/* Items Table */}
                        <table className="w-full text-sm">
                          <tbody>
                            {cat.items.map((item, idx) => (
                              <tr key={idx} className={idx % 2 === 0 ? 'bg-white' : 'bg-gray-50/40'}>
                                <td className="px-5 py-2.5 font-bold text-gray-700">{item.lab_items_name}</td>
                                <td className="px-5 py-2.5 text-right font-bold text-blue-700 w-44">
                                  {item.lab_order_result ?? '-'}
                                </td>
                              </tr>
                            ))}
                          </tbody>
                        </table>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            ))}
          </div>
        </>
      )}

      {/* Empty state before search */}
      {!searched && (
        <div className="text-center py-16 text-gray-400">
          <span className="material-symbols-outlined text-5xl" style={{ fontVariationSettings: "'wght' 200" }}>
            biotech
          </span>
          <p className="font-bold mt-3">กรอกข้อมูลด้านบนแล้วกด ค้นหา</p>
        </div>
      )}

    </div>
  );
}