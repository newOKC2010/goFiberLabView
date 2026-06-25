'use client';

import { useState, useEffect, useCallback } from 'react';
import Table from '@/components/table/mainTable';
import Pagination from '@/components/pagination/mainPagination';
import { PaginationData } from '@/components/pagination/handler/handlerPagination';
import { handleGetUsers, handleUpdateStatus } from '@/app/labView/users/handler/handlerUsers';
import { UserItem, GetUsersResponse } from '@/app/labView/users/utils/types';
import { showToast, showConfirm } from '@/global/globalSwal';

const PAGE_SIZE = 10;

const formatThaiDateTime = (iso: string): string => {
  if (!iso) return '-';
  const d = new Date(iso);
  return d.toLocaleDateString('th-TH', { year: 'numeric', month: 'short', day: 'numeric' });
};

export default function UsersPage() {
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [searchInput, setSearchInput] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [togglingId, setTogglingId] = useState<number | null>(null);
  const [responseData, setResponseData] = useState<GetUsersResponse | null>(null);

  const loadUsers = useCallback(async (page: number, searchVal: string) => {
    setLoading(true);
    const res = await handleGetUsers(page, PAGE_SIZE, searchVal);
    setLoading(false);
    if (!res.success) {
      showToast(res.message || 'โหลดข้อมูลไม่สำเร็จ', 'error');
      return;
    }
    setResponseData(res.data || null);
  }, []);

  useEffect(() => {
    loadUsers(currentPage, search);
  }, [currentPage, search, loadUsers]);

  const handleSearchSubmit = () => {
    setCurrentPage(1);
    setSearch(searchInput);
  };

  const handleSearchKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') handleSearchSubmit();
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const pageData: UserItem[] = responseData?.data || [];

  const paginationData: PaginationData = {
    count: pageData.length,
    total_count: responseData?.total_count || 0,
    total_pages: responseData?.total_pages || 1,
    current_page: responseData?.current_page || 1,
  };

  const handleToggleStatus = async (user: UserItem) => {
    const newStatus = !user.status;
    const action = newStatus ? 'อนุมัติ' : 'ระงับ';
    const confirmed = await showConfirm(
      `ต้องการ${action}ผู้ใช้ "${user.full_name}" ใช่หรือไม่?`,
      {},
      { confirm: newStatus ? '#3b82f6' : '#ef4444', cancel: '#6b7280' },
      { confirm: action, cancel: 'ยกเลิก' }
    );
    if (!confirmed) return;

    setTogglingId(user.id);
    const res = await handleUpdateStatus(user.id, newStatus);
    setTogglingId(null);

    if (!res.success) {
      showToast(res.message || 'อัพเดทสถานะไม่สำเร็จ', 'error');
      return;
    }
    showToast(`${action}ผู้ใช้สำเร็จ`, 'success');
    await loadUsers(currentPage, search);
  };

  const columns = [
    {
      key: 'index',
      label: '#',
      render: (_: UserItem) => {
        const idx = pageData.indexOf(_);
        return <span>{(paginationData.current_page - 1) * PAGE_SIZE + idx + 1}</span>;
      },
    },
    {
      key: 'full_name',
      label: 'ชื่อ-นามสกุล',
      render: (u: UserItem) => <span className="font-bold text-gray-800">{u.full_name}</span>,
    },
    {
      key: 'facility_name',
      label: 'สถานพยาบาล',
      render: (u: UserItem) => <span className="font-bold text-gray-600">{u.facility_name || '-'}</span>,
    },
    {
      key: 'created_at',
      label: 'วันที่สมัคร',
      render: (u: UserItem) => <span>{formatThaiDateTime(u.created_at)}</span>,
    },
    {
      key: 'status',
      label: 'สถานะ',
      render: (u: UserItem) => (
        <span className={`inline-flex items-center gap-1 px-3 py-1 rounded-full text-xs font-bold
          ${u.status ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-600'}`}>
          <span className={`w-1.5 h-1.5 rounded-full ${u.status ? 'bg-green-500' : 'bg-red-400'}`} />
          {u.status ? 'อนุมัติแล้ว' : 'รอการอนุมัติ'}
        </span>
      ),
    },
  ];

  return (
    <div className="p-4 sm:p-6 max-w-5xl mx-auto">

      {/* Header */}
      <div className="mb-6 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
        <div>
          <h1 className="text-2xl font-bold text-gray-800 flex items-center gap-2">
            <span className="material-symbols-outlined text-blue-500" style={{ fontVariationSettings: "'wght' 700" }}>
              manage_accounts
            </span>
            จัดการผู้ใช้งาน
          </h1>
          <p className="text-sm text-gray-500 font-bold mt-1 ml-8">
            ทั้งหมด {responseData?.total_count ?? '-'} บัญชี
          </p>
        </div>

        {/* Search */}
        <div className="flex gap-2 w-full sm:w-auto">
          <div className="relative flex-1 sm:w-64">
            <span className="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 text-base pointer-events-none"
              style={{ fontVariationSettings: "'wght' 400" }}>
              search
            </span>
            <input
              type="text"
              value={searchInput}
              onChange={e => setSearchInput(e.target.value)}
              onKeyDown={handleSearchKeyDown}
              placeholder="ค้นหาชื่อ..."
              className="w-full pl-9 pr-4 py-2.5 text-sm font-bold border border-gray-200 rounded-xl bg-white focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-blue-400 transition"
            />
          </div>
          <button
            onClick={handleSearchSubmit}
            className="px-4 py-2.5 text-sm font-bold bg-blue-500 hover:bg-blue-600 text-white rounded-xl transition cursor-pointer"
          >
            ค้นหา
          </button>
        </div>
      </div>

      {/* Table Card */}
      <div className="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden mb-4">
        <Table<UserItem>
          columns={columns}
          data={pageData}
          loading={loading || togglingId !== null}
          onToggleStatus={handleToggleStatus}
          getItemId={u => u.id}
          getItemStatus={u => u.status}
          statusButtonText="อนุมัติ"
        />
      </div>

      {/* Pagination */}
      {!loading && (responseData?.total_count ?? 0) > 0 && (
        <Pagination
          paginationData={paginationData}
          onPageChange={handlePageChange}
          loading={loading}
        />
      )}
    </div>
  );
}