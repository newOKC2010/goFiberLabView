'use client'

import { PaginationData } from '@/components/pagination/handler/handlerPagination'
import { formatPaginationInfo } from '@/components/pagination/handler/handlerPagination'

interface PaginationInfoProps {
  paginationData: PaginationData
}

export default function PaginationInfo({ paginationData }: PaginationInfoProps) {
  const infoText = formatPaginationInfo(paginationData)
  
  return (
    <div className="flex items-center justify-center sm:justify-start">
      <div className="flex items-center gap-2 text-sm text-gray-600 bg-gray-50 px-4 py-2 rounded-lg border">
        <span className="material-symbols-outlined text-blue-500 text-lg">info</span>
        <span className="font-bold">{infoText}</span>
      </div>
    </div>
  )
}