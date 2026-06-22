'use client'

import { PaginationProps } from '@/components/pagination/handler/handlerPagination'
import PaginationControls from '@/components/pagination/component/controlPagiantion'
import PaginationInfo from '@/components/pagination/component/infoPagiantion'

export default function Pagination({ paginationData, onPageChange, loading }: PaginationProps) {
  const { total_pages, current_page } = paginationData

  if (total_pages <= 1) {
    return (
      <div className="flex justify-center">
        <PaginationInfo paginationData={paginationData} />
      </div>
    )
  }

  return (
    <div className="flex flex-col sm:flex-row items-center justify-between gap-4 p-4 bg-white rounded-xl border border-gray-200 shadow-sm">
      <PaginationInfo paginationData={paginationData} />
      
      <PaginationControls 
        currentPage={current_page}
        totalPages={total_pages}
        onPageChange={onPageChange}
        loading={loading}
      />
    </div>
  )
}