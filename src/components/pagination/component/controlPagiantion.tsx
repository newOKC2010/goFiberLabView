'use client'

import { PaginationControlsProps } from '@/components/pagination/handler/handlerPagination'
import { calculatePageNumbers, validatePageNumber } from '@/components/pagination/handler/handlerPagination'

export default function PaginationControls({ currentPage, totalPages, onPageChange, loading }: PaginationControlsProps) {
  const pages = calculatePageNumbers(currentPage, totalPages)
  
  const handlePageChange = (page: number) => {
    if (loading) return
    const validPage = validatePageNumber(page, totalPages)
    onPageChange(validPage)
  }

  const buttonClass = (isActive = false, isDisabled = false) => {
    const baseClass = "px-2 py-1 rounded-lg font-bold text-xs transition-all duration-200 "
    
    if (isDisabled) {
      return baseClass + "bg-gray-200 text-gray-400 cursor-not-allowed"
    }
    if (isActive) {
      return baseClass + "bg-gradient-to-r from-blue-500 to-blue-600 text-white shadow-md cursor-pointer"
    }
    return baseClass + "bg-white text-gray-700 border border-gray-300 hover:bg-blue-50 hover:border-blue-300 hover:scale-105 shadow-sm cursor-pointer"
  }

  return (
    <div className="flex items-center justify-center space-x-1">
      {/* First Page Button */}
      <button
        onClick={() => handlePageChange(1)}
        disabled={currentPage <= 1 || loading}
        className={buttonClass(false, currentPage <= 1 || loading)}
        title="หน้าแรก"
      >
        <span className="material-symbols-outlined text-base">first_page</span>
      </button>

      {/* Previous Button */}
      <button
        onClick={() => handlePageChange(currentPage - 1)}
        disabled={currentPage <= 1 || loading}
        className={buttonClass(false, currentPage <= 1 || loading)}
        title="หน้าก่อน"
      >
        <span className="material-symbols-outlined text-base">chevron_left</span>
      </button>

      {/* Page Numbers - Hidden on Mobile */}
      <div className="hidden sm:flex space-x-1">
        {pages.map((page, index) => (
          <button
            key={index}
            onClick={() => typeof page === 'number' ? handlePageChange(page) : undefined}
            disabled={loading || typeof page === 'string'}
            className={
              typeof page === 'string' 
                ? "px-2 py-1 text-gray-500 cursor-default font-bold text-xs"
                : buttonClass(page === currentPage, loading)
            }
          >
            {page}
          </button>
        ))}
      </div>

      {/* Current Page Info - Mobile Only */}
      <div className="sm:hidden px-2 py-1 bg-gradient-to-r from-blue-500 to-blue-600 text-white rounded-lg font-bold text-xs">
        {currentPage}/{totalPages}
      </div>

      {/* Next Button */}
      <button
        onClick={() => handlePageChange(currentPage + 1)}
        disabled={currentPage >= totalPages || loading}
        className={buttonClass(false, currentPage >= totalPages || loading)}
        title="หน้าถัดไป"
      >
        <span className="material-symbols-outlined text-base">chevron_right</span>
      </button>

      {/* Last Page Button */}
      <button
        onClick={() => handlePageChange(totalPages)}
        disabled={currentPage >= totalPages || loading}
        className={buttonClass(false, currentPage >= totalPages || loading)}
        title="หน้าสุดท้าย"
      >
        <span className="material-symbols-outlined text-base">last_page</span>
      </button>
    </div>
  )
}